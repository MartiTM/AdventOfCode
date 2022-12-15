package util

import (
	"os"
	"strconv"
)

func GetRawData(path string) []byte {
	data, err := os.ReadFile(path)
	CheckErr(err)
	return data
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func BytesStringToInt(b []byte) int {
	s := string(b)
	i, err := strconv.Atoi(s)
	CheckErr(err)
	return i
}