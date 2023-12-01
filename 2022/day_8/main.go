package main

import (
	"bytes"
	"fmt"

	"github.com/MartiTM/AdventOfCode2022/util"
)

func main() {
	part1()
	part2()
}

func part1() {
	rawData := util.GetRawData("./day_8/input")
	matrix := bytes.Split(rawData, []byte("\n"))
	visibleTree := getVisibleTrees(matrix)
	fmt.Printf("Visible trees %v\n", visibleTree)
}

func getVisibleTrees(matrix [][]byte) int {
	total := 0
	// on mesure l'interieure de la matrice
	for y := 1; y<len(matrix)-1; y++ {
		for x := 1; x<len(matrix[0])-1; x++ {
			// fmt.Printf("The %q in %v, %v is : ", matrix[y][x], x, y)
			// un 0 ne sera jamais vue car on ne voit pas 2 arbre de la meme hauteur
			if matrix[y][x] == byte('0') {
				continue
				// fmt.Printf("not visible because 0\n")
			}
			if isVisibleFromTheTop(y, x, matrix) {
				total++
				// fmt.Printf("visible from top\n")
				continue
			}
			if isVisibleFromTheBottom(y, x, matrix) {
				total++
				// fmt.Printf("visible from bottom\n")
				continue
			}
			if isVisibleFromTheLeft(y, x, matrix) {
				total++
				// fmt.Printf("visible from Left\n")
				continue
			}
			if isVisibleFromTheRight(y, x, matrix) {
				total++
				// fmt.Printf("visible from Right\n")
				continue
			}
			// fmt.Printf("not visible\n")
		}
	}
	edgeTrees := len(matrix)*2 + len(matrix[0])*2 - 4
	return total + edgeTrees
}

func isVisibleFromTheTop(y int, x int, matrix [][]byte) bool {
	visible := true
	for t := y-1; t>=0; t-- {
		if matrix[t][x] >= matrix[y][x] {
			visible = false
			break
		} 
	}
	return visible
}

func isVisibleFromTheBottom(y int, x int, matrix [][]byte) bool {
	visible := true
	for t := y+1; t<len(matrix); t++ {
		if matrix[t][x] >= matrix[y][x] {
			visible = false
			break
		} 
	}
	return visible
}

func isVisibleFromTheRight(y int, x int, matrix [][]byte) bool {
	visible := true
	for t := x+1; t<len(matrix[0]); t++ {
		if matrix[y][t] >= matrix[y][x] {
			visible = false
			break
		} 
	}
	return visible
}

func isVisibleFromTheLeft(y int, x int, matrix [][]byte) bool {
	visible := true
	for t := x-1; t>=0; t-- {
		if matrix[y][t] >= matrix[y][x] {
			visible = false
			break
		} 
	}
	return visible
}

func part2() {
	rawData := util.GetRawData("./day_8/input")
	matrix := bytes.Split(rawData, []byte("\n"))
	score := getScenicScore(matrix)
	fmt.Printf("Score is %v\n", score)
}

func getScenicScore(matrix [][]byte) int {
	max := 0
	// on mesure l'interieure de la matrice
	for y := 1; y<len(matrix)-1; y++ {
		for x := 1; x<len(matrix[0])-1; x++ {
			if matrix[y][x] == byte('0') {
				continue
			}
			score := scoreFromTop(y, x, matrix)
			score = score * scoreFromLeft(y, x, matrix)
			score = score * scoreFromBottom(y, x, matrix)
			score = score * scoreFromRight(y, x, matrix)
			
			if score >= max {
				max = score
			}
		}
	}
	return max
}

func scoreFromTop(y int, x int, matrix [][]byte) int {
	score := 1
	for t := y-1; t>0; t-- {
		if matrix[t][x] >= matrix[y][x] {
			break
		}
		score++
	}
	return score
}

func scoreFromBottom(y int, x int, matrix [][]byte) int {
	score := 1
	for t := y+1; t<len(matrix)-1; t++ {
		if matrix[t][x] >= matrix[y][x] {
			break
		} 
		score++
	}
	return score
}

func scoreFromRight(y int, x int, matrix [][]byte) int {
	score := 1
	for t := x+1; t<len(matrix[0])-1; t++ {
		if matrix[y][t] >= matrix[y][x] {
			break
		} 
		score++
	}
	return score
}

func scoreFromLeft(y int, x int, matrix [][]byte) int {
	score := 1
	for t := x-1; t>0; t-- {
		if matrix[y][t] >= matrix[y][x] {
			break
		} 
		score++
	}
	return score
}