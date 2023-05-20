package main

import (
	_ "embed"
	"fmt"
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
  
  var max int
  for _, x := range caloriesByElf{
    if x>max {
     max=x 
    }
  }
  fmt.Printf("part 1: %d", max)

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

  totals:= []int{}
  for _, items := range input {
    var sum int
    for _,n := range items {
       sum += n
    }
    totals = append(totals, sum)
  }
  return totals
}
