package main

import (
	"bytes"
	"fmt"
	"math"
	"strconv"

	"github.com/MartiTM/AdventOfCode2022/util"
)

func getReports(data []byte) [][]int {
	reports := [][]int{}

	dataByLine := bytes.Split(data, []byte("\n"))

	for _, line := range(dataByLine) {
		numbers := bytes.Split(line, []byte(" "))
		col := []int{}
		for _, num := range(numbers) {
			n, _ := strconv.Atoi(string(num))
			col = append(col, n)
		}
		reports = append(reports, col)
	}

	return reports
}

func checkReport(report []int) (string, bool) {

	previousNum := 0
	isIncreasing := true

	for i, num := range(report) {
		if i == 0 {
			previousNum = num
			continue
		}
		if (num == previousNum) || (math.Abs(float64(num) - float64(previousNum)) > 3) {
			return "0 ou supp à 3", false
		}
		if i == 1 {
			isIncreasing = (num > previousNum)
		}
		if isIncreasing != (num > previousNum) {
			return "pas constant", false
		}
		previousNum = num
	}

	return "", true
}

func part1() {
	rawData := util.GetRawData("./2024/02/file1")
	fmt.Printf("rawData :\n%s\n", rawData)
	reports := getReports(rawData)
	fmt.Printf("reports :\n%d\n", reports)
	
	result := 0
	
	for _, report := range(reports) {
		if err, ok := checkReport(report); ok {
			result++
		} else {
			fmt.Printf("report :\n%d\nerr :%s\n", report, err)
		} 
	}
	
	fmt.Printf("reports :%d\n", result)
}

func checkReport2(report []int) (int, bool) {

	previousNum := 0
	isIncreasing := true

	for i, num := range(report) {
		if i == 0 {
			previousNum = num
			continue
		}
		if (num == previousNum) || (math.Abs(float64(num) - float64(previousNum)) > 3) {
			return i, false
		}
		if i == 1 {
			isIncreasing = (num > previousNum)
		}
		if isIncreasing != (num > previousNum) {
			return i, false
		}
		previousNum = num
	}

	return 0, true
}

func part2() {
	rawData := util.GetRawData("./2024/02/file2")
	fmt.Printf("rawData :\n%s\n", rawData)
	reports := getReports(rawData)
	fmt.Printf("reports :\n%d\n", reports)
	
	result := 0
	
	for _, report := range(reports) {
		if _, ok := checkReport2(report); ok {
			result++
		} else {
			l := len(report)
			fmt.Println(report)
			for i, _ := range(report) {
				tmpReport := make([]int, l)
				copy(tmpReport, report)
				if i == len(report) {
					tmpReport = tmpReport[:i-1]
				} else {
					tmpReport = append(tmpReport[:i], tmpReport[i+1:]...)
				}
				fmt.Println(report)
				fmt.Println(tmpReport)
				fmt.Printf("max lenght : %d, - tmp lenght : %d\n", l, len(tmpReport))
				
				if _, ok := checkReport2(tmpReport); ok {
					result++
					break
				}
			}
		} 
	}
	
	fmt.Printf("reports : %d\n", result)
}

func main() {
	part2()
}