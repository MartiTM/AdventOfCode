package main

import (
	"bytes"
	"fmt"

	"github.com/MartiTM/AdventOfCode2022/util"
)

func main() {
	rawData := util.GetRawData("./day_10/input")
	sumSignal := getSumSignals(rawData, []int{20, 60, 100, 140, 180, 220})
	fmt.Printf("Score : %v\n", sumSignal)
}

func getSumSignals(data []byte, targetSignals []int) int {
	sum := 0
	rows := bytes.Split(data, []byte("\n"))

	countCycles := 0
	x := 1

	for _, row := range rows {
		sum = applySignal(sum, x, countCycles, targetSignals)
		countCycles++
		if bytes.Equal(row, []byte("noop")) {
			continue
		}

		sum = applySignal(sum, x, countCycles, targetSignals)
		signal := util.BytesStringToInt(row[5:])
		x+=signal
		countCycles++
	}

	return sum
}

func applySignal(total int, x int, cycles int, targetSignals []int) int {
	if !isIn(targetSignals, cycles) {
		return total
	}
	fmt.Printf("Cycle %v, x %v, res %v, total %v\n", cycles, x, (x*cycles), total + (x*cycles))
	return total + (x*cycles)
}

func isIn(targetSignals []int, x int) bool {
	for _, target := range targetSignals {
		if target == x {
			return true
		}
	}
	return false
}