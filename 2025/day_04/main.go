package main

import (
	"bytes"
	"fmt"
	"os"
)

type Coordinate struct {
	X int
	Y int
}

func foundAccessibleRollsCoord(table [][]byte)  []Coordinate {

	coordinates := []Coordinate{}

	for y, row := range table {
		for x, _ := range row {
			num, err := getNumberCloseRolls(table, x, y)
			// fmt.Printf("%d < %d\n", num, lessThen)
			if err != nil {
				continue
			}
			if num < 4 {
				coordinates = append(coordinates, Coordinate{x, y})
			}
		}
	}

	return coordinates
}

func getNumberCloseRolls(table [][]byte, x int , y int) (int, error){
	maxLenY := len(table)
	maxLenX := len(table[0])

	sum := 0

	if table[y][x] != byte('@') {
		return sum, fmt.Errorf("not a @")
	}

	if (y != 0) && (x != 0) {
		if table[y-1][x-1] == byte('@') {
			sum ++
		}
	}

	if (y != maxLenY-1) && (x != maxLenX-1) {
		if table[y+1][x+1] == byte('@') {
			sum ++
		}
	}
	
	if (y != maxLenY-1) && (x != 0) {
		if table[y+1][x-1] == byte('@') {
			sum ++
		}
	}

	if (y != 0) && (x != maxLenX-1) {
		if table[y-1][x+1] == byte('@') {
			sum ++
		}
	}

	if (x != 0) {
		if table[y][x-1] == byte('@') {
			sum ++
		}
	}

	if (x != maxLenX-1) {
		if table[y][x+1] == byte('@') {
			sum ++
		}
	}

	if (y != 0) {
		if table[y-1][x] == byte('@') {
			sum ++
		}
	}

	if (y != maxLenY-1) {
		if table[y+1][x] == byte('@') {
			sum ++
		}
	}

	return sum, nil
}

func removeRollsFromTable(table [][]byte, listCoord []Coordinate) [][]byte{
	for _, coord := range listCoord {
		table[coord.Y][coord.X] = byte('x')
	}

	return table
}

func main() {
	rawInput, err := os.ReadFile("./chall1")
	if err != nil {
		panic(err)
	}

	rawTable := bytes.Split(rawInput, []byte("\n"))

	
	listRemove := foundAccessibleRollsCoord(rawTable)
	totalRemove := len(listRemove)
	
	for len(listRemove) != 0  {
		rawTable = removeRollsFromTable(rawTable, listRemove)

		listRemove = foundAccessibleRollsCoord(rawTable)
		totalRemove += len(listRemove)
		
	}

	fmt.Printf("accessibleRolls: %d\n", totalRemove)
}
