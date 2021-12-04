package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

func main() {
	file, err := os.Open("input1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var values []int

	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		values = append(values, value)
	}

	fmt.Println(countIncreasesSliding(values, 1))
	fmt.Println(countIncreasesSliding(values, 3))
}

func countIncreases(values []int, windowSize int) int {
	count := 0

	for i, s := range values {
		if i < windowSize {
			continue
		}

		// We can ignore all of the values which are shared 
		// by the current window and the previous window, 
		// as they will have the same impact on both windows.
		// We only need to check if the new added value
		// is greater than the one falling out
		if s > values[i - windowSize] {
			count++
		}
	}

	return count
}

