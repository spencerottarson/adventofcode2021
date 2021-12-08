package main

import (
    "bufio"
    "fmt"
	"sort"
	"strings"
    
	"../util"
)

func median(numbers []int) int {
	sort.Ints(numbers)

	var median int
	if (len(numbers) % 2 == 1) {
		median = numbers[len(numbers) / 2]
	} else {
		median = (numbers[(len(numbers) / 2) - 1] + numbers[(len(numbers) / 2)]) / 2
	}

	return median
}

func mean(numbers []int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum / len(numbers)
}

func intAbs(value int) int {
	if (value >= 0) {
		return value
	} else {
		return -value
	}
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

	median := median(startingNumbers)
	mean := mean(startingNumbers)

	sumDiffMedian := 0
	sumDiffMean := 0

	for _, value := range startingNumbers {
		sumDiffMedian += intAbs(value - median)
		n := intAbs(value - mean)
		sumDiffMean += (n * (n + 1)) / 2
	}

	fmt.Println(sumDiffMedian)
	fmt.Println(sumDiffMean)
}