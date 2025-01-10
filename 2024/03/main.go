package main

import (
	"bytes"
	"fmt"
	"regexp"

	"github.com/MartiTM/AdventOfCode2022/util"
)

func part1() {
	rawData := util.GetRawData("./2024/03/file1")

	reg := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	muls := reg.FindAllSubmatch(rawData, -1)
	fmt.Printf("mul : %s\n", muls)

	res := 0
	for _, mul := range(muls) {
		g1 := util.BytesStringToInt(mul[1])
		g2 := util.BytesStringToInt(mul[2])
		res += g1 * g2
	}

	fmt.Printf("res : %d\n", res)
}

func part2() {
	rawData := util.GetRawData("./2024/03/file2")

	reg := regexp.MustCompile(`(mul\((\d{1,3}),(\d{1,3})\))|(do\(\))|(don't\(\))`)

	muls := reg.FindAllSubmatch(rawData, -1)
	fmt.Printf("mul : %s\n", muls)
	
	res := 0
	isDo := true
	for _, mul := range(muls) {
		if bytes.Equal([]byte("do()"), mul[0]) {
			isDo = true
			continue
		}
		if bytes.Equal([]byte("don't()"), mul[0]) {
			isDo = false
			continue
		}
		if !isDo {
			continue
		}

		g1 := util.BytesStringToInt(mul[2])
		g2 := util.BytesStringToInt(mul[3])

		res += g1 * g2
	}

	fmt.Printf("res : %d\n", res)
}

func main() {
	part2()
}