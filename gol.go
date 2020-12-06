package main

import (
	tl "github.com/JoelOtter/termloop"
)

const cellWidth = 5
const cellHeight = 2

type Cell struct {
	*tl.Rectangle
	alive bool
	game  *tl.Game
}

type GameOfLife struct {
	*tl.Entity
	width   int
	height  int
	board   [100][100]bool
	running bool
	ticks int
}

func mod(a, b int) int {
	m := a % b
	if a < 0 && b < 0 {
		m -= b
	}
	if a < 0 && b > 0 {
		m += b
	}
	return m
}

func (gol *GameOfLife) numberOfNeighbours(xPos int, yPos int) int {
	neighbors := 0
	nums := []int{-1, 0, 1}
	for _, x := range nums {
		for _, y := range nums {
			if x == 0 && y == 0 {
				continue
			}
			if gol.board[mod(xPos+x, 100)][mod(yPos+y, 100)] {
				neighbors++
			}
		}
	}
	return neighbors
}

func (gol *GameOfLife) iterate() {
	var new_board = [100][100]bool{}
	for x, y_arr := range gol.board {
		for y, cell := range y_arr {
			numNeighbours := gol.numberOfNeighbours(x, y)
			if cell && (numNeighbours == 2 || numNeighbours == 3) {
				new_board[x][y] = true
			}
			if !cell && numNeighbours == 3 {
				new_board[x][y] = true
			}
		}
	}
	gol.board = new_board
}

func (gol *GameOfLife) Tick(event tl.Event) {
	if event.Type == tl.EventMouse {
		board.board[event.MouseX/cell_width][event.MouseY/cell_height] = true
	} else if event.Type == tl.EventKey {
		if event.Key == tl.KeySpace {
			gol.running = false
		}
		if event.Key == tl.KeyEnter {
			gol.running = true
		}
	}
	if gol.running {
		if gol.ticks == 45 {
			gol.iterate()
			gol.ticks = 0
		} else {
			gol.ticks++
		}
	}
}

func (cell *Cell) Tick(event tl.Event) {
	x, y := cell.Position()

	x = x / cell_width
	y = y / cell_height

	if board.board[x][y] {
		cell.Rectangle.SetColor(tl.ColorRed)
	} else {
		cell.Rectangle.SetColor(tl.ColorBlack)
	}
}

var board = GameOfLife{tl.NewEntity(1, 1, 1, 1), 100, 100, [100][100]bool{}, true, 0}

const cell_width = 5
const cell_height = 2

func main() {
	game := tl.NewGame()
	game.Screen().SetFps(1000)

	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
		Fg: tl.ColorBlack,
		Ch: ' ',
	})

	// cell1 := Cell{tl.NewRectangle(20, 20, cellWidth, cellHeight, tl.ColorRed), true}
	// cell2 := Cell{tl.NewRectangle(20, 22, cellWidth, cellHeight, tl.ColorRed), true}
	// cell3 := Cell{tl.NewRectangle(20, 24, cellWidth, cellHeight, tl.ColorRed), true}

	// cell4 := Cell{tl.NewRectangle(15, 22, cellWidth, cellHeight, tl.ColorRed), false}
	// cell5 := Cell{tl.NewRectangle(25, 22, cellWidth, cellHeight, tl.ColorRed), false}

	// board.board[4][10] = true
	// board.board[4][11] = true
	// board.board[4][12] = true

	// board.board[7][10] = true
	// board.board[8][10] = true
	// board.board[9][10] = true
	// board.board[9][9] = true
	// board.board[7][8] = true

	// level.AddEntity(&cell1)
	// level.AddEntity(&cell2)
	// level.AddEntity(&cell3)
	// level.AddEntity(&cell4)
	// level.AddEntity(&cell5)
	level.AddEntity(&board)

	for x, yArr := range board.board {
		for y, alive := range yArr {
			level.AddEntity(&Cell{tl.NewRectangle(x*cell_width, y*cell_height, cellWidth, cellHeight, tl.ColorRed), alive, game})
		}
	}

	game.Screen().SetLevel(level)

	game.Start()
	// fmt.Println("============================")
	// fmt.Println(board.board[3][10], board.board[4][10], board.board[5][10])
	// fmt.Println(board.board[3][11], board.board[4][11], board.board[5][11])
	// fmt.Println(board.board[3][12], board.board[4][12], board.board[5][12])
	// fmt.Println("============================")
	// board.iterate()
	// fmt.Println(board.board[3][10], board.board[4][10], board.board[5][10])
	// fmt.Println(board.board[3][11], board.board[4][11], board.board[5][11])
	// fmt.Println(board.board[3][12], board.board[4][12], board.board[5][12])
	// fmt.Println("============================")
	// board.iterate()
	// fmt.Println(board.board[3][10], board.board[4][10], board.board[5][10])
	// fmt.Println(board.board[3][11], board.board[4][11], board.board[5][11])
	// fmt.Println(board.board[3][12], board.board[4][12], board.board[5][12])
	// fmt.Println("============================")
}
