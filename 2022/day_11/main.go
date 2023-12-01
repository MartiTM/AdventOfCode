package main

import (
	"bytes"
	"fmt"

	"github.com/MartiTM/AdventOfCode2022/util"
)

type Monkey struct {
	items []byte
	operationSigne byte
	operationVal []byte
	testVal int
	testTrue int
	testFalse int
}

func (m Monkey) Operation(item int) int {
	var value int

	if bytes.Equal(m.operationVal, []byte("old")) {
		value = item
	} else {
		value = util.BytesStringToInt(m.operationVal)
	}

	if m.operationSigne == byte('*') {
		return value*item
	} else {
		return value+item
	}
}

func (m Monkey) Test(item int) int {
	if item % m.testVal == 0 {
		return m.testTrue
	}
	return m.testFalse
}

func main() {
	rawData := util.GetRawData("./day_11/sample")
	rawDataToMonkeys(rawData)

	fmt.Printf("Score : %v\n", "e")
}

func rawDataToMonkeys(data []byte) []Monkey {
	rawData := bytes.Split(data, []byte("\n"))
	var monkeys []Monkey
	for i := 0; i < len(rawData); i+=7 {
		var monkey Monkey

		monkey.items = rawData[i+1][18:]
		monkey.operationSigne = rawData[i+2][23]

		monkeys = append(monkeys, monkey)
	}
	return monkeys
}