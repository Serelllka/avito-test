package handler

import (
	"avito-test/dto"
	"avito-test/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createRemittance(c *gin.Context) {
	h.createCustomTransaction(c, model.Remittance)
}

func (h *Handler) createReservation(c *gin.Context) {
	h.createCustomTransaction(c, model.Reservation)
}

func (h *Handler) createDeposit(c *gin.Context) {
	h.createCustomTransaction(c, model.Deposit)
}

func (h *Handler) createCustomTransaction(c *gin.Context, trType model.TransactionType) {
	var transaction dto.Transaction

	if err := c.BindJSON(&transaction); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := validateTransaction(&transaction, trType); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Transaction.CreateTransaction(transaction, trType)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func validateTransaction(transaction *dto.Transaction, trType model.TransactionType) error {
	switch trType {
	case model.Remittance:
		if transaction.ProducerId == 0 || transaction.ConsumerId == 0 {
			break
		}
		return nil
	case model.Deposit:
		transaction.ProducerId = 0
		if transaction.ConsumerId == 0 {
			break
		}
		return nil
	case model.Reservation:
		transaction.ConsumerId = 0
		if transaction.ProducerId == 0 || transaction.ServiceId == 0 {
			break
		}
		return nil
	default:
		return fmt.Errorf("uresolved transaction type %d", trType)
	}
	return fmt.Errorf("consumer, producer, service id can't be null with given transaction type: %d", trType)
}
