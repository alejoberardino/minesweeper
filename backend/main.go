package main

import (
	"fmt"
	"os"

	"github.com/alejoberardino/minesweeper/controllers"
	"github.com/alejoberardino/minesweeper/docs"
	"github.com/alejoberardino/minesweeper/services"
	"github.com/alejoberardino/minesweeper/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"   // gin-swagger middleware
	"github.com/swaggo/gin-swagger/swaggerFiles" // swagger embed files
)

func main() {
	// Create new Gin REST API
	router := gin.Default()
	r := router.Group("/api")

	// Configure CORS
	r.Use(cors.Default())

	// Connect to mongo
	client, ctx, cancel := utils.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	// Games
	gameController := controllers.GameController{
		GameService: &services.GameService{
			Client:  client,
			Context: ctx,
		},
	}
	games := r.Group("/games")
	{
		games.POST("/", gameController.Create)
		games.GET("/:id", gameController.Get)
		games.POST("/:id/click", gameController.Click)
	}

	// Declarative swagger info
	docs.SwaggerInfo.Title = "MinesweeperAPI"
	docs.SwaggerInfo.Description = "This is a test project implementing a Minesweeper game, using a REST API built with Go"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", os.Getenv("API_HOST"), os.Getenv("API_PORT"))
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// Run swagger docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Run the api
	router.Run(":80")
}
