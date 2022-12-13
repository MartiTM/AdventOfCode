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
	res := getFullyContains(data)
	fmt.Printf("There is %v fully contains\n", res)
}

func getFullyContains(data [][][]byte) int {
	res := 0
	for _, row := range data {
		if isFullyContain(row[1], row[2], row[3], row[4]) {
			res++
		}
	}
	return res
}

func isFullyContain(a []byte, b []byte, c []byte, d []byte) bool {
	
	if bytesStringToInt(a) <= bytesStringToInt(c) && bytesStringToInt(b) >= bytesStringToInt(d) {
		return true
	}
	if bytesStringToInt(a) >= bytesStringToInt(c) && bytesStringToInt(b) <= bytesStringToInt(d) {
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