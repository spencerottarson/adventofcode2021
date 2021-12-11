package main

import (
    "bufio"
    "fmt"
    
	"../util"
)

func increaseAll(grid [][]util.Point) {
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			grid[row][col].Value++
		}
	}
}

func increaseFromFlashes(grid [][]util.Point) {
	stillToSearch := util.NewPointSet()
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col].Value == 10 {
				stillToSearch.Add(&grid[row][col])
			}
		}
	}

	for !stillToSearch.IsEmpty() {
		point := stillToSearch.Pop()
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				// Skip the current point and any points off the grid
				if (i == 0 && j == 0) || point.Row + i < 0 || point.Row + i >= len(grid) || point.Col + j < 0 || point.Col + j >= len(grid[point.Row]) {
					continue
				}

				grid[point.Row + i][point.Col + j].Value++
				if grid[point.Row + i][point.Col + j].Value == 10 {
					stillToSearch.Add(&grid[point.Row + i][point.Col + j])
				}
			}
		}
	}
}

func resetAll(grid [][]util.Point) int {
	countFlashes := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col].Value > 9 {
				countFlashes++
				grid[row][col].Value = 0
			}
		}
	}

	return countFlashes
}

func main() {
	file := util.RequireFile("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var grid [][]util.Point

	for row := 0; scanner.Scan(); row++ {
		line := scanner.Text()
		var lineValues []util.Point
		for col, char := range line {
			lineValues = append(lineValues, util.Point{row, col, int(char - '0')})
		}
		grid = append(grid, lineValues)
	}

	totalFlashes := 0
	numberOfOctopi := len(grid) * len(grid[0])
	stepCounter := 0
	flashesThisStep := -1
	for flashesThisStep != numberOfOctopi {  // Assumes it will take more than 100 steps to synchronize
		increaseAll(grid)
		increaseFromFlashes(grid)

		flashesThisStep = resetAll(grid)
		totalFlashes += flashesThisStep
		if stepCounter == 99 {
			fmt.Println("Flashes at 100: ", totalFlashes)
		}
		stepCounter++
	}

	fmt.Println("Steps to synchronize: ", stepCounter)
}
