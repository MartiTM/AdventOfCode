package main

import (
	"bytes"
	"fmt"
	"regexp"

	"github.com/MartiTM/AdventOfCode2022/util"
)

type Filo []byte

func (f *Filo) Push(v byte) {
    *f = append(*f, v)
}

func (f *Filo) Pop() byte {
    ret := (*f)[len(*f)-1]
	*f = (*f)[0 : len(*f)-1]

	return ret
}

func part1() {
	rawData := util.GetRawData("")

	ship := getShip(rawData, 8, 9)
	moveArray := getMovements(rawData)

	after := applyMovements(ship, moveArray)
	fmt.Print("Le code est : ")
	for _, row := range after {
		fmt.Printf("%c", row.Pop())
	}
	fmt.Println()
}

func applyMovements(ship []Filo, move [][][]byte) []Filo {
	for _, row := range move {
		for i:=0; i<util.BytesStringToInt(row[1]); i++ {
			crate := ship[util.BytesStringToInt(row[2])-1].Pop()
			ship[util.BytesStringToInt(row[3])-1].Push(crate)
		}
	}
	return ship
}

func getMovements(data []byte) [][][]byte {
	re := regexp.MustCompile(`move (\d*) from (\d*) to (\d*)`)
	return re.FindAllSubmatch(data, -1)
}

func getShip(data []byte, startingRow int, nb int) []Filo {
	var res []Filo = make([]Filo, nb)

	rows := bytes.Split(data, []byte("\n"))

	for i:=startingRow-1; i>=0; i-- {
		rowNum:=0
		for y:=1; y<len(rows[0]); y+=4 {
			if rows[i][y] != byte(' ') {
				res[rowNum].Push(rows[i][y])
			}
			rowNum++
		}
	}
	return res
}

// type Fifo []byte

// func (f *Fifo) Push(v []byte) {
//     for _, e := range v {
// 		*f = append(*f, e)
// 	}
// }

// func (f *Fifo) PushB(v byte) {
//     *f = append(*f, v)
// }

// func (f *Fifo) Pop(i int) []byte {
//     ret := (*f)[len(*f)-i:]
// 	*f = (*f)[:len(*f)-i]

// 	return ret
// }

func part2() {
	// rawData := util.GetRawData("")

	// ship := getShip(rawData, 8, 9)
	// moveArray := getMovements(rawData)

	// after := applyMovements(ship, moveArray)
	fmt.Print("Le code est : ")
	// for _, row := range after {
		// fmt.Printf("%c", row.Pop(1)[0])
	// }
	fmt.Println()
}

// func applyMovements(ship []Fifo, move [][][]byte) []Fifo {
// 	for _, row := range move {
// 		crate := ship[bytesStringToInt(row[2])-1].Pop(bytesStringToInt(row[1]))
// 		ship[bytesStringToInt(row[3])-1].Push(crate)
// 	}
// 	return ship
// }

// func getMovements(data []byte) [][][]byte {
// 	re := regexp.MustCompile(`move (\d*) from (\d*) to (\d*)`)
// 	return re.FindAllSubmatch(data, -1)
// }

// func getShip(data []byte, startingRow int, nb int) []Fifo {
// 	var res []Fifo = make([]Fifo, nb)

// 	rows := bytes.Split(data, []byte("\n"))

// 	for i:=startingRow-1; i>=0; i-- {
// 		rowNum:=0
// 		for y:=1; y<len(rows[0]); y+=4 {
// 			if rows[i][y] != byte(' ') {
// 				res[rowNum].PushB(rows[i][y])
// 			}
// 			rowNum++
// 		}
// 	}
// 	return res
// }