package main

import (
	"context"
	"fx_101/config"
	"fx_101/controllers"
	v1 "fx_101/handlers/v1"
	v2 "fx_101/handlers/v2"
	"fx_101/mongo"
	"fx_101/postgres"
	"fx_101/pubsub"
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Server struct {
	router *gin.Engine
}

func NewServer(lc fx.Lifecycle, apiRouter *ApiRouter) *Server {

	server := &Server{
		router: apiRouter.Router,
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			log.Println("Starting server")
			go server.Run()
			return nil
		},
		OnStop: func(context.Context) error {
			log.Println("Shutting down server")
			return nil
		},
	})

	return server
}

func (s *Server) Run() {
	s.router.Run(":8080")
}

type ApiRouter struct {
	Router *gin.Engine
}

func SetupRoutes(v1Router *controllers.V1Router, v2Router *controllers.V2Router) *ApiRouter {
	router := gin.Default()

	baseRouter := router.Group("/")

	v1Router.SetupRoutes(baseRouter)
	v2Router.SetupRoutes(baseRouter)

	return &ApiRouter{Router: router}
}

func main() {
	fx.New(
		fx.Provide(
			config.NewConfig,
			mongo.NewMongo,
			postgres.NewPostgres,
			pubsub.NewPubSub,
			controllers.NewV1Router,
			controllers.NewV2Router,
			SetupRoutes,
			v1.NewRouteHandler,
			v2.NewRouteHandler,
		),
		fx.Invoke(NewServer),
	).Run()
}
