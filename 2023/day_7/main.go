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

func (h Hand) getType() int {
	cardMap := make(map[byte]int)
	for _, card := range h.cards {
		if _, ok := cardMap[card]; !ok {
			cardMap[card]=0
		}
		cardMap[card]++
	}

	if len(cardMap) == 1 {
		return 7
	}

	if nbJ, ok := cardMap[byte('J')]; ok {
		highestCard := byte('J')
		numberOfCard := 0
		for card, nbCard := range cardMap {
			if card == byte('J') {
				continue
			}
			if nbCard > numberOfCard {
				highestCard = card
				numberOfCard = nbCard
			}
		}
		cardMap[highestCard] += nbJ
		delete(cardMap, byte('J'))
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

type ByHand []Hand

func (a ByHand) Len() int           { return len(a) }
func (a ByHand) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func (a ByHand) Less(i, j int) bool {
	if a[i].getType() != a[j].getType() {
		return a[i].getType() > a[j].getType()
	}

	for i, card := range a[i].cards {
		if card == a[j].cards[i] {
			continue
		}
		return cardToPoint(card) > cardToPoint(a[j].cards[i])
	}

	fmt.Printf("Aie Aie Aie\n")
	return false
}

func main() {
	data := util.GetRawData("./2023/day_7/input")

	hands := rawDataToHands(data)

	sort.Sort(ByHand(hands))

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
		return 1
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