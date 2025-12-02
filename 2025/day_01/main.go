package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

type Safe struct {
	Dial				int
	Zero_counter		int
	Zero_pass_counter	int
}

func New() Safe {
	return Safe{50, 0, 0}
}

func (s *Safe) turnRight(count int) {
	s.Zero_pass_counter += (s.Dial + count) / 100
	s.Dial = (s.Dial + count) % 100
	if (s.Dial == 0) {
		s.Zero_pass_counter--
	}
}

func (s *Safe) turnLeft(count int) {
	s.Zero_pass_counter += (count / 100)
	count = count % 100

	tmp := s.Dial - count
	if (tmp < 0) {
		tmp = 100 + tmp
		if (s.Dial != 0) {
			s.Zero_pass_counter++
		}
	}
	s.Dial = tmp
}

func (s *Safe) countZero() {
	if (s.Dial == 0) {
		s.Zero_counter++
	}
}

func (s *Safe) process(input []byte) error{
	
	rotation, err := strconv.Atoi(string(input[1:]))
	if (err != nil) {
		return err
	}
	// fmt.Printf("rot : %s\n", rotation)
	switch string(input[0]) {
	case "L":
		s.turnLeft(rotation)
	case "R":
		s.turnRight(rotation)
	default:
		return fmt.Errorf("wrong instrution, expected : R/L, received : %v", input)
	}

	s.countZero()
	return nil
}

func main() {
	// rawFile, err := os.ReadFile("test_1")
	rawFile, err := os.ReadFile("chall_1")

	if (err != nil) {
		panic(err)
	}

	rawFileByLine := bytes.Split(rawFile, []byte("\n"))

	safe := New()

	fmt.Printf("Début:\nCoffre sur : %v\n\n", safe.Dial)
	for _, line := range rawFileByLine {
		err := safe.process(line)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Mouvement : %s\n", line)
		fmt.Printf("Coffre sur : %v\n", safe.Dial)
		fmt.Printf("Nombre 0 vue : %v\n\n", safe.Zero_counter + safe.Zero_pass_counter)
	}

	fmt.Printf("Nombre de 0 : %v\n", safe.Zero_counter)
	fmt.Printf("Nombre de 0 vue : %v\n", (safe.Zero_counter + safe.Zero_pass_counter))
}