package main

import (
    "bufio"
    "fmt"
	"sort"
    
	"../util"
)

type Point struct {
	Row int
	Col int
	Height int
}

type PointSet struct {
	points map[*Point]bool
}

func CreatePointSet() PointSet {
	return PointSet{make(map[*Point]bool)}
}

func (s *PointSet) add(point *Point) {
	s.points[point] = true
}

func (s *PointSet) remove(point *Point) {
	s.points[point] = false
}

func (s *PointSet) pop() *Point {
	for point, value := range s.points {
		if value == true {
			s.points[point] = false
			return point
		}
	}

	return nil
}

func (s *PointSet) contains(point *Point) bool {
	return s.points[point] == true
}

func (s *PointSet) size() int {
	count := 0
	for _, value := range s.points {
		if value == true {
			count++
		}
	}

	return count
}

func (s *PointSet) isEmpty() bool {
	for _, value := range s.points {
		if value == true {
			return false
		}
	}

	return true
}

func isLowPoint(values [][]Point, row int, col int) bool {
	lowerThanLeft := false
	if (col == 0 || values[row][col].Height < values[row][col-1].Height) {
		lowerThanLeft = true
	}

	lowerThanRight := false
	if (col == len(values[row]) - 1 || values[row][col].Height < values[row][col+1].Height) {
		lowerThanRight = true
	}

	lowerThanAbove := false
	if (row == 0 || values[row][col].Height < values[row-1][col].Height) {
		lowerThanAbove = true
	}

	lowerThanBelow := false
	if (row == len(values) - 1  || values[row][col].Height < values[row+1][col].Height) {
		lowerThanBelow = true
	}

	return lowerThanLeft && lowerThanRight && lowerThanAbove && lowerThanBelow
}

func findBasinSize(values [][]Point, lowPoint *Point) int {
	stillToSearch := CreatePointSet()
	visited := CreatePointSet()
	stillToSearch.add(lowPoint)
	
	for (!stillToSearch.isEmpty()) {
		point := stillToSearch.pop()
		visited.add(point)

		// left one
		if (point.Col > 0 && values[point.Row][point.Col - 1].Height < 9 && !visited.contains(&values[point.Row][point.Col - 1])) {
			stillToSearch.add(&values[point.Row][point.Col - 1])
		}
		// right one
		if (point.Col < len(values[point.Row]) - 1 && values[point.Row][point.Col + 1].Height < 9 && !visited.contains(&values[point.Row][point.Col + 1])) {
			stillToSearch.add(&values[point.Row][point.Col + 1])
		}

		// up one
		if (point.Row > 0 && values[point.Row - 1][point.Col].Height < 9 && !visited.contains(&values[point.Row - 1][point.Col])) {
			stillToSearch.add(&values[point.Row - 1][point.Col])
		}
		// down one
		if (point.Row < len(values) - 1 && values[point.Row + 1][point.Col].Height < 9 && !visited.contains(&values[point.Row + 1][point.Col])) {
			stillToSearch.add(&values[point.Row + 1][point.Col])
		}
	}

	return visited.size()
}

func findLowPoints(values [][]Point) []*Point {
	var lowPoints []*Point
	for row := 0; row < len(values); row++ {
		for col := 0; col < len(values[row]); col++ {
			if (isLowPoint(values, row, col)) {
				lowPoints = append(lowPoints, &values[row][col])
			}
		}
	}

	return lowPoints
}

func main() {
	file := util.RequireFile("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var values [][]Point

	for row := 0; scanner.Scan(); row++ {
		line := scanner.Text()
		var lineValues []Point
		for col, char := range line {
			lineValues = append(lineValues, Point{row, col, int(char - '0')})
		}
		values = append(values, lineValues)
	}

	lowPoints := findLowPoints(values)

	riskLevel := 0
	var basinSizes []int
	for _, point := range lowPoints {
		riskLevel += point.Height + 1
		basinSizes = append(basinSizes, findBasinSize(values, point))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))

	fmt.Println(riskLevel)
	fmt.Println(basinSizes[0] * basinSizes[1] * basinSizes[2])
}