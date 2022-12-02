package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var (
		elves []int
		elf   []int
	)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 1 {
			elves, elf = append(elves, sum(elf)), nil
		}
		a, _ := strconv.Atoi(line)
		elf = append(elf, a)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(elves)))

	fmt.Printf("Part One: %d\n", elves[0])
	fmt.Printf("Part Two: %d\n", sum(elves[:3]))
}
