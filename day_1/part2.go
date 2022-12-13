package main

import (
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	rawData, err := getRawData()
	checkErr(err)
	max := foundBiggestElf(rawData)
	fmt.Printf("%v\n", max)
}

func getRawData() ([]byte, error) {
	rawData, err := os.ReadFile("./day_1/input")
	return rawData, err
}

func foundBiggestElf(rawData []byte) (int) {
	// \n\n en hexa c'est 0a0a, si il y en a 2 de suite alors on saute une ligne donc on marque la fin de la poche de l'elf
	rawDataByElf := bytes.Split(rawData, []byte("\n\n"))
	var scoreBoard []int = []int{}
	for _, pocketElf := range rawDataByElf {
		var count int = 0
		foods := bytes.Split(pocketElf, []byte("\n"))
		for _, food := range foods {
			add, _ := strconv.Atoi(string(food))
			count += add
		}

		if len(scoreBoard) < 3 {
			scoreBoard = append(scoreBoard, count)
			continue
		}
		sort.Ints(scoreBoard)
		if count > scoreBoard[0] {
			scoreBoard[0] = count
		}
	}
	sum := 0
	for _, val := range scoreBoard {
		sum += val
	}
	return sum
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}