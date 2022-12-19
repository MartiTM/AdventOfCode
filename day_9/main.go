package main

import (
	"bytes"
	"fmt"

	"github.com/MartiTM/AdventOfCode2022/util"
)

func main() {
	rawData := util.GetRawData("./day_9/input")
	game := newGame(1000, 10)
	movements := getMovements(rawData)
	game.play(movements)
	score := game.getMarks()
	fmt.Printf("Score : %v\n", score)
}

func getMovements(data []byte) []Movement {
	movements := []Movement{}
	for _, move := range bytes.Split(data, []byte("\n")) {
		newMove := Movement{0, 0}
		switch move[0] {
		case byte('U'):
			newMove.y--
		case byte('D'):
			newMove.y++
		case byte('R'):
			newMove.x++
		case byte('L'):
			newMove.x--
		default:
			panic("incorrect movement")
		}
		for i:=0; i<util.BytesStringToInt(move[2:]); i++ {
			movements = append(movements, newMove)
		}

	}
	return movements
}

func newGame(boardSize int, ropeSize int) *Game {
	game := new(Game)
	game.board = make([][]byte, boardSize)
	for i := range game.board {
		game.board[i] = make([]byte, boardSize)
	}
	game.rope = []Position{}
	for i := 0; i < ropeSize; i++ {
		game.rope = append(game.rope, game.getStartingPosition())
	}
	game.markTail()

	return game
}

type Game struct {
	board [][]byte
	rope []Position
}

func (g Game) getMarks() int {
	sum := 0
	for _, rows := range g.board {
		for _, b := range rows {
			if b == 1 {
				sum++
			}
		}
	}
	return sum
}

func (g Game) getStartingPosition() Position {
	return Position{len(g.board)/2,len(g.board)/2}
}

func (g *Game) play(movements []Movement) {
	for _, move := range movements {
		g.applyMovement(move)
		g.markTail()
	}
}

func (g *Game) applyMovement(move Movement) {	
	g.rope[0].x+=move.x
	g.rope[0].y+=move.y
	
	if g.rope[0].x >= len(g.board) || g.rope[0].y >= len(g.board) || g.rope[0].x < 0 || g.rope[0].y < 0 {
		panic(fmt.Errorf("Out of the board. Head x: %v, y: %v. Move: x: %v, y: %v", g.rope[0].x, g.rope[0].y, move.x, move.y))
	}
	
	for i := 1; i < len(g.rope); i++ {
		if g.rope[i].isClose(g.rope[i-1]) {
			return
		}
		addX := g.rope[i-1].x - g.rope[i].x
		addY := g.rope[i-1].y - g.rope[i].y
		switch addX {
		case 2:
			addX = 1
		case -2:
			addX = -1
		}
		switch addY {
		case 2:
			addY = 1
		case -2:
			addY = -1
		}
		g.rope[i].x += addX
		g.rope[i].y += addY
	}
}
// func (g *Game) applyMovement(move Movement) {
// 	oldPos := Position{g.rope[0].x, g.rope[0].y}
	
// 	g.rope[0].x+=move.x
// 	g.rope[0].y+=move.y
	
// 	if g.rope[0].x >= len(g.board) || g.rope[0].y >= len(g.board) || g.rope[0].x < 0 || g.rope[0].y < 0 {
// 		panic(fmt.Errorf("Out of the board. Head x: %v, y: %v. Move: x: %v, y: %v", g.rope[0].x, g.rope[0].y, move.x, move.y))
// 	}
	
// 	for i := 1; i < len(g.rope); i++ {
// 		if g.rope[i].isClose(g.rope[i-1]) {
// 			return
// 		}
// 		temp := Position{g.rope[i].x, g.rope[i].y}
// 		g.rope[i] = Position{oldPos.x, oldPos.y}
// 		oldPos = Position{temp.x, temp.y}
// 	}
// }

func (g Game) displayBoard() {
	for y, row := range g.board {
		for x, pos := range row {
			if g.getStartingPosition().x == x && g.getStartingPosition().y == y {
				fmt.Print("s")
				continue
			}
			found:=false
			for i, pos := range g.rope {
				if pos.x == x && pos.y == y {
					found=true
					if i == 0 {
						fmt.Print("H")
						break
					}
					if i == len(g.rope) {
						fmt.Print("T")
						break
					}
					fmt.Printf("%v", i)
					break
				}
			}
			if found {
				continue
			}
			if pos == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}

func (g *Game) markTail() {
	x:= g.rope[len(g.rope)-1].x
	y:= g.rope[len(g.rope)-1].y
	(*g).board[y][x] = 1 
}

type Position struct {
	x int
	y int
}

func (main Position) isClose(target Position) bool {
	if (main.x == target.x || main.x == target.x+1 || main.x == target.x-1) && (main.y == target.y || main.y == target.y+1 || main.y == target.y-1) {
		return true
	}
	return false
}

type Movement struct {
	x int
	y int
}