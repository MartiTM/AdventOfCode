package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"

	"github.com/MartiTM/AdventOfCode2022/util"
)

func main() {
	data := util.GetRawData("./2023/day_6/input")

	mapRaceTimeAndRecord := getRaceInfo(data, true)

	total := 0
	for time, record := range mapRaceTimeAndRecord {
		ways := countWayOfWinning(time, record)
		if total == 0 {
			total = ways
		} else {
			total*=ways
		}
	}
	fmt.Printf("%v\n", total)
}

func countWayOfWinning(time int, record int) int {
	total := 0
	for timeButonPush := 0; timeButonPush < time; timeButonPush++ {
		distance := timeButonPush * (time-timeButonPush)
		if distance > record {
			total++
		}
	}
	return total
}

func getRaceInfo(data []byte, isPart2 bool) map[int]int {
	mapRaceTimeAndRecord := make(map[int]int)

	if isPart2 {
		data = bytes.Replace(data, []byte(" "), []byte(""), -1)
	}
	twoLine := bytes.Split(data, []byte("\n"))

	regex := regexp.MustCompile(`[0-9]+`)

	raceTime := regex.FindAll(twoLine[0], -1)
	record := regex.FindAll(twoLine[1], -1)

	for i := 0; i < len(raceTime); i++ {
		time, _ := strconv.Atoi(string(raceTime[i]))
		distance, _ := strconv.Atoi(string(record[i]))
		mapRaceTimeAndRecord[time] = distance
	}

	return mapRaceTimeAndRecord
}