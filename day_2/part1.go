package main

import (
	"bytes"
	"fmt"
	"os"
)

type Game struct {
	opposant 	Shape
	me 			Shape
}

type Shape interface {
	getType() string
	getShapeValue() int
	getWinningPoint(Shape) int
}


type Rock struct {}

func (r Rock) getType() string {
	return "rock"
}

func (r Rock) getShapeValue() int {
	return 1
}

func (r Rock) getWinningPoint(s Shape) int {
	if r.getType() == s.getType() {
		return 3
	}
	if s.getType() == new(Paper).getType() {
		return 0
	}
	return 6
}

type Paper struct {}

func (p Paper) getType() string {
	return "paper"
}

func (p Paper) getShapeValue() int {
	return 2
}

func (p Paper) getWinningPoint(s Shape) int {
	if p.getType() == s.getType() {
		return 3
	}
	if s.getType() == new(Scissors).getType() {
		return 0
	}
	return 6
}

type Scissors struct {}

func (sc Scissors) getType() string {
	return "scissors"
}

func (sc Scissors) getShapeValue() int {
	return 3
}

func (sc Scissors) getWinningPoint(s Shape) int {
	if sc.getType() == s.getType() {
		return 3
	}
	if s.getType() == new(Rock).getType() {
		return 0
	}
	return 6
}

func (g Game) getMyScore() int {
	return g.me.getWinningPoint(g.opposant) + g.me.getShapeValue()
}

func main() {
	rawData, err := getRawData()
	checkErr(err)
	gameList := getListOfGames(rawData)
	score := sumMyScores(gameList)
	fmt.Printf("My score is %v\n", score)
}

func getRawData() ([]byte, error) {
	rawData, err := os.ReadFile("./day_2/input")
	return rawData, err
}

func getListOfGames(rawData []byte) []Game {

	rowList := bytes.Split(rawData, []byte("\n"))

	gameList := []Game{}

	// chaque ligne contient 3 bit -> jeux de l'opposant + espace + notre jeux
	for _, row := range rowList {
		var me Shape
		var opposant Shape

		switch string(row[0]) {
		case "A":
			opposant = new(Rock)
		case "B":
			opposant = new(Paper)
		case "C":
			opposant = new(Scissors)
		}

		switch string(row[2]) {
		case "X":
			me = new(Rock)
		case "Y":
			me = new(Paper)
		case "Z":
			me = new(Scissors)
		}

		gameList = append(gameList, Game{opposant, me})
	}

	return gameList
}

func sumMyScores(list []Game) int {
	sum := 0
	for _, game := range list {
		sum += game.getMyScore()
	}
	return sum
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}