package main

import (
	"fmt"
	"log"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/aman-lf/event-management/config"
	"github.com/aman-lf/event-management/database"
	"github.com/aman-lf/event-management/graph"
	"github.com/aman-lf/event-management/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatal("failed to get SQL database connection:", err)
		}
		sqlDB.Close()
	}()

	gin.SetMode(config.Cfg.Mode)

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.JSONLoggerMiddleware())
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.Auth())

	addAPIRoutes(router)
	router.Run(fmt.Sprintf("%s:%s", config.Cfg.Host, config.Cfg.Port))
}

func addAPIRoutes(router *gin.Engine) {
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	router.POST("/query", gin.WrapH(srv))

	// Playground endpoint
	router.GET("/", gin.WrapH(playground.Handler("GraphQL Playground", "/query")))
}
