package main

import (
	"fmt"

	"github.com/MartiTM/AdventOfCode2022/util"
)

func part1() {
	rawData := util.GetRawData("")
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

func part2() {
	rawData := util.GetRawData("")
	markerPosition := findMarkerPosition(rawData)
	fmt.Printf("Market at %v\n", markerPosition)
}

// type Fifo []byte

// func (f *Fifo) Push(b byte) {
// 	*f = append(*f, b)
// }

// func (f *Fifo) Pop() byte {
// 	res := (*f)[0]
// 	*f = (*f)[1:]
// 	return res
// }

// func (f Fifo) isAllDifferents() bool {
// 	for i, ref := range f {
// 		for y, test := range f {
// 			if i == y {
// 				continue
// 			}
// 			if ref == test {
// 				return false
// 			}
// 		}

// 	}
// 	return true
// }


// func findMarkerPosition(data []byte) int {
// 	var forthlastest Fifo = Fifo{}
// 	for i, char := range data {
// 		if len(forthlastest) < 14 {
// 			forthlastest.Push(char)
// 			continue
// 		}
// 		if forthlastest.isAllDifferents() {
// 			return i
// 		}
// 		forthlastest.Push(char)
// 		forthlastest.Pop()
// 	}
// 	return -1
// }