package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Hello World!")
	})

	return router
}
