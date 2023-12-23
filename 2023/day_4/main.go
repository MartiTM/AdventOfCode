package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"

	"github.com/MartiTM/AdventOfCode2022/util"
)

type Scratchcard struct {
	winningNumbers []string
	numbers []string
	cardNumber int
}

func (c Scratchcard) getMatches() int {
	matches := 0
	for _, winNum := range c.winningNumbers {
		for _, num := range c.numbers {
			if winNum == num {
				matches+=1
			}
		}
	}
	return matches
}

func (c Scratchcard) getPoints() int {
	points := 0
	for i := 0; i < c.getMatches(); i++ {
		if points == 0 {
			points+=1
		} else {
			points*=2
		} 
	}
	return points
}

func main() {
	data := util.GetRawData("./2023/day_4/input")

	cards := getCards(data)

	total := 0

	for _, card := range cards {
		total += card.getPoints()
	}

	fmt.Printf("points %v\n", total)

	for i := 0; i < len(cards); i++ {
		match := cards[i].getMatches()
		for x := 0; x < cards[i].cardNumber; x++ {
			for y := 1; y <= match; y++ {
				cards[i+y].cardNumber++
			}
		}
	}

	total = 0
	for _, card := range cards {
		total += card.cardNumber
	}

	fmt.Printf("cards %v\n", total)
}

func getCards(data []byte) []Scratchcard {
	result := []Scratchcard{}
	
	regex := regexp.MustCompile(`( |[0-9]+)+\|( |[0-9]+)+`)
	regexNum := regexp.MustCompile(`([0-9]+)`)
	
	for _, line := range bytes.Split(data, []byte("\n")){
		lineFilter := regex.FindString(string(line))
		parts := strings.Split(lineFilter, "|")
		result = append(result, Scratchcard{regexNum.FindAllString(parts[0], -1), regexNum.FindAllString(parts[1], -1), 1})
	}

	return result
}