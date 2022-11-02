package handler

import (
	"avito-test/dto"
	"avito-test/pkg/service"
	mock_service "avito-test/pkg/service/mock"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"net/http/httptest"
	"testing"
)

func TestHandler_Service(t *testing.T) {
	type mockBehavior func(s *mock_service.MockMaintenance, dto dto.Service)

	formRequest := func(title, description string) string {
		return fmt.Sprintf(`
			{
				"title" : "%s",
				"description" : "%s"
			}
		`, title, description)
	}

	formExpectedBodyFromId := func(id int) string {
		return fmt.Sprintf(`{"id":%d}`, id)
	}

	for _, testCase := range []struct {
		name               string
		input              string
		inputDto           dto.Service
		mockBehavior       mockBehavior
		expectedStatusCode int
		expectedBody       string
		reqBody            bool
	}{
		{
			name:  "OK",
			input: formRequest("Service1", "Description1"),
			inputDto: dto.Service{
				Title:       "Service1",
				Description: "Description1",
			},
			mockBehavior: func(s *mock_service.MockMaintenance, dto dto.Service) {
				s.EXPECT().CreateService(dto).Return(1, nil)
			},
			expectedStatusCode: 200,
			reqBody:            true,
			expectedBody:       formExpectedBodyFromId(1),
		},
		{
			name:               "BadRequest",
			input:              `{"description":"%s"}`,
			mockBehavior:       func(s *mock_service.MockMaintenance, dto dto.Service) {},
			expectedStatusCode: 400,
		},
		{
			name:  "Server Error",
			input: formRequest("Service1", "Description1"),
			inputDto: dto.Service{
				Title:       "Service1",
				Description: "Description1",
			},
			mockBehavior: func(s *mock_service.MockMaintenance, dto dto.Service) {
				s.EXPECT().CreateService(dto).Return(0, fmt.Errorf("oops"))
			},
			expectedStatusCode: 500,
			reqBody:            true,
			expectedBody:       `{"message":"oops"}`,
		},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockMaintenance(c)
			testCase.mockBehavior(repo, testCase.inputDto)

			services := &service.Service{
				Maintenance: repo,
			}
			handler := Handler{
				services,
			}

			r := gin.New()
			r.POST("/service/", handler.createService)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/service/",
				bytes.NewBufferString(testCase.input))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			if testCase.reqBody {
				assert.Equal(t, w.Body.String(), testCase.expectedBody)
			}
		})
	}
}
