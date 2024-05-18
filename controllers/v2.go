package controllers

import (
	v2 "fx_101/handlers/v2"

	"github.com/gin-gonic/gin"
)

type V2Router struct {
	handler *v2.RouteHandler
}

func NewV2Router(handler *v2.RouteHandler) *V2Router {
	router := &V2Router{
		handler: handler,
	}

	return router
}

func (r *V2Router) SetupRoutes(baseRouter *gin.RouterGroup) {
	v2Router := baseRouter.Group("/v2")
	v2Router.GET("/mongo", r.handler.GetMongo)
}
