package handler

import (
	"avito-test/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createRemittance(c *gin.Context) {
	var transaction dto.Remittance

	if err := c.BindJSON(&transaction); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := h.services.Transaction.CreateRemittance(transaction)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) createReservation(c *gin.Context) {
	var transaction dto.Reservation

	if err := c.BindJSON(&transaction); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := h.services.Transaction.CreateReservation(transaction)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) createDeposit(c *gin.Context) {
	var transaction dto.Deposit

	if err := c.BindJSON(&transaction); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := h.services.Transaction.CreateDeposit(transaction)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
