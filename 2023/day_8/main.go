package main

import (
	"bytes"
	"fmt"
	"regexp"

	"github.com/MartiTM/AdventOfCode2022/util"
)

func main() {
	data := util.GetRawData("./2023/day_8/input")

	instructions, currentNodes := getPuzzleInput(data)


	for _, currentNode := range currentNodes {

		fmt.Printf("----start node %s-------\n", currentNode.name)
		for i := 0; i < 3; i++ {
			for !bytes.Contains(currentNode.name, []byte("Z")) {
				switch instructions.Read() {
				case byte('L'):
					currentNode = *currentNode.leftNode
				case byte('R'):
					currentNode = *currentNode.rightNode
				}
			}
			fmt.Printf("node %s take %v steps\n", currentNode.name, instructions.steps)
			switch instructions.Read() {
			case byte('L'):
				currentNode = *currentNode.leftNode
			case byte('R'):
				currentNode = *currentNode.rightNode
			}
		}

		instructions.steps = 0
		instructions.cursor = 0
	}

	// PPCM on the 6 roads give us the awser
	fmt.Printf("%v\n", 283*73*71*53*61*79*59)
	

	// Brut force methods not very conclusive
	// for !isAllAtTheEnd(currentNodes) {
	// 	instruction := instructions.Read()
	// 	nextNodes := []Node{}
	// 	for _, currentNode := range currentNodes {
	// 		switch instruction {
	// 		case byte('L'):
	// 			nextNodes = append(nextNodes, *currentNode.leftNode)
	// 		case byte('R'):
	// 			nextNodes = append(nextNodes, *currentNode.rightNode)
	// 		default:
	// 			log.Fatal("KC")
	// 		}
	// 	}
	// 	currentNodes = nextNodes
	// 	if instructions.steps % 1000000 == 0 {
	// 		fmt.Printf("%v\n", instructions.steps)
	// 	}
	// }
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

func isAllAtTheEnd(nodes []Node) bool {
	for _, node := range nodes {
		if !bytes.Contains(node.name, []byte("Z")) {
			return false
		}
	}
	return true
}

func getPuzzleInput(data []byte) (Instruction, []Node) {
	dataByLine := bytes.Split(data, []byte("\n"))

	instruction := Instruction{dataByLine[0], 0, 0}

	nodeMap := make(map[string]*Node)

	regletter := regexp.MustCompile(`[A-Z0-9]{3}`)

	rootNodes := []Node{}

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

		if bytes.Contains(currentNode.name, []byte("A")) {
			rootNodes = append(rootNodes, *currentNode)
		}
	}

	return instruction, rootNodes
}