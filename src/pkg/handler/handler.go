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
		transaction.POST("/", h.createTransaction)
	}
	router.POST("/", h.createUser)
	return router
}
