package main

import (
    "bufio"
    "fmt"
	"strings"
    
	"../util"
)

func advanceDay(currentDay [9]int) [9]int {
	var newDay [9]int
	newDay[8] = currentDay[0]
	newDay[6] = currentDay[0]

	for i :=1; i < 9; i++ {
		newDay[i-1] += currentDay[i]
	}

	return newDay
}

func main() {
	file := util.RequireFile("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	input := strings.Split(scanner.Text(), ",")
	var startingNumbers []int
	for _, stringValue := range input {
		startingNumbers = append(startingNumbers, util.RequireAtoI(stringValue))
	}

	var currentDay [9]int

	for _, value := range startingNumbers {
		currentDay[value]++
	}

	for i := 0; i < 256; i++ {
		currentDay = advanceDay(currentDay)
	}

	sum := 0
	for _, value := range currentDay {
		sum += value
	}
	fmt.Println(sum)
}