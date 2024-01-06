package main

import (
	"bytes"
	"fmt"
	"sort"

	"github.com/MartiTM/AdventOfCode2022/util"
)

type Hand struct {
	cards []byte
	bid int
}

func (h Hand) isBetter(other Hand) bool {
	if h.getType() != other.getType() {
		return h.getType() > other.getType()
	}

	for i, card := range h.cards {
		if card == other.cards[i] {
			continue
		}
		return cardToPoint(card) > cardToPoint(other.cards[i])
	}

	fmt.Printf("Aie Aie Aie\n")
	return false
}

func (h Hand) getType() int {
	cardMap := make(map[byte]int)
	for _, card := range h.cards {
		if _, ok := cardMap[card]; !ok {
			cardMap[card]=0
		}
		cardMap[card]++
	}

	// Carte haute
	if len(cardMap) == 5 {
		return 1
	}

	// 1 paire 
	if len(cardMap) == 4 {
		return 2
	}

	for _, count := range cardMap {
		switch count {
		// 5 king
		case 5:
			return 7
		// 4 king
		case 4:
			return 6
		case 3:
			// full house
			if len(cardMap) == 2 {
				return 5
			}
			// 3 king
			return 4
		}
	}
	
	// 2 pairs
	return 3
}

func main() {
	data := util.GetRawData("./2023/day_7/input")

	hands := rawDataToHands(data)

	sort.Slice(hands, func(i, j int) bool {return hands[i].isBetter(hands[j])})

	total := 0

	for i, hand := range hands {
		total += hand.bid * (len(hands) - i)
	}

	fmt.Printf("Total : %v\n", total)
}

func rawDataToHands(data []byte) []Hand {
	hands := []Hand{}

	for _, line := range bytes.Split(data, []byte("\n")) {
		cards := line[0:5]
		bid := util.BytesStringToInt(line[6:])
		hand := Hand{cards, bid}
		hands = append(hands, hand)
	}

	return hands
}

func cardToPoint(card byte) int {
	switch card {
	case byte('A'):
		return 14
	case byte('K'):
		return 13
	case byte('Q'):
		return 12
	case byte('J'):
		return 11
	case byte('T'):
		return 10		
	case byte('9'):
		return 9		
	case byte('8'):
		return 8		
	case byte('7'):
		return 7		
	case byte('6'):
		return 6		
	case byte('5'):
		return 5		
	case byte('4'):
		return 4		
	case byte('3'):
		return 3
	}
	return 2
}