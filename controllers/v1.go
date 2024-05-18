package controllers

import (
	v1 "fx_101/handlers/v1"
	"fx_101/mongo"
	"fx_101/postgres"
	"fx_101/pubsub"

	"github.com/gin-gonic/gin"
)

type v1Router struct {
	handlers *v1.RouteHandler
	mongo    *mongo.Client
	postgres *postgres.Client
	pubSub   *pubsub.Client
}

func NewRouter(mongo *mongo.Client, postgres *postgres.Client, pubSub *pubsub.Client) *v1Router {
	return &v1Router{
		handlers: v1.NewRouteHandler(mongo, postgres, pubSub),
		mongo:    mongo,
		postgres: postgres,
		pubSub:   pubSub,
	}
}

func (r *v1Router) SetupRoutes(route *gin.RouterGroup) {
	v1Route := route.Group("/v1")

	v1Route.GET("/mongo", r.handlers.GetMongo)
	v1Route.GET("/postgres", r.handlers.GetPostgres)
	v1Route.GET("/pubsub", r.handlers.GetPubSub)
}
