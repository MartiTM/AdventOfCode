package main

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/MartiTM/AdventOfCode2022/util"
)

func part1() {
	rawData := util.GetRawData("")
	at, max := foundBiggestElf(rawData)
	fmt.Printf("Biggest Elf at %v and carrying %v\n", at, max)
}

func foundBiggestElf(rawData []byte) (int, int) {
	// \n\n en hexa c'est 0a0a, si il y en a 2 de suite alors on saute une ligne donc on marque la fin de la poche de l'elf
	rawDataByElf := bytes.Split(rawData, []byte("\n\n"))
	maxAt := 0
	var max int = 0
	for i, pocketElf := range rawDataByElf {
		var count int = 0
		foods := bytes.Split(pocketElf, []byte("\n"))
		for _, food := range foods {
			add, _ := strconv.Atoi(string(food))
			count += add
		}
		if count > max {
			maxAt = i
			max = count
		}
	}
	return maxAt+1, max
}

func part2() {
	// rawData := util.GetRawData("")
	// max := foundBiggestElf(rawData)
	// fmt.Printf("%v\n", max)
}

// func foundBiggestElf(rawData []byte) (int) {
// 	// \n\n en hexa c'est 0a0a, si il y en a 2 de suite alors on saute une ligne donc on marque la fin de la poche de l'elf
// 	rawDataByElf := bytes.Split(rawData, []byte("\n\n"))
// 	var scoreBoard []int = []int{}
// 	for _, pocketElf := range rawDataByElf {
// 		var count int = 0
// 		foods := bytes.Split(pocketElf, []byte("\n"))
// 		for _, food := range foods {
// 			add, _ := strconv.Atoi(string(food))
// 			count += add
// 		}

// 		if len(scoreBoard) < 3 {
// 			scoreBoard = append(scoreBoard, count)
// 			continue
// 		}
// 		sort.Ints(scoreBoard)
// 		if count > scoreBoard[0] {
// 			scoreBoard[0] = count
// 		}
// 	}
// 	sum := 0
// 	for _, val := range scoreBoard {
// 		sum += val
// 	}
// 	return sum
// }