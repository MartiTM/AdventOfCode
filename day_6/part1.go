package main

import (
	"fmt"
	"os"
)

func main() {
	rawData := getRawData()
	markerPosition := findMarkerPosition(rawData)
	fmt.Printf("Market at %v\n", markerPosition)
}

type Fifo []byte

func (f *Fifo) Push(b byte) {
	*f = append(*f, b)
}

func (f *Fifo) Pop() byte {
	res := (*f)[0]
	*f = (*f)[1:]
	return res
}

func (f Fifo) isAllDifferents() bool {
	for i, ref := range f {
		for y, test := range f {
			if i == y {
				continue
			}
			if ref == test {
				return false
			}
		}

	}
	return true
}


func findMarkerPosition(data []byte) int {
	var forthlastest Fifo = Fifo{}
	for i, char := range data {
		if len(forthlastest) < 4 {
			forthlastest.Push(char)
			continue
		}
		if forthlastest.isAllDifferents() {
			return i
		}
		forthlastest.Push(char)
		forthlastest.Pop()
	}
	return -1
}

func getRawData() []byte {
	data, err := os.ReadFile("./day_6/input")
	checkErr(err)
	return data
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}