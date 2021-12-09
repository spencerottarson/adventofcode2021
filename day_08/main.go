package main

import (
    "bufio"
    "fmt"
	"strings"
    
	"../util"
)

func findFirstOfLength(input []map[rune]bool, length int) map[rune]bool {
	for _, signal := range input {
		if (len(signal) == length) {
			return signal
		}
	}

	panic(1)
}

func findOne(input []map[rune]bool) map[rune]bool {
	return findFirstOfLength(input, 2)
}

func findFour(input []map[rune]bool) map[rune]bool {
	return findFirstOfLength(input, 4)
}

func findSeven(input []map[rune]bool) map[rune]bool {
	return findFirstOfLength(input, 3)
}

func findEight(input []map[rune]bool) map[rune]bool {
	return findFirstOfLength(input, 7)
}

func findThree(input []map[rune]bool, sevenSignal map[rune]bool) map[rune]bool {
	for _, signal := range input {
		if (len(signal) == 5 && isSubset(signal, sevenSignal)) {
			return signal
		}
	}

	panic(1)
}

func findNine(input []map[rune]bool, threeSignal map[rune]bool) map[rune]bool {
	for _, signal := range input {
		if (len(signal) == 6 && isSubset(signal, threeSignal)) {
			return signal
		}
	}

	panic(1)
}

func findFive(input []map[rune]bool, nineSignal map[rune]bool, threeSignal map[rune]bool) map[rune]bool {
	for _, signal := range input {
		if (len(signal) == 5 && isSubset(nineSignal, signal) && !mapEquals(signal, threeSignal)) {
			return signal
		}
	}

	panic(1)
}

func findSix(input []map[rune]bool, fiveSignal map[rune]bool, nineSignal map[rune]bool) map[rune]bool {
	for _, signal := range input {
		if (len(signal) == 6 && isSubset(signal, fiveSignal) && !mapEquals(signal, nineSignal)) {
			return signal
		}
	}

	fmt.Println(input)
	fmt.Println("Five: ", fiveSignal)
	fmt.Println("Nine:", nineSignal)

	panic(1)
}

func findZero(input []map[rune]bool, sixSignal map[rune]bool, nineSignal map[rune]bool) map[rune]bool {
	for _, signal := range input {
		if (len(signal) == 6 && !mapEquals(signal, sixSignal) && !mapEquals(signal, nineSignal)) {
			return signal
		}
	}

	panic(1)
}

func findTwo(input []map[rune]bool, threeSignal map[rune]bool, fiveSignal map[rune]bool) map[rune]bool {
	for _, signal := range input {
		if (len(signal) == 5 && !mapEquals(signal, threeSignal) && !mapEquals(signal, fiveSignal)) {
			return signal
		}
	}

	panic(1)
}

// Returns true if input2 is a subset of input1 (or if they are equivalent)
func isSubset(input1 map[rune]bool, input2 map[rune]bool) bool {
	for key, value := range input2 {
		value2, present := input1[key]
		if (!present || value != value2) {
			return false
		}
	}

	return len(input1) >= len(input2)
}

func mapEquals(input1 map[rune]bool, input2 map[rune]bool) bool {
	for key, value := range input1 {
		value2, present := input2[key]
		if (!present || value != value2) {
			return false
		}
	}

	return len(input1) == len(input2)
}

func findDigit(digitSignal map[rune]bool, digitMap map[string]map[rune]bool) string {
	for digit, signal := range digitMap {
		if (mapEquals(digitSignal, signal)) {
			return digit
		}
	}

	panic(1)
}

type Combination struct {
	signals []map[rune]bool
	outputs []map[rune]bool
}

func parseCombination(input string) Combination {
	signalStrings := strings.Split(strings.TrimSpace(strings.Split(input, "|")[0]), " ")
	var signals []map[rune]bool
	for _, value := range signalStrings {
		newMap := make(map[rune]bool)
		for _, char := range value {
			newMap[char] = true
		}
		signals = append(signals, newMap)
	}

	outputStrings := strings.Split(strings.TrimSpace(strings.Split(input, "|")[1]), " ")
	var outputs []map[rune]bool
	for _, value := range outputStrings {
		newMap := make(map[rune]bool)
		for _, char := range value {
			newMap[char] = true
		}
		outputs = append(outputs, newMap)
	}

	return Combination{signals, outputs}
}

func (c *Combination) countOutputsOfLength(length int) int {
	count := 0
	for _, output := range c.outputs {
		if (len(output) == length) {
			count++
		}
	}

	return count
}

func (c *Combination) decode() int {
	digitMap := make(map[string]map[rune]bool)
	digitMap["1"] = findOne(c.signals)
	digitMap["4"] = findFour(c.signals)
	digitMap["7"] = findSeven(c.signals)
	digitMap["8"] = findEight(c.signals)
	digitMap["3"] = findThree(c.signals, digitMap["7"])
	digitMap["9"] = findNine(c.signals, digitMap["3"])
	digitMap["5"] = findFive(c.signals, digitMap["9"], digitMap["3"])
	digitMap["6"] = findSix(c.signals, digitMap["5"], digitMap["9"])
	digitMap["0"] = findZero(c.signals, digitMap["6"], digitMap["9"])
	digitMap["2"] = findTwo(c.signals, digitMap["3"], digitMap["5"])

	outputString := ""
	for _, output := range c.outputs {
		outputString += findDigit(output, digitMap)
	}

	return util.RequireAtoI(outputString)
}

func main() {
	file := util.RequireFile("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var combinations []Combination

	for scanner.Scan() {
		inputLine := scanner.Text()
		combinations = append(combinations, parseCombination(inputLine))
	}

	sumTotal := 0
	sumPartOne := 0
	for _, combination := range combinations {
		sumTotal += combination.decode()
		sumPartOne += combination.countOutputsOfLength(2)
		sumPartOne += combination.countOutputsOfLength(4)
		sumPartOne += combination.countOutputsOfLength(3)
		sumPartOne += combination.countOutputsOfLength(7)
	}

	fmt.Println(sumPartOne)
	fmt.Println(sumTotal)
}