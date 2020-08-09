package model

import (
	"log"
	"math/rand"
	"time"

	"github.com/alejoberardino/minesweeper/utils"
)

const (
	BLANK   = 0
	UNKNOWN = -1
	FLAGGED = -2
	MINE    = -3
)

type Cell struct {
	State   int
	Clicked bool
}

type Game struct {
	Rows       int
	Columns    int
	Mines      int
	State      string
	Started_at time.Time
	Cells      [][]Cell // Could have also stored it as a simple array, this is nicer (maybe)
}

func BuildGame(rows int, columns int, mines int) Game {
	log.Printf("Creating game with %d rows, %d columns and %d mines", rows, columns, mines)
	var game Game = Game{
		Rows:       rows,
		Columns:    columns,
		Mines:      mines,
		State:      "new",
		Started_at: time.Now(),
		Cells:      make([][]Cell, rows),
	}

	log.Print("Initializing board")     // (there should be a cleaner way, no?)
	cells := make([]Cell, rows*columns) // Allocate all to ensure it's contiguous in memory
	for y := range game.Cells {
		game.Cells[y], cells = cells[:columns], cells[columns:]
		for x := range game.Cells[y] {
			game.Cells[y][x] = Cell{
				State:   0,
				Clicked: false,
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
		x := rand.Intn(rows)

		if game.Cells[y][x].State != MINE {
			log.Printf("Added mine at (%d;%d)", x, y)
			game.Cells[y][x].State = MINE
		} else {
			log.Printf("(%d;%d) was already a mine (%d). Retrying...", x, y, game.Cells[y][x].State)
			i-- // Retry
		}
		counter++
	}

	log.Print("Calculating mine-adjacent scores")
	for y := range game.Cells {
		for x := range game.Cells[y] {
			if game.Cells[y][x].State == MINE {
				game.calculateAdjacentScores(x, y)
			}
		}
	}

	return game
}

func (game *Game) calculateAdjacentScores(x int, y int) {
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

			log.Printf("Increasing score for (%d;%d)", j, i)
			game.Cells[i][j].State++
		}
	}
}
