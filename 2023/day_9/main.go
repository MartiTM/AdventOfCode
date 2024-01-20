package main

import (
	"fmt"
	"strings"

	"github.com/MartiTM/AdventOfCode2022/util"
)

type Report struct {
	values []int
}

func (r Report) getNextValue() int {

	lastestVal := []int{r.values[len(r.values)-1]}
	currentArray := r.values
	
	for(!isTheSame(currentArray)) {
		nextArray := []int{}
		var oldVal int
		
		for i, val := range currentArray {
			if i == 0 {
				oldVal = val
				continue
			}
	
			nextArray = append(nextArray, val - oldVal)
	
			oldVal = val
		}
	
		currentArray = nextArray
		lastestVal = append(lastestVal, nextArray[len(nextArray)-1])
	}

	addToTheNextVal := lastestVal[len(lastestVal)-1]
	for i := len(lastestVal)-2; i >= 0; i-- {
		addToTheNextVal = lastestVal[i] + addToTheNextVal
	}

	return addToTheNextVal
}

func (r Report) getPreviousValue() int {

	firstValue := []int{r.values[0]}
	currentArray := r.values
	
	for(!isTheSame(currentArray)) {
		nextArray := []int{}
		var oldVal int
		
		for i, val := range currentArray {
			if i == 0 {
				oldVal = val
				continue
			}
	
			nextArray = append(nextArray, val - oldVal)
	
			oldVal = val
		}
	
		currentArray = nextArray
		firstValue = append(firstValue, nextArray[0])
	}

	addToTheNextVal := firstValue[len(firstValue)-1]
	for i := len(firstValue)-2; i >= 0; i-- {
		addToTheNextVal = firstValue[i] - addToTheNextVal
	}

	return addToTheNextVal
}

func main() {
	data := util.GetStringDataFromFile("./2023/day_9/input")

	total := 0
	for _, line := range strings.Split(data, "\n") {
		report := Report{util.StringsToInt(strings.Split(line, " "))}
		next_value := report.getPreviousValue()

		total += next_value
	}

	fmt.Printf("%v\n", total)
}

func isTheSame(array []int) bool {
	for i := 1; i < len(array); i++ {
		if array[0] != array[i] {
			return false
		}
	}
	return true
}