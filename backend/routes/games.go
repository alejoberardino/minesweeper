package routes

import (
	"log"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/alejoberardino/minesweeper/model"
	"github.com/alejoberardino/minesweeper/utils"
	"github.com/gin-gonic/gin"
)

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
func CreateGame(c *gin.Context) {
	var dto CreateGameRequestDTO

	// Parse dto
	if err := c.ShouldBindJSON(&dto); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}

	game := model.BuildGame(dto.Rows, dto.Columns, dto.Mines)
	log.Printf("Built game %v", game)

	// Connect to mongo
	client, ctx, cancel := utils.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	// Insert to database
	result, err := client.Database(os.Getenv("MONGO_DATABASE")).Collection("games").InsertOne(ctx, game)
	if err != nil {
		log.Printf("Could not create game: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}

	// Return the id of the new game
	c.JSON(200, gin.H{
		"id": result.InsertedID,
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
func GetGame(c *gin.Context) {
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

	log.Printf("Object id: %v", objectId)

	var game GetGameResponseDTO

	// Connect to mongo
	client, ctx, cancel := utils.GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)

	// Get from the database
	err = client.Database(os.Getenv("MONGO_DATABASE")).Collection("games").FindOne(ctx, bson.M{"_id": objectId}).Decode(&game)
	if err != nil || game.Id == primitive.NilObjectID {
		log.Printf("Could not find game with id %s", id)
		c.JSON(http.StatusNotFound, gin.H{"msg": "Could not find game"})
		return
	}
	log.Printf("Game: %v", game)

	// Return the id of the new game
	c.JSON(200, game)
}
