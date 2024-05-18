package v1

import (
	"fx_101/mongo"
	"fx_101/postgres"
	"fx_101/pubsub"

	"github.com/gin-gonic/gin"
)

type RouteHandler struct {
	mongoClient    *mongo.Client
	postgresClient *postgres.Client
	pubsubClient   *pubsub.Client
}

func NewRouteHandler(mongoClient *mongo.Client, postgresClient *postgres.Client, pubsubClient *pubsub.Client) *RouteHandler {
	return &RouteHandler{
		mongoClient:    mongoClient,
		postgresClient: postgresClient,
		pubsubClient:   pubsubClient,
	}
}

func (h *RouteHandler) GetMongo(ctx *gin.Context) {
	h.mongoClient.Ping()
}

func (h *RouteHandler) GetPostgres(ctx *gin.Context) {
	h.postgresClient.Ping()
}

func (h *RouteHandler) GetPubSub(ctx *gin.Context) {
	h.pubsubClient.Publish("hello")
}
