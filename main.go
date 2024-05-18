package main

import (
	"fx_101/config"
	"fx_101/controllers"
	"fx_101/mongo"
	"fx_101/postgres"
	"fx_101/pubsub"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router   *gin.Engine
	mongo    *mongo.Client
	postgres *postgres.Client
	pubSub   *pubsub.Client
}

func NewServer() *Server {
	config := config.NewConfig()

	return &Server{
		mongo:    mongo.NewMongo(config.MongoUrl),
		postgres: postgres.NewPostgres(config.PostgresUrl),
		pubSub:   pubsub.NewPubSub(config.Topic),
	}
}

func (s *Server) Run() {
	s.setupRoutes()

	s.router.Run(":8080")
}

func (s *Server) setupRoutes() {
	router := gin.Default()

	baseRouter := router.Group("/")

	v1Router := controllers.NewRouter(s.mongo, s.postgres, s.pubSub)
	v1Router.SetupRoutes(baseRouter)

	s.router = router
}

func main() {
	server := NewServer()
	server.Run()
}