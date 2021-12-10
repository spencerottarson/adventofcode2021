package main

import (
    "bufio"
    "fmt"
	"sort"
    
	"../util"
)

var Points = map[rune]int {
    ')': 3,
    ']': 57,
    '}': 1197,
    '>': 25137,
}

var PartTwoPoints = map[rune]int {
    '(': 1,
    '[': 2,
    '{': 3,
    '<': 4,
}

var Openers = map[rune]bool {
    '(': true,
    '[': true,
    '{': true,
    '<': true,
}

var ClosingMatch = map[rune]rune {
    '(': ')',
    '[': ']',
    '{': '}',
    '<': '>',
}

func checkString(input string) (rune, int, bool) {
	stack := NewStack()
	for _, char := range input {
		if Openers[char] {
			stack.Add(char)
		} else if ClosingMatch[stack.Peek()] == char {
			stack.Pop()
		} else {
			return char, 0, false
		}
	}

	finishingScore := 0
	for stack.Size() > 0 {
		top := stack.Pop()
		finishingScore *= 5
		finishingScore += PartTwoPoints[top]
	}

	return 0, finishingScore, true
}

func main() {
	file := util.RequireFile("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	sum := 0
	var finishingScores []int
	for _, line := range lines {
		illegalChar, finishingScore, isValid := checkString(line)
		if isValid {
			finishingScores = append(finishingScores, finishingScore)
		} else {
			sum += Points[illegalChar]
		}
	}

	sort.Ints(finishingScores)

	fmt.Println(sum)
	fmt.Println(finishingScores[len(finishingScores) / 2])
}