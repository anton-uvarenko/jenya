package handlers

import "github.com/gin-gonic/gin"

type SomeHandler struct {
	deps someDep // some dependencies
}

type someDep interface{}

func NewSomeHandler(deps someDep) *SomeHandler {
	return &SomeHandler{
		deps: deps,
	}
}

func (h *SomeHandler) RegisterRouts(engine *gin.Engine) {
	router := engine.Group("/path")
}

func (h *SomeHandler) SomeFunction(ctx *gin.Context) {
}
