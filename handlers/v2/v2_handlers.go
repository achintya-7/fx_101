package v2

import (
	"fx_101/mongo"

	"github.com/gin-gonic/gin"
)

type RouteHandler struct {
	mongoClient *mongo.Client
}

func NewRouteHandler(mongoClient *mongo.Client) *RouteHandler {
	return &RouteHandler{
		mongoClient: mongoClient,
	}
}

func (h *RouteHandler) GetMongo(ctx *gin.Context) {
	h.mongoClient.Ping()
}
