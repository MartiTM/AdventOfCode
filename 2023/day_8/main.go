package main

import (
	"bytes"
	"fmt"
	"log"
	"regexp"

	"github.com/MartiTM/AdventOfCode2022/util"
)

func main() {
	data := util.GetRawData("./2023/day_8/input")

	instructions, currentNode := getPuzzleInput(data)

	for !bytes.Equal(currentNode.name, []byte("ZZZ")) {
		// fmt.Printf("%v\n", string(currentNode.name))
		switch instructions.Read() {
		case byte('L'):
			currentNode = *currentNode.leftNode
		case byte('R'):
			currentNode = *currentNode.rightNode
		default:
			log.Fatal("KC")
		}
	}

	fmt.Printf("%v steps\n", instructions.steps)
}

type Instruction struct {
	instructions []byte
	cursor int
	steps int
}

func (i *Instruction) Read() byte {
	instruction := i.instructions[i.cursor]
	i.cursor++
	i.steps++
	if i.cursor > len(i.instructions) - 1 {
		i.cursor = 0
	}

	return instruction
}

type Node struct {
	name []byte
	leftNode *Node
	rightNode *Node
}

func getPuzzleInput(data []byte) (Instruction, Node) {
	dataByLine := bytes.Split(data, []byte("\n"))

	instruction := Instruction{dataByLine[0], 0, 0}

	nodeMap := make(map[string]*Node)

	regletter := regexp.MustCompile(`[A-Z]{3}`)

	var rootNode Node

	for i := 2; i < len(dataByLine); i++ {
		line := dataByLine[i]
		match := regletter.FindAll(line, 3)

		// Create if not existe
		if _, ok := nodeMap[string(match[0])]; !ok {
			nodeMap[string(match[0])] = &Node{match[0], nil, nil}
		}
		if _, ok := nodeMap[string(match[1])]; !ok {
			nodeMap[string(match[1])] = &Node{match[1], nil, nil}
		}
		if _, ok := nodeMap[string(match[2])]; !ok {
			nodeMap[string(match[2])] = &Node{match[2], nil, nil}
		}

		currentNode := nodeMap[string(match[0])]
		leftNode := nodeMap[string(match[1])]
		rightNode := nodeMap[string(match[2])]

		currentNode.leftNode = leftNode
		currentNode.rightNode = rightNode

		if bytes.Equal(currentNode.name, []byte("AAA")) {
			rootNode = *currentNode
		}
	}

	return instruction, rootNode
}