package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Fifo []byte

func (f *Fifo) Push(v byte) {
    *f = append(*f, v)
}

func (f *Fifo) Pop() byte {
    ret := (*f)[len(*f)-1]
	*f = (*f)[0 : len(*f)-1]

	return ret
}

func main() {
	rawData := getRawData()

	ship := getShip(rawData, 8, 9)
	moveArray := getMovements(rawData)

	after := applyMovements(ship, moveArray)
	fmt.Print("Le code est : ")
	for _, row := range after {
		fmt.Printf("%c", row.Pop())
	}
	fmt.Println()
}

func applyMovements(ship []Fifo, move [][][]byte) []Fifo {
	for _, row := range move {
		for i:=0; i<bytesStringToInt(row[1]); i++ {
			crate := ship[bytesStringToInt(row[2])-1].Pop()
			ship[bytesStringToInt(row[3])-1].Push(crate)
		}
	}
	return ship
}

func getMovements(data []byte) [][][]byte {
	re := regexp.MustCompile(`move (\d*) from (\d*) to (\d*)`)
	return re.FindAllSubmatch(data, -1)
}

func getShip(data []byte, startingRow int, nb int) []Fifo {
	var res []Fifo = make([]Fifo, nb)

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

func getRawData() []byte {
	raw, err := os.ReadFile("./day_5/input")
	checkErr(err)
	return raw
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func bytesStringToInt(b []byte) int {
	s := string(b)
	i, err := strconv.Atoi(s)
	checkErr(err)
	return i
}