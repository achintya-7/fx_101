package controllers

import (
	v1 "fx_101/handlers/v1"

	"github.com/gin-gonic/gin"
)

type V1Router struct {
	handler *v1.RouteHandler
}

func NewV1Router(handler *v1.RouteHandler) *V1Router {
	router := &V1Router{
		handler: handler,
	}

	return router
}

func (r *V1Router) SetupRoutes(route *gin.RouterGroup) {
	v1Route := route.Group("/v1")

	v1Route.GET("/mongo", r.handler.GetMongo)
	v1Route.GET("/postgres", r.handler.GetPostgres)
	v1Route.GET("/pubsub", r.handler.GetPubSub)
}
