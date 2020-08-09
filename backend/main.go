package main

import (
	"fmt"
	"os"

	"github.com/alejoberardino/minesweeper/docs"
	"github.com/alejoberardino/minesweeper/routes"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"   // gin-swagger middleware
	"github.com/swaggo/gin-swagger/swaggerFiles" // swagger embed files
)

func main() {
	// Create new Gin REST API
	r := gin.Default()

	// Games
	games := r.Group("/games")
	{
		games.POST("/", routes.CreateGame)
	}

	// Declarative swagger info
	docs.SwaggerInfo.Title = "MinesweeperAPI"
	docs.SwaggerInfo.Description = "This is a test project implementing a Minesweeper game, using a REST API built with Go"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", os.Getenv("API_HOST"), os.Getenv("API_PORT"))
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// Add sample path
	r.GET("/test", routes.Test)

	// Run swagger docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
