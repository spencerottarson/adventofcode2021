package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
	"strings"
)

type movement struct {
    direction string
    value  int
}


func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var movements []movement

	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		
		direction := input[0]
		value, err := strconv.Atoi(input[1])

		movement := movement{direction, value}

		if err != nil {
			log.Fatal(err)
		}

		movements = append(movements, movement)
	}

	fmt.Println(partOne(movements))
	fmt.Println(partTwo(movements))
}

func partOne(movements []movement) int {
	horizontalPosition := 0
	verticalPosition := 0

	for _, move := range movements {
		switch move.direction {
		case "forward":
			horizontalPosition += move.value
		case "up":
			verticalPosition -= move.value
		case "down":
			verticalPosition += move.value
		}
	}

	return verticalPosition * horizontalPosition
}

func partTwo(movements []movement) int {
	horizontalPosition := 0
	verticalPosition := 0
	aim := 0

	for _, move := range movements {
		switch move.direction {
		case "forward":
			horizontalPosition += move.value
			verticalPosition += move.value * aim
		case "up":
			aim -= move.value
		case "down":
			aim += move.value
		}
	}

	return verticalPosition * horizontalPosition
}