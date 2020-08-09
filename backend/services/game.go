package services

import (
	"context"
	"log"
	"os"

	"github.com/alejoberardino/minesweeper/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type GameService struct {
	Client  *mongo.Client
	Context context.Context
}

func (service *GameService) Create(rows int, columns int, mines int) (primitive.ObjectID, error) {
	game := model.BuildGame(rows, columns, mines)
	log.Printf("Built game %v", game)

	// Insert to database
	result, err := service.Client.Database(os.Getenv("MONGO_DATABASE")).Collection("games").InsertOne(service.Context, game)
	if err != nil {
		log.Printf("Could not create game: %v", err)
		return primitive.NilObjectID, err
	}

	return result.InsertedID.(primitive.ObjectID), nil
}

func (service *GameService) Get(id primitive.ObjectID) (model.Game, error) {
	var game model.Game

	// Get from the database
	err := service.Client.Database(os.Getenv("MONGO_DATABASE")).Collection("games").FindOne(service.Context, bson.M{"_id": id}).Decode(&game)

	return game, err
}
