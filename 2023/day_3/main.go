package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"

	"github.com/MartiTM/AdventOfCode2022/util"
)

func getCoordByLine(enginMap [][]byte, regex *regexp.Regexp) [][][]int {
	result := [][][]int{}

	for _, line := range enginMap {
		result = append(result, regex.FindAllIndex(line, -1))
	}

	return result
}

func getNumberCloseToSymbolCoord(nbCoord [][][]int , symbolCoord [][][]int, neededMatch int) [][][]int {
	result := [][][]int{}
	for y, line := range nbCoord {
		resultLine := [][]int{}
		for _, coord := range line {
			xStart := coord[0]
			xEnd := coord[1] - 1
			if isCloseFromSymbol(y, xStart, xEnd, symbolCoord, neededMatch) {
				resultLine = append(resultLine, coord)
			}
		}
		result = append(result, resultLine)
	}
	return result
}

func isCloseFromSymbol(numberY int , numberStartX int, numberEndX int, symbolCoord [][][]int, neededMatch int) bool {
	startX := numberStartX - 1
	endX := numberEndX + 1
	
	startY := numberY - 1
	endY := numberY + 1

	if startY < 0 {
		startY = numberY
	}
	if endY >= len(symbolCoord) {
		endY = len(symbolCoord) - 1
	}

	match := 0
	
	// fmt.Printf("Ligne %v : coord %v:%v\n", numberY, numberStartX, numberEndX)
	for y := startY; y <= endY; y++ {
		for _, symboleCoord := range symbolCoord[y] {
			// fmt.Println(startX, " >= ", symboleCoord[0], " <= ", endX )
			if startX <= symboleCoord[0] && symboleCoord[0] <= endX {
				// fmt.Printf("trouvé sur %v\n", y)
				match++
			}
		}
	}

	// fmt.Printf("pas trouvé de %v a %v \n", startY, endY)
	return match >= neededMatch
}

func getTotal(enginMap [][]byte, numberCoord [][][]int) int {
	total := 0
	
	for i, coords := range numberCoord {
		for _, coord := range coords {
			bArray := enginMap[i][coord[0]:coord[1]]
			num, _ := strconv.Atoi(string(bArray))
			total+= num
		} 
	}

	return total
}

func part1(){
	data := util.GetRawData("./2023/day_3/input")
	enginMap := bytes.Split(data, []byte("\n"))

	regexGetNumber := regexp.MustCompile(`[0-9]+`)
	numberCoord := getCoordByLine(enginMap, regexGetNumber)

	regexGetSymbol := regexp.MustCompile(`[^0-9|a-zA-Z|.]`)
	symbolCoord := getCoordByLine(enginMap, regexGetSymbol)

	numberCloseToSymbolCoord := getNumberCloseToSymbolCoord(numberCoord, symbolCoord, 1)

	total := getTotal(enginMap, numberCloseToSymbolCoord)

	fmt.Printf("%v\n", total)
}

func part2() {
	data := util.GetRawData("./2023/day_3/input")
	enginMap := bytes.Split(data, []byte("\n"))

	regexGetNumber := regexp.MustCompile(`[0-9]+`)
	numberCoord := getCoordByLine(enginMap, regexGetNumber)

	regexGetSymbol := regexp.MustCompile(`[*]`)
	symbolCoord := getCoordByLine(enginMap, regexGetSymbol)

	numberCloseToSymbolCoord := getNumberCloseToSymbolCoord(numberCoord, symbolCoord, 2)

	total := getTotal(enginMap, numberCloseToSymbolCoord)

	fmt.Printf("%v\n", total)
}

func main() {
	part2()
}