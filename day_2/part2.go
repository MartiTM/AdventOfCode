package main

import (
	"bytes"
	"fmt"
	"os"
)

type Game struct {
	opposant 	Shape
	result 		Result
}
func (g Game) getMyShape() Shape {
	switch g.result.getResult() {
	case new(Win).getResult():
		switch g.opposant.getType() {
		case new(Rock).getType():
			return new(Paper)
		case new(Paper).getType():
			return new(Scissors)
		case new(Scissors).getType():
			return new(Rock)
		}
	case new(Lose).getResult():
		switch g.opposant.getType() {
		case new(Rock).getType():
			return new(Scissors)
		case new(Paper).getType():
			return new(Rock)
		case new(Scissors).getType():
			return new(Paper)
		}
	}
	
	return g.opposant
}
func (g Game) getMyScore() int { return g.result.getResultPoint() + g.getMyShape().getShapeValue() }

type Shape interface {
	getType() string
	getShapeValue() int
	getWinningPoint(Shape) int
}

type Result interface {
	getResult() string
	getResultPoint() int
}

type Win struct {}
func (w Win) getResult() 		string 	{ return "win" }
func (w Win) getResultPoint() 	int 	{ return 6 }

type Draw struct {}
func (d Draw) getResult() 		string 	{ return "draw" }
func (d Draw) getResultPoint() 	int 	{ return 3 }

type Lose struct {}
func (d Lose) getResult() 		string 	{ return "lose" }
func (d Lose) getResultPoint() 	int 	{ return 0 }

type Rock struct {}
func (r Rock) getType() string { return "rock" }
func (r Rock) getShapeValue() int { return 1 }
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
func (p Paper) getType() string {return "paper" }
func (p Paper) getShapeValue() int { return 2 }
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
func (sc Scissors) getType() string { return "scissors" }
func (sc Scissors) getShapeValue() int { return 3 }
func (sc Scissors) getWinningPoint(s Shape) int {
	if sc.getType() == s.getType() {
		return 3
	}
	if s.getType() == new(Rock).getType() {
		return 0
	}
	return 6
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
		var result Result
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
			result = new(Lose)
		case "Y":
			result = new(Draw)
		case "Z":
			result = new(Win)
		}

		gameList = append(gameList, Game{opposant, result})
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