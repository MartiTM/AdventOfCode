package main

import (
	"fmt"
	"regexp"

	"github.com/MartiTM/AdventOfCode2022/util"
)

func part1() {
	rawData := util.GetRawData("")
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
	
	if util.BytesStringToInt(a) <= util.BytesStringToInt(c) && util.BytesStringToInt(b) >= util.BytesStringToInt(d) {
		return true
	}
	if util.BytesStringToInt(a) >= util.BytesStringToInt(c) && util.BytesStringToInt(b) <= util.BytesStringToInt(d) {
		return true
	}
	
	return false
}

func processData(raw []byte) [][][]byte {
	re := regexp.MustCompile(`(\d*)-(\d*),(\d*)-(\d*)`)
	return re.FindAllSubmatch(raw, -1)
}

func part2() {
	rawData := util.GetRawData("")
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
	
	if util.BytesStringToInt(a) <= util.BytesStringToInt(c) && util.BytesStringToInt(b) >= util.BytesStringToInt(c) {
		return true
	}
	if util.BytesStringToInt(a) <= util.BytesStringToInt(d) && util.BytesStringToInt(b) >= util.BytesStringToInt(d) {
		return true
	}
	if util.BytesStringToInt(c) <= util.BytesStringToInt(a) && util.BytesStringToInt(d) >= util.BytesStringToInt(a) {
		return true
	}
	if util.BytesStringToInt(c) <= util.BytesStringToInt(b) && util.BytesStringToInt(d) >= util.BytesStringToInt(b) {
		return true
	}
	
	return false
}

// func processData(raw []byte) [][][]byte {
// 	re := regexp.MustCompile(`(\d*)-(\d*),(\d*)-(\d*)`)
// 	return re.FindAllSubmatch(raw, -1)
// }