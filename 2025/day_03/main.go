package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

type Bank struct {
	Battries []byte
}

func NewBank(row []byte) Bank {
	return Bank{row}
}

// For 2 digits
// func (b Bank) getLargestVoltage() (int, error){
// 	indexLargeFirst := 0
	
// 	for i := 0; i < len(b.Battries)-1; i++ {
// 		if b.Battries[i] > b.Battries[indexLargeFirst] {
// 			indexLargeFirst = i
// 		}
// 	}

// 	indexLargeSecond := indexLargeFirst+1
	
// 	for i := indexLargeSecond; i < len(b.Battries); i++ {
// 		if b.Battries[i] > b.Battries[indexLargeSecond] {
// 			indexLargeSecond = i
// 		}
// 	}

// 	return strconv.Atoi(string([]byte{b.Battries[indexLargeFirst], b.Battries[indexLargeSecond]}))
// }

func (b Bank) getLargestVoltage(digit int) (int, error){
	index := []int{}

	for len(index) < digit {
		var startingIndex int
		if len(index) == 0 {
			startingIndex = 0
		} else {
			startingIndex = index[len(index)-1]+1
		}
		maxIndex := startingIndex

		for i := startingIndex; i < len(b.Battries)-(digit-len(index)-1); i++ {
			if b.Battries[i] > b.Battries[maxIndex] {
				maxIndex = i
			}
		}

		index = append(index, maxIndex)
	}

	result := []byte{}

	for _, i := range index {
		result = append(result, byte(b.Battries[i]))		
	}

	return strconv.Atoi(string(result))
}

func main() {
	rawInput, err := os.ReadFile("./chall1")
	if err != nil {
		panic(err)
	}

	sum := 0
	for _, row := range bytes.Split(rawInput, []byte("\n")) {
		bank := NewBank(row)
		voltage, err := bank.getLargestVoltage(12)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Bank : %s\n", row)
		fmt.Printf("Voltage : %d\n", voltage)
		sum += voltage
	}

	fmt.Printf("Total voltage max = %d\n", sum)
}