package main

import (
    "bufio"
    "fmt"
	"log"
	"strconv"

	"../util"
)

func main() {
	file := util.RequireFile("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var binaryNumbers []string

	for scanner.Scan() {
		input := scanner.Text()
		binaryNumbers = append(binaryNumbers, input)
	}

	if (len(binaryNumbers) == 0) {
		log.Fatal("Error: input cannot be empty")
	}

	inputSize := len(binaryNumbers)
	numberLength := len(binaryNumbers[0])

	fmt.Println(partOne(binaryNumbers, inputSize, numberLength))
	fmt.Println(partTwo(binaryNumbers))
}

func partOne(binaryNumbers []string, inputSize int, numberLength int) int64 {
	sums := make([]int, numberLength)

	for _, binaryNumber := range binaryNumbers {
		for i := 0; i < numberLength; i++ {
			sums[i] += intAtPosition(binaryNumber, i)
		}
	}

	gammaRateString := ""
	epsilonRateString := ""

	for _, value := range sums {
		if (value > inputSize / 2) {
			gammaRateString += "1"
			epsilonRateString += "0"
		} else {
			gammaRateString += "0"
			epsilonRateString += "1"
		}
	}

	gammaRate, _ := strconv.ParseInt(gammaRateString, 2, 32);
	epsilonRate, _ := strconv.ParseInt(epsilonRateString, 2, 32);

	return gammaRate * epsilonRate
}

func partTwo(binaryNumbers []string) int64 {
	oxygenString := findOxygenRating(binaryNumbers, 0)
	co2String := findCO2Rating(binaryNumbers, 0)

	oxygen, _ := strconv.ParseInt(oxygenString, 2, 32);
	co2, _ := strconv.ParseInt(co2String, 2, 32);

	return oxygen * co2
}

func findOxygenRating(binaryNumbers []string, position int) string {
	if (len(binaryNumbers) == 1) {
		return binaryNumbers[0]
	}

	var startsWithZero []string
	var startsWithOne []string

	for _, number := range binaryNumbers {
		if (intAtPosition(number, position) == 0) {
			startsWithZero = append(startsWithZero, number)
		} else {
			startsWithOne = append(startsWithOne, number)
		}
	}

	if (len(startsWithZero) > len(startsWithOne)) {
		return findOxygenRating(startsWithZero, position + 1)
	} else {
		return findOxygenRating(startsWithOne, position + 1)
	}
}

func findCO2Rating(binaryNumbers []string, position int) string {
	if (len(binaryNumbers) == 1) {
		return binaryNumbers[0]
	}

	var startsWithZero []string
	var startsWithOne []string

	for _, number := range binaryNumbers {
		if (intAtPosition(number, position) == 0) {
			startsWithZero = append(startsWithZero, number)
		} else {
			startsWithOne = append(startsWithOne, number)
		}
	}

	if (len(startsWithZero) <= len(startsWithOne)) {
		return findCO2Rating(startsWithZero, position + 1)
	} else {
		return findCO2Rating(startsWithOne, position + 1)
	}
}

func intAtPosition(input string, position int) int {
	return util.RequireAtoI(string(input[position]))
}