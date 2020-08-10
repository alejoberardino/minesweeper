package model

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/alejoberardino/minesweeper/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// States
	CLICKED  = 1
	UNKNOWN  = 0
	FLAGGED  = -1
	POSSIBLE = -2

	// Special values
	BLANK = 0
	MINE  = -1
)

type Cell struct {
	Value int `json:"value" bson:"value"`
	State int `json:"state" bson:"state"`
}

type Game struct {
	Id        primitive.ObjectID `bson:"_id"`
	Rows      int                `bson:"rows"`
	Columns   int                `bson:"columns"`
	Mines     int                `bson:"mines"`
	Value     string             `bson:"value"`
	StartedAt time.Time          `bson:"started_at"`
	Cells     [][]Cell           `bson:"cells"` // Could have also stored it as a simple array, this is nicer (maybe)
}

func BuildGame(rows int, columns int, mines int) Game {
	log.Printf("Creating game with %d rows, %d columns and %d mines", rows, columns, mines)
	var game Game = Game{
		Id:        primitive.NewObjectID(),
		Rows:      rows,
		Columns:   columns,
		Mines:     mines,
		Value:     "new",
		StartedAt: time.Now(),
		Cells:     make([][]Cell, rows),
	}

	log.Print("Initializing board")     // (there should be a cleaner way, no?)
	cells := make([]Cell, rows*columns) // Allocate all to ensure it's contiguous in memory
	for y := range game.Cells {
		game.Cells[y], cells = cells[:columns], cells[columns:]
		for x := range game.Cells[y] {
			game.Cells[y][x] = Cell{
				Value: 0,
				State: 0,
			}
		}
	}

	log.Print("Adding random mines")
	counter := 0
	for i := 0; i < mines; i++ {
		if counter > 10 {
			break
		}
		log.Printf("Adding mine no. %d", i)
		y := rand.Intn(rows)
		x := rand.Intn(columns)

		if game.Cells[y][x].Value != MINE {
			log.Printf("Added mine at (%d;%d)", x, y)
			game.Cells[y][x].Value = MINE
		} else {
			log.Printf("(%d;%d) was already a mine (%d). Retrying...", x, y, game.Cells[y][x].Value)
			i-- // Retry
		}
		counter++
	}

	log.Print("Calculating mine-adjacent scores")
	for y := range game.Cells {
		for x := range game.Cells[y] {
			if game.Cells[y][x].Value == MINE {
				game.calculateAdjacentScores(x, y)
			}
		}
	}

	return game
}

func (game *Game) traverseAdjacent(x int, y int, f func(x int, y int)) {
	log.Printf("Calculating adjacent Scores for (%d;%d)", x, y)
	startRow := utils.Max(0, y-1)
	endRow := utils.Min(y+1, game.Rows-1)
	log.Printf("Looping rows from %d to %d", startRow, endRow)
	startCol := utils.Max(0, x-1)
	endCol := utils.Min(x+1, game.Columns-1)
	log.Printf("Looping cols from %d to %d", startCol, endCol)
	for i := startRow; i <= endRow; i++ {
		for j := startCol; j <= endCol; j++ {
			if i == y && j == x {
				continue
			}
			f(j, i)
		}
	}
}

func (game *Game) calculateAdjacentScores(x0 int, y0 int) {
	log.Printf("Calculating adjacent Scores for (%d;%d)", x0, y0)

	game.traverseAdjacent(x0, y0, func(x int, y int) {
		log.Printf("Increasing score for (%d;%d)", x, y)
		if game.Cells[y][x].Value != MINE {
			game.Cells[y][x].Value++
		}
	})
}

func (game *Game) UncoverBlank(x0 int, y0 int) {
	log.Printf("Uncovering blank area starting at (%d;%d)", x0, y0)
	game.Cells[y0][x0].State = CLICKED // We have to do this first to have an escape from infinite recursion

	game.traverseAdjacent(x0, y0, func(x int, y int) {
		cell := &game.Cells[y][x]

		// We want only orthogonal matches here
		if cell.Value == BLANK && cell.State != CLICKED {
			log.Printf("Uncover (recurse) (%d;%d)", x, y)
			game.UncoverBlank(x, y)
		}

		cell.State = CLICKED
	})
}

func (game *Game) UncoverAll() {
	log.Print("Uncovering all cells")
	for y := range game.Cells {
		for x := range game.Cells[y] {
			game.Cells[y][x].State = CLICKED
		}
	}
}

func (game *Game) PrintBoard() {
	log.Print("Game board:")
	var sb strings.Builder
	sb.WriteString("\n")
	for y := range game.Cells {
		for x := range game.Cells[y] {
			sb.WriteString(fmt.Sprintf("%2d ", game.Cells[y][x].Value))
		}
		sb.WriteString("\n")
	}
	log.Print(sb.String())
}
