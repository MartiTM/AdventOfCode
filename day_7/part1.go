package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Folder struct {
	size int
	folders []*Folder
}

func (f Folder) getSize() int {
	sum := 0
	for _, folder := range f.folders {
		sum+=folder.getSize()
	}
	return sum+f.size
}

type FileSystem map[string]*Folder

func main() {
	rawData := getRawData()
	commands := getCommands(rawData)
	system := createFileSystem(commands)
	size := getSizes(system)
	fmt.Printf("size %v\n", size)
}

func getSizes(system FileSystem) int {
	sum := 0
	for _, folder := range system {
		size := folder.getSize()
		if size <= 100000 {
			sum+=size
		}
	}
	return sum
}

func createFileSystem(commands [][]byte) FileSystem {
	var system FileSystem = make(FileSystem)
	var currentFolder string = ""

	for _, command := range commands {
		if regexp.MustCompile(`ls\s`).Match(command) {
			applyLs(command, &system, currentFolder)
		} else {
			currentFolder = applyCd(command, &system, currentFolder)
		}
	}
	return system
}

func applyLs(command []byte, system *FileSystem, currentFolder string) {
	rows := bytes.Split(command, []byte("\n"))[1:]
	for _, row := range rows {
		if bytes.Compare(row, []byte("")) == 0 {
			continue
		}
		if regexp.MustCompile(`dir`).Match(row) {
			name := row[4:]
			folder := &Folder{0, []*Folder{}}
			path := currentFolder
			if currentFolder != "/" {
				path+="/"
			}
			path+=string(name)
			(*system)[path] = folder
			(*system)[currentFolder].folders = append((*system)[currentFolder].folders, folder)
			continue
		}
		size := regexp.MustCompile(`^\d*`).Find(row)
		(*system)[string(currentFolder)].size += bytesStringToInt(size)
	}
}

func applyCd(command []byte, system *FileSystem, currentFolder string) string {
	folderName := command[4:len(command)-1]
	
	if bytes.Compare(folderName, []byte("..")) == 0 {
		split :=  strings.Split(currentFolder, "/")
		return strings.Join(split[:len(split)-1], "/")
	}

	if bytes.Compare(folderName, []byte("/")) == 0 {
		root := &Folder{0, []*Folder{}}
		(*system)["/"] = root
		return "/"
	}

	path := currentFolder
	if currentFolder != "/" {
		path+="/"
	}
	path+=string(folderName)
	return path
}

func getCommands(data []byte) [][]byte {
	commands := bytes.Split(data, []byte("$"))[1:]
	return commands
}

func getRawData() []byte {
	data, err := os.ReadFile("./day_7/input")
	checkErr(err)
	return data
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func bytesStringToInt(b []byte) int {
	s := string(b)
	i, err := strconv.Atoi(s)
	checkErr(err)
	return i
}