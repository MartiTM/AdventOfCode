package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	rawData, err := getRawData()
	checkErr(err)
	at, max := foundBiggestElf(rawData)
	fmt.Printf("Biggest Elf at %v and carrying %v\n", at, max)
}

func getRawData() ([]byte, error) {
	rawData, err := os.ReadFile("./day_1/input_day_1.txt")
	return rawData, err
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

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}