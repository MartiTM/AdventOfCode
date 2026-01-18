package main

import (
	"bytes"
	"fmt"
	"sort"

	"github.com/MartiTM/AdventOfCode2022/util"
)

type RangeFreshIngredients struct {
	start int
	end int
}

func (r RangeFreshIngredients) isFresh(ingredients int) bool {
	return (r.start <= ingredients) && (ingredients <= r.end)
}

func NewRangeFreshIngredients(data []byte) RangeFreshIngredients{
	listNum := bytes.Split(data, []byte("-"))

	return RangeFreshIngredients{start: util.BytesStringToInt(listNum[0]), end: util.BytesStringToInt(listNum[1])}
}

type ListeRange struct {
	liste []RangeFreshIngredients
}

func NewListeRange() ListeRange{
	return ListeRange{liste: []RangeFreshIngredients{}}
}

func main() {
	rawData := util.GetRawData("./chall1")

	rawInputs := bytes.Split(rawData, []byte("\n\n"))

	listRangeFreshIngredients := bytes.Split(rawInputs[0], []byte("\n"))
	listIngredients := bytes.Split(rawInputs[1], []byte("\n"))

	listRange := []RangeFreshIngredients{}

	for _, freshIngredients := range listRangeFreshIngredients {
		listRange = append(listRange, NewRangeFreshIngredients(freshIngredients))
	}

	total := 0

	for _, ingredients := range listIngredients {
		for _, r := range listRange {
			if r.isFresh(util.BytesStringToInt(ingredients)) {
				// fmt.Printf("%d <= %s <= %d\n", r.start, ingredients, r.end)
				total++
				break
			}
		}
	}

	fmt.Printf("Total fresh ingredients in list : %d\n", total)
	
	
	// Trier les pairs en fonction du début

	sort.SliceStable(listRange, func(i, j int) bool {
		return listRange[i].start < listRange[j].start
	})

	finalRange := []RangeFreshIngredients{}

	y := 0

	for i := 0; i < len(listRange); {
		if len(finalRange) == 0 {
			finalRange = append(finalRange, listRange[i])
			i++
			continue
		}

		// début de n+1 est compris dans n
		if finalRange[y].start <= listRange[i].start && listRange[i].start <= finalRange[y].end {
			if listRange[i].end > finalRange[y].end {
				finalRange[y].end = listRange[i].end
			}
			i++
			continue
		}
		finalRange = append(finalRange, listRange[i])
		y++
		i++
	}

	totalFresh := 0

	for _, rangeFresh := range finalRange {
		totalFresh += rangeFresh.end - rangeFresh.start + 1
	}

	// fmt.Printf("sorted range : \n%s\n", finalRange)
	
	fmt.Printf("Total fresh ingredients in range : %d\n", totalFresh)
}