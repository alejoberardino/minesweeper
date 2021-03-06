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
	game.PrintBoard()

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

	log.Print("Geting game from the database")
	err := service.Client.Database(os.Getenv("MONGO_DATABASE")).Collection("games").FindOne(service.Context, bson.M{"_id": id}).Decode(&game)

	return game, err
}

func (service *GameService) Click(id primitive.ObjectID, x int, y int, state int) (model.Game, error) {
	// TODO: This is very inefficient, but mongo deep queries require aggregation & projections,
	// and the mongo-go-driver isn't super documented. I'll leave it for later.
	game, err := service.Get(id)
	if err != nil || game.Id == primitive.NilObjectID {
		return game, err
	}
	cell := &game.Cells[y][x]
	log.Printf("Got game, cell: %v", cell)

	log.Printf("Modifying local instance, setting state %d in cell (%d;%d)", state, x, y)
	cell.State = state
	game.Value = "in_progress"

	if cell.State == model.CLICKED {
		if cell.Value == model.MINE {
			log.Print("Clicked on a mine, game over!")
			game.Value = "over"
			game.UncoverAll()
		} else if cell.Value == model.BLANK {
			log.Print("Clicked on a blank cell, need to traverse and uncover all neighboring blank cells")
			game.UncoverBlank(x, y)
		} else if game.CheckWin() {
			log.Print("All mines were sweeped! Congrats!")
			game.Value = "complete"
		}
	}

	log.Print("Updating state in the db")
	if result, err := service.Client.Database(os.Getenv("MONGO_DATABASE")).Collection("games").UpdateOne(service.Context, bson.M{"_id": id}, bson.D{
		{
			"$set", bson.D{
				{
					"cells", game.Cells,
				},
			},
		},
	}); err != nil || result.ModifiedCount > 1 {
		return game, err
	}

	log.Print("Update successful")
	return game, err
}
