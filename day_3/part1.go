package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	rawData := getRawData()
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

func getRawData() []byte {
	res, err := os.ReadFile("./day_3/input")
	checkErr(err)
	return res
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}