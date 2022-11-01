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

func TestHandler_Deposit(t *testing.T) {
	type mockBehavior func(s *mock_service.MockTransaction, dto dto.Deposit)

	formRequest := func(consumerId int, description string, amount int) string {
		return fmt.Sprintf(`
			{
				"consumerId": %d,
				"description" : "%s",
				"amount" : %d
			}
		`, consumerId, description, amount)
	}

	formExpectedBodyFromId := func(id int) string {
		return fmt.Sprintf(`{"id":%d}`, id)
	}

	for _, testCase := range []struct {
		name               string
		input              string
		inputDto           dto.Deposit
		mockBehavior       mockBehavior
		expectedStatusCode int
		expectedBody       string
		reqBody            bool
	}{
		{
			name:  "OK",
			input: formRequest(1, "description1", 100),
			inputDto: dto.Deposit{
				ConsumerId:  1,
				Description: "description1",
				Amount:      100,
			},
			mockBehavior: func(s *mock_service.MockTransaction, dto dto.Deposit) {
				s.EXPECT().CreateDeposit(dto).Return(1, nil)
			},
			expectedStatusCode: 200,
			reqBody:            true,
			expectedBody:       formExpectedBodyFromId(1),
		},
		{
			name:               "BadRequest",
			input:              `{"description":"%s"}`,
			mockBehavior:       func(s *mock_service.MockTransaction, dto dto.Deposit) {},
			expectedStatusCode: 400,
		},
		{
			name:  "Server Error",
			input: formRequest(1, "description1", 100),
			inputDto: dto.Deposit{
				ConsumerId:  1,
				Description: "description1",
				Amount:      100,
			},
			mockBehavior: func(s *mock_service.MockTransaction, dto dto.Deposit) {
				s.EXPECT().CreateDeposit(dto).Return(0, fmt.Errorf("oops"))
			},
			expectedStatusCode: 500,
			reqBody:            true,
			expectedBody:       `{"message":"oops"}`,
		},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			repo := mock_service.NewMockTransaction(c)
			testCase.mockBehavior(repo, testCase.inputDto)

			services := &service.Service{
				Transaction: repo,
			}
			handler := Handler{
				services,
			}

			r := gin.New()
			r.POST("/transaction/deposit/", handler.createDeposit)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/transaction/deposit/",
				bytes.NewBufferString(testCase.input))

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			if testCase.reqBody {
				assert.Equal(t, w.Body.String(), testCase.expectedBody)
			}
		})
	}
}
