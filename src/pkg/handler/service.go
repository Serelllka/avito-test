package handler

import (
	"avito-test/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) CreateService(c *gin.Context) {
	var input dto.Service
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Maintenance.CreateService(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
