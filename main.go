package main

import (
	"fx_101/mongo"
	"fx_101/postgres"
	"fx_101/pubsub"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Topic       string
	MongoUrl    string
	PostgresUrl string
}

func NewConfig() *Config {
	return &Config{
		Topic:       "my-topic",
		MongoUrl:    "mongodb://localhost:27017",
		PostgresUrl: "postgres://localhost:5432",
	}
}

type Server struct {
	router   *gin.Engine
	mongo    *mongo.Client
	postgres *postgres.Client
	pubSub   *pubsub.Client
}

func NewServer() *Server {
	config := NewConfig()

	return &Server{
		router:   gin.Default(),
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
	
}
