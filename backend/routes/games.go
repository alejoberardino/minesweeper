package routes

import (
	"log"
	"net/http"
	"os"

	"github.com/alejoberardino/minesweeper/utils"
	"github.com/gin-gonic/gin"
)

type CreateGameDTO struct {
	Rows    int `json:"rows"`
	Columns int `json:"columns"`
	Mines   int `json:"mines"`
}

// Create game godoc
// @Summary Create game endpoint
// @Description Create a new game
// @Produce  json
// @Param dto body CreateGameDTO true "Details for the new game"
// @Success 200 {object} model.MessageResponse
// @Router /games/ [post]
func CreateGame(c *gin.Context) {
	var dto CreateGameDTO

	// Parse dto
	if err := c.ShouldBindJSON(&dto); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}

	// Connect to mongo
	client, ctx, cancel := utils.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	// Insert to database
	result, err := client.Database(os.Getenv("MONGO_DATABASE")).Collection("games").InsertOne(ctx, dto)
	if err != nil {
		log.Printf("Could not create game: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}

	c.JSON(200, result)
}
