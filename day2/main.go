package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	result_map_1 := map[string]int{
		"AX": 4,
		"AY": 8,
		"AZ": 3,
		"BX": 1,
		"BY": 5,
		"BZ": 9,
		"CX": 7,
		"CY": 2,
		"CZ": 6,
	}

	result_map_2 := map[string]int{
		"AX": 3,
		"AY": 4,
		"AZ": 8,
		"BX": 1,
		"BY": 5,
		"BZ": 9,
		"CX": 2,
		"CY": 6,
		"CZ": 7,
	}

	result1 := 0
	result2 := 0

	for scanner.Scan() {
		line := scanner.Text()
		data := strings.ReplaceAll(line, " ", "")
		if val, ok := result_map_1[data]; ok {
			result1 += val
		}
		if val, ok := result_map_2[data]; ok {
			result2 += val
		}
	}

	fmt.Println("result_1: ", result1)
	fmt.Println("result_2: ", result2)

}
