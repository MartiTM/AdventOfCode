package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/MartiTM/AdventOfCode2022/util"
)

func getPossibleGameId(dataPerGame []string, maxRed int, maxGreen int, maxBlue int) []int {
	result := []int{}
	regexGetByColor := regexp.MustCompile(`([0-9]+)[ ]?([a-z]+)`)
	gameLoop:
	for id, game := range dataPerGame {
		setResultByColor := regexGetByColor.FindAllStringSubmatch(game, -1)
		for _, setResult := range setResultByColor {
			number, _ := strconv.Atoi(setResult[1])
			color := setResult [2]
			switch color{
			case "red":
				if number > maxRed {
					continue gameLoop
				}
			case "green":
				if number > maxGreen {
					continue gameLoop
				}
			case "blue":
				if number > maxBlue {
					continue gameLoop
				}
			}
		}
		result = append(result, id+1)
	}
	return result
}


func part1() {
	data := util.GetStringDataFromFile("./2023/day_2/input")
	dataPerGame := strings.Split(data, "\n")
	possibleGameId := getPossibleGameId(dataPerGame, 12, 13, 14)
	total := 0
	for _, id := range possibleGameId {
		total+=id
	}
	fmt.Printf("%v\n", total)
}

func part2Res(dataPerGame []string) int {
	result := 0
	regexGetByColor := regexp.MustCompile(`([0-9]+)[ ]?([a-z]+)`)
	for _, game := range dataPerGame {
		setResultByColor := regexGetByColor.FindAllStringSubmatch(game, -1)
		maxRed := 0
		maxGreen := 0
		maxBlue := 0
		for _, setResult := range setResultByColor {
			number, _ := strconv.Atoi(setResult[1])
			color := setResult [2]
			switch color{
			case "red":
				if number > maxRed {
					maxRed = number
				}
			case "green":
				if number > maxGreen {
					maxGreen = number
				}
			case "blue":
				if number > maxBlue {
					maxBlue = number
				}
			}
		}
		result+= (maxBlue * maxRed * maxGreen)
	}
	return result
}

func part2() {
	data := util.GetStringDataFromFile("./2023/day_2/input")
	dataPerGame := strings.Split(data, "\n")
	total := part2Res(dataPerGame)
	fmt.Printf("%v\n", total)
}

func main() {
	part2()
}