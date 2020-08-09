package controllers

import (
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/alejoberardino/minesweeper/services"
	"github.com/gin-gonic/gin"
)

type GameController struct {
	GameService *services.GameService
}

type CreateGameRequestDTO struct {
	Rows    int `json:"rows" example:"10"`
	Columns int `json:"columns" example:"9"`
	Mines   int `json:"mines" example:"8"`
}

type CreateGameResponseDTO struct {
	Id int `json:"id"`
}

// Create game godoc
// @Summary Create game endpoint
// @Description Create a new game
// @Produce  json
// @Param dto body CreateGameRequestDTO true "Details for the new game"
// @Success 200 {object} CreateGameResponseDTO
// @Router /games/ [post]
func (controller *GameController) Create(c *gin.Context) {
	var dto CreateGameRequestDTO

	// Parse dto
	if err := c.ShouldBindJSON(&dto); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}

	// Create game
	id, err := controller.GameService.Create(dto.Rows, dto.Columns, dto.Mines)
	if err != nil {
		log.Printf("Could not create game: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}

	// Return the id of the new game
	c.JSON(200, gin.H{
		"id": id,
	})
}

type NestedCellDTO struct {
	Value int `json:"value" bson:"value"`
	State int `json:"state" bson:"state"`
}

type GetGameResponseDTO struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	Rows      int                `json:"rows" bson:"rows"`
	Columns   int                `json:"columns" bson:"columns"`
	Mines     int                `json:"mines" bson:"mines"`
	Value     string             `json:"value" bson:"value"`
	StartedAt time.Time          `json:"startedAt" bson:"started_at"`
	Cells     [][]NestedCellDTO  `json:"cells" bson:"cells"`
}

// Get game godoc
// @Summary Get game endpoint
// @Description Gets an existing game from the db
// @Produce  json
// @Param id path string true "Id of the game to get"
// @Success 200 {object} GetGameResponseDTO
// @Router /games/{id} [get]
func (controller *GameController) Get(c *gin.Context) {
	id := c.Param("id")

	// if id == nil {
	// 	log.Print("Id was not present in the request")
	// 	c.JSON(http.StatusBadRequest, gin.H{"msg": "Id was not present in the request"})
	// 	return
	// }
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Print("The provided id was invalid")
		c.JSON(http.StatusBadRequest, gin.H{"msg": "The provided id was invalid"})
	}

	log.Printf("Got Object id: %v", objectId)
	game, err := controller.GameService.Get(objectId)
	if err != nil {
		log.Printf("There was an error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "There was an error"})
	} else if game.Id == primitive.NilObjectID {
		log.Printf("Could not find game with id %s", id)
		c.JSON(http.StatusNotFound, gin.H{"msg": "Could not find game with provided ID"})
	} else {
		log.Printf("Found Game: %v", game)
		c.JSON(200, game)
	}
}