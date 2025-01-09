package main

import (
	"bytes"
	"fmt"
	"sort"

	"github.com/MartiTM/AdventOfCode2022/util"
)

func GetBothLists(data [][]byte) ([]int, []int) {
	list1 := []int{}
	list2 := []int{}

	for _, line := range data {
		before, after, _ := bytes.Cut(line, []byte("   "))

		val1 := util.BytesStringToInt(before)
		val2 := util.BytesStringToInt(after)

		list1 = append(list1, val1)
		list2 = append(list2, val2)
	}

	return list1, list2
}

func getDistance(list1 []int, list2[]int) int {
	distance := 0

	for i, val := range list1 {
		val1 := val
		val2 := list2[i]
		x := val1 - val2
		if x < 0 {
			x = -x
		}
		distance += x
	}
	return distance
}

type sortInts []int

func (s sortInts) Less(i, j int) bool {
    return s[i] < s[j]
}

func (s sortInts) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s sortInts) Len() int {
    return len(s)
}

func SortInts(r []int) []int {
    sort.Sort(sortInts(r))
    return r
}

func part1() {
	rawData := util.GetRawData("./2024/01_day/file1")
	data := bytes.Split(rawData, []byte("\n"))
	list1, list2 := GetBothLists(data)
	fmt.Printf("list1 :\n%d\nlist2 :\n%d\n", list1, list2)
	SortInts(list1)
	SortInts(list2)
	fmt.Printf("list1 :\n%d\nlist2 :\n%d\n", list1, list2)
	distance := getDistance(list1, list2)
	fmt.Printf("distance : %d\n", distance)
}

func toHashMap(list []int) map[int]int {

	m := make(map[int]int)

	for _, item := range(list) {
		if _, ok := m[item]; ok {
			m[item]++
		} else {
			m[item] = 1
		}
	}

	return m
}

func calcResultPart2(firstList []int, secondHashMap map[int]int) int {
	result := 0

	for _, item := range(firstList) {
		result += item * secondHashMap[item]
	}

	return result
}

func part2() {
	rawData := util.GetRawData("./2024/01_day/file2")
	data := bytes.Split(rawData, []byte("\n"))
	list1, list2 := GetBothLists(data)
	fmt.Printf("list1 :\n%d\nlist2 :\n%d\n", list1, list2)
	list2Hash := toHashMap(list2)
	result := calcResultPart2(list1, list2Hash)
	fmt.Printf("result : %d\n", result)
}

func main() {
	part2()
}