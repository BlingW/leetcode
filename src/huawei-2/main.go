package main

import (
	"bufio"
	"fmt"
	"os"

	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var recordArray []string
	fileMap := make(map[string]int)

	for {
		scanner.Scan()
		input := scanner.Text()
		if len(input) == 0{
			break
		}

		urlSlice := strings.Split(input, "\\")
		contentSlice := strings.Split(urlSlice[len(urlSlice) - 1], " ")
		fileName := contentSlice[0]
		if len(fileName) > 16{
			fileName = fileName[len(fileName) - 1 - 15 :]
		}
		lineValue := contentSlice[1]
		key := fileName + " " + lineValue


		_, exist := fileMap[key]
		if !exist{
			fileMap[key] = 1
			recordArray = append(recordArray, key)
		} else{
			fileMap[key] = fileMap[key] + 1
		}



	}
	if len(recordArray) <= 8{
		for _, record := range recordArray{
			time := fileMap[record]
			fmt.Printf("%s %d\n", record, time)
		}
	} else{
		for _, record := range recordArray[len(recordArray) - 1 - 7 :]{
			time := fileMap[record]
			fmt.Printf("%s %d\n", record, time)
		}
	}



}