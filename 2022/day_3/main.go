package main

import (
	"bytes"
	"fmt"

	"github.com/MartiTM/AdventOfCode2022/util"
)

func part1() {
	rawData := util.GetRawData("")
	rowList := toRowList(rawData)
	score := getScore(rowList)
	fmt.Printf("Score : %v\n", score)
}

func getScore(rowList [][]byte) int {
	score := 0
	
	for _, row := range rowList {
		letter := getCommunLetter(row)
		score += getPrioritize(letter)
	}

	return score
}

func getCommunLetter(row []byte) byte {
	middle := len(row)/2
	part1 := row[:middle]
	part2 := row[middle:]

	for _, letter1 := range part1 {
		for _, letter2 := range part2 {
			if letter2 == letter1 {
				return letter2
			}
		}
	}

	panic("Carton rouge")
}

func getPrioritize(letter byte) int {
	if isUpperCase(letter) {
		return getAlphabetPosition(letter)-198
	}
	return getAlphabetPosition(letter)
}

func isUpperCase(letter byte) bool {
	return letter >= 'A' && letter <= 'Z'
}

func getAlphabetPosition(letter byte) int {
	return int(letter - 'a' + 1)
}

func toRowList(rawData []byte) [][]byte {
	return bytes.Split(rawData, []byte("\n"))
}

func part2() {
	rawData := util.GetRawData("")
	rowList := toRowList(rawData)
	score := getScore(rowList)
	fmt.Printf("Score : %v\n", score)
}

// func getScore(rowList [][]byte) int {
// 	score := 0
	
// 	for i := 0; i<len(rowList); i+=3  {
// 		letter := getCommunLetter(rowList[i], rowList[i+1], rowList[i+2])
// 		score += getPrioritize(letter)
// 	}

// 	return score
// }

// func getCommunLetter(row1 []byte, row2 []byte, row3 []byte) byte {

// 	for _, letter1 := range row1 {
// 		for _, letter2 := range row2 {
// 			if letter2 != letter1 {
// 				continue
// 			}
// 			for _, letter3 := range row3 {
// 				if letter3 == letter2 {
// 					return letter3
// 				}
// 			}
// 		}
// 	}

// 	panic("Carton rouge")
// }

// func getPrioritize(letter byte) int {
// 	if isUpperCase(letter) {
// 		return getAlphabetPosition(letter)-198
// 	}
// 	return getAlphabetPosition(letter)
// }

// func isUpperCase(letter byte) bool {
// 	return letter >= 'A' && letter <= 'Z'
// }

// func getAlphabetPosition(letter byte) int {
// 	return int(letter - 'a' + 1)
// }

// func toRowList(rawData []byte) [][]byte {
// 	return bytes.Split(rawData, []byte("\n"))
// }