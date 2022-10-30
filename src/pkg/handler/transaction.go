package handler

import (
	"avito-test/dto"
	"avito-test/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createTransaction(c *gin.Context) {
	var transaction dto.Transaction

	if err := c.BindJSON(&transaction); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := h.services.Transaction.CreateTransaction(transaction, model.Remittance)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
