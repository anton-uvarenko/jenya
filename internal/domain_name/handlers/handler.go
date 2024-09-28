package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	SomeHandler *SomeHandler
}

func NewHandler(deps any) *Handler {
	return &Handler{
		SomeHandler: NewSomeHandler(deps),
	}
}

func (h *Handler) RegisterRouts() *gin.Engine {
	router := gin.New()

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello from upper level handler")
	})

	someHandlerRouts := router.Group("/path")
	h.SomeHandler.RegisterRouts(someHandlerRouts)

	return router
}
