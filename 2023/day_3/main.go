package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"

	"github.com/MartiTM/AdventOfCode2022/util"
)

type Coord struct {
	y int
	x int
	value string
}

// Only work if we use engine part on base
func (base Coord) isClose(target Coord) bool {
	return base.y - 1 <= target.y && target.y <= base.y + 1 && base.x - 1 <= target.x && target.x <= (base.x + len(base.value))
}

func main() {
	data := util.GetRawData("./2023/day_3/inputTest")
	engineMap := bytes.Split(data, []byte("\n"))

	engineParts, symbols := getEngineAndSymbols(engineMap)

	enginePartsClose := getEnginePartsCloseToSymbol(engineParts, symbols)

	total := getTotalEngineParts(enginePartsClose)

	fmt.Printf("%v\n", total)
	
	total = getTotalPart2(engineParts, symbols)

	fmt.Printf("%v\n", total)
}

func getTotalPart2(engineParts []Coord, symbols []Coord) int {
	total := 0

	for _, symbol := range symbols {
		if symbol.value != "*" {
			continue
		}
		matches := []Coord{}
		for _, part := range engineParts {
			if part.isClose(symbol) {
				matches = append(matches, part)
			}
		}
		if len(matches) == 2 {
			num1, _ := strconv.Atoi(matches[0].value)
			num2, _ := strconv.Atoi(matches[1].value)
			total+= (num1*num2)
		}
	}
	return total
} 

func getTotalEngineParts(enginePars []Coord) int {
	total := 0

	for _, part := range enginePars {
		num, _ := strconv.Atoi(part.value)
		total+= num
	}

	return total
}

func getEnginePartsCloseToSymbol(engineParts []Coord, symbols []Coord) []Coord {
	result := []Coord{}

	for _, part := range engineParts {
		for _, symbol := range symbols {
			if part.isClose(symbol) {
				result = append(result, part)
				break
			}
		}
	}

	return result
}

func getEngineAndSymbols(engineMap [][]byte) ([]Coord, []Coord) {

	regexGetNumber := regexp.MustCompile(`[0-9]+`)
	engineParts := getCoord(engineMap, regexGetNumber)

	regexGetSymbol := regexp.MustCompile(`[^0-9|a-zA-Z|.]`)
	symbols := getCoord(engineMap, regexGetSymbol)

	return engineParts, symbols
}

func getCoord(engineMap [][]byte, regex *regexp.Regexp) []Coord {
	result := []Coord{}

	for y, line := range engineMap {
		matchesIndex := regex.FindAllIndex(line, -1)
		for _, index := range matchesIndex {
			value := line[index[0]:index[1]]
			coord := Coord{y, index[0], string(value)}
			result = append(result, coord)
		}
	}

	return result
}