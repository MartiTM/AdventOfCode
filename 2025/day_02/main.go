package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

type IdRange struct {
	firstId int
	lastId int
}

func NewIdRange(rawRange []byte) (IdRange, error){
	ids := bytes.Split(rawRange, []byte("-"))
	fId, err := strconv.Atoi(string(ids[0]))
	if err != nil {
		return IdRange{}, fmt.Errorf("problème convertion int : %s", ids[0])
	}
	secondId, err := strconv.Atoi(string(ids[1]))
	if err != nil {
		return IdRange{}, fmt.Errorf("problème convertion int : %s", ids[0])
	}
	return IdRange{fId, secondId}, nil
}

// you can find the invalid IDs by looking for any ID which is made only of some sequence of digits repeated twice. 
// So, 55 (5 twice), 6464 (64 twice), and 123123 (123 twice) would all be invalid IDs.
func (i IdRange) sumInvalidIds() int {
	sum := 0

	for id := i.firstId; id <= i.lastId; id++ {
		// fmt.Printf("Id actuel : %d\n", id)
		if isInvalidIdTwice(id) {
			sum+=id
		}
	}

	return sum
}

func isInvalidIdTwice(id int) bool{
	strId := []byte(strconv.Itoa(id))

	middle := len(strId)/2
	// fmt.Printf("Test : %s et %s\n", strId[:middle], strId[middle:])
	return bytes.Equal(strId[:middle],strId[middle:])
}

func main() {
	rawInput, err := os.ReadFile("./test_1")
	if err != nil {
		panic(err)
	}
	sum := 0
	for _, rawRange := range bytes.Split(rawInput, []byte(",")) {
		idRange, err := NewIdRange(rawRange)
		if err != nil {
			panic(err)
		}
		sum += idRange.sumInvalidIds()
	}

	fmt.Printf("Somme des Ids invalides : %d\n", sum)
}