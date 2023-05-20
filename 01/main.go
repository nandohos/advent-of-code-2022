package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	elvesWithItems := getArrayElvesItems(input)

	caloriesByElf := getTotalCaloriesByElf(elvesWithItems)
	sort.Ints(caloriesByElf)

	fmt.Printf("part 1: %d\n", caloriesByElf[len(caloriesByElf)-1])

	caloriesTop3 := 0
	for i := 0; i < 3; i++ {
		caloriesTop3 += caloriesByElf[len(caloriesByElf)-1-i]
	}

	fmt.Printf("part 2: %d\n", caloriesTop3)

}

func getArrayElvesItems(input string) (arrayRes [][]int) {
	// Get the elves from the input
	for _, elvesItems := range strings.Split(input, "\n\n") {
		row := []int{}
		// Get the items from each elf
		for _, item := range strings.Split(elvesItems, "\n") {
			itemAsInt, _ := strconv.Atoi(item)
			row = append(row, itemAsInt)
		}
		// Group them in a matrix [elf][items]
		arrayRes = append(arrayRes, row)
	}
	return arrayRes
}

func getTotalCaloriesByElf(input [][]int) []int {

	totals := []int{}
	for _, items := range input {
		var sum int
		for _, n := range items {
			sum += n
		}
		totals = append(totals, sum)
	}
	return totals
}
