package util

import (
	"bytes"
	"os"
	"sort"
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

func GetStringDataFromFile(path string) string {
	byteData := GetRawData(path)
	buffer := bytes.NewBuffer(byteData)
	return buffer.String()
}

func StringsToInt(s []string) []int {
	result := []int{}

	for _, val := range s {
		v, err := strconv.Atoi(val)
		CheckErr(err)
		result = append(result, v)
	}

	return result
}

type sortBytes []byte

func (s sortBytes) Less(i, j int) bool {
    return s[i] < s[j]
}

func (s sortBytes) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s sortBytes) Len() int {
    return len(s)
}

func SortString(r []byte) []byte {
    sort.Sort(sortBytes(r))
    return r
}