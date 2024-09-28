package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SomeHandler struct {
	deps someDep // some dependencies
}

type someDep interface{}

func NewSomeHandler(deps someDep) *SomeHandler {
	return &SomeHandler{
		deps: deps,
	}
}

func (h *SomeHandler) RegisterRouts(engine *gin.RouterGroup) {
	engine.GET("/", h.SomeFunction)
}

func (h *SomeHandler) SomeFunction(ctx *gin.Context) {
	ctx.String(http.StatusOK, "jenya zdarova")
}
