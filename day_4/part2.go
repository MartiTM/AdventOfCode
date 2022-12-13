package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	rawData := importRawData()
	data := processData(rawData)
	res := getOverlaps(data)
	fmt.Printf("There is %v overlap\n", res)
}

func getOverlaps(data [][][]byte) int {
	res := 0
	for _, row := range data {
		if isOverlap(row[1], row[2], row[3], row[4]) {
			res++
		}
	}
	return res
}

func isOverlap(a []byte, b []byte, c []byte, d []byte) bool {
	
	if bytesStringToInt(a) <= bytesStringToInt(c) && bytesStringToInt(b) >= bytesStringToInt(c) {
		return true
	}
	if bytesStringToInt(a) <= bytesStringToInt(d) && bytesStringToInt(b) >= bytesStringToInt(d) {
		return true
	}
	if bytesStringToInt(c) <= bytesStringToInt(a) && bytesStringToInt(d) >= bytesStringToInt(a) {
		return true
	}
	if bytesStringToInt(c) <= bytesStringToInt(b) && bytesStringToInt(d) >= bytesStringToInt(b) {
		return true
	}
	
	return false
}

func bytesStringToInt(b []byte) int {
	s := string(b)
	i, err := strconv.Atoi(s)
	checkErr(err)
	return i
}

func processData(raw []byte) [][][]byte {
	re := regexp.MustCompile(`(\d*)-(\d*),(\d*)-(\d*)`)
	return re.FindAllSubmatch(raw, -1)
}

func importRawData() []byte {
	res, err := os.ReadFile("./day_4/input")
	checkErr(err)
	return res
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}