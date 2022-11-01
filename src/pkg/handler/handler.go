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
		payment := transaction.Group("/payment")
		{
			payment.POST("/", h.createPayment)
		}
	}

	user := router.Group("/user")
	{
		user.POST("/", h.createUser)

		balance := user.Group("/balance")
		{
			balance.GET("/:id", h.getUserBalance)
		}
	}

	serv := router.Group("/service")
	{
		serv.POST("/", h.createService)
	}
	return router
}
