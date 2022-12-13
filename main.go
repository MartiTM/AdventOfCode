package main

import "fmt"

func main() {
	
	fmt.Println("%v\n", getAlphabetPosition('Z'))
}

func getAlphabetPosition(letter byte) int {
	return int(letter - 'a' + 1)-198
}