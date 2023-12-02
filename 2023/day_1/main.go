package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MartiTM/AdventOfCode2022/util"
)

func isInt(r rune) bool {
	switch {
	case r >= '0' && r <= '9':
		return false
	case r == '\n':
		return false
	default:
		return true
	}
}

func filterDigitOfEachLine(data string) []string {
	arrayOnlyInt := strings.FieldsFunc(data, isInt)
	dataOnlyInt := strings.Join(arrayOnlyInt, "")
	return strings.Split(dataOnlyInt, "\n")
}

func getFirstAndLastChar(array []string) []string {
	result := []string{}
	for _, line := range array {
		firstChar := line[0]
		lastChar := line[len(line)-1]
		result = append(result, string(firstChar) + string(lastChar))
	}
	return result
}

func mapReduceStringsToTotalInt(array []string) int {
	total := 0
	for _, line := range array {
		num, err := strconv.Atoi(line)
		util.CheckErr(err)
		total += num
	}
	return total
}

func part1() {
	data := util.GetStringDataFromFile("./2023/day_1/input")
	arrayDigit := filterDigitOfEachLine(data)
	filterArray := getFirstAndLastChar(arrayDigit)
	total := mapReduceStringsToTotalInt(filterArray)
	fmt.Printf("%v\n", total)
}

// Spelling problem 
// exemple eightwothree we transform two before eight so the t of 8 is gone and we didn't transform the number
// Same with one and two
func transformStringDigitToDigit(data string) string {
	r := strings.NewReplacer(
		"one", 		"1e",
		"two", 		"2o",
		"three", 	"3e",
		"four", 	"4r",
		"five", 	"5e",
		"six", 		"6x",
		"seven", 	"7n",
		"eight", 	"8t",
		"nine", 	"9e",
	)
	data = r.Replace(data)
	return r.Replace(data)
}

func part2() {
	data := util.GetStringDataFromFile("./2023/day_1/input")
	dataDigited := transformStringDigitToDigit(data)
	arrayDigit := filterDigitOfEachLine(dataDigited)
	filterArray := getFirstAndLastChar(arrayDigit)
	total := mapReduceStringsToTotalInt(filterArray)
	fmt.Printf("%v\n", total)
}

func main() {
	part2()
}