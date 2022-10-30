package handler

import (
	"avito-test/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	transaction := router.Group("/transaction")
	{
		remittance := transaction.Group("/remittance")
		{
			remittance.POST("/", h.createRemittance)
		}
		deposit := transaction.Group("/deposit")
		{
			deposit.POST("/", h.createDeposit)
		}
		reservation := transaction.Group("/reservation")
		{
			reservation.POST("/", h.createReservation)
		}
	}

	user := router.Group("/user")
	{
		user.POST("/", h.createUser)
	}
	return router
}
