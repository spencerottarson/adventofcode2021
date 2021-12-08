package main

import (
    "bufio"
    "fmt"
	"strings"

	"../util"
)

type Point struct {
    x int
	y int
}

type Line struct {
	startPoint Point
	endPoint Point
}

func parseLine(input string) Line {
	stringPoints := strings.Split(input, " -> ")
	startPoint := Point{util.RequireAtoI(strings.Split(stringPoints[0], ",")[0]), util.RequireAtoI(strings.Split(stringPoints[0], ",")[1])}
	endPoint := Point{util.RequireAtoI(strings.Split(stringPoints[1], ",")[0]), util.RequireAtoI(strings.Split(stringPoints[1], ",")[1])}

	return Line{startPoint, endPoint}
}

func getSign(number int) int {
	if (number > 0) {
		return 1
	} else if (number < 0) {
		return -1
	} else {
		return 0
	}
}

func (line *Line) getPoints(includeDiags bool) []Point {
	var points []Point

	if (!includeDiags && line.startPoint.x != line.endPoint.x && line.startPoint.y != line.endPoint.y) {
		return points
	}

	xSign := getSign(line.endPoint.x - line.startPoint.x)
	ySign := getSign(line.endPoint.y - line.startPoint.y)

	var xValues []int
	var yValues []int

	if (xSign != 0) {
		for x := line.startPoint.x; x != line.endPoint.x; x += xSign {
			xValues = append(xValues, x)
		}
	} else {
		for i := 0; i < (line.endPoint.y - line.startPoint.y) * ySign; i++ {
			xValues = append(xValues, line.startPoint.x)
		}
	}

	xValues = append(xValues, line.endPoint.x)

	if (ySign != 0) {
		for y := line.startPoint.y; y != line.endPoint.y; y += ySign {
			yValues = append(yValues, y)
		}
	} else {
		for i := 0; i < (line.endPoint.x - line.startPoint.x) * xSign; i++ {
			yValues = append(yValues, line.startPoint.y)
		}
	}

	yValues = append(yValues, line.endPoint.y)

	for i := 0; i < len(xValues); i++ {
		points = append(points, Point{xValues[i], yValues[i]})
	}

	return points
}

func main() {
	file := util.RequireFile("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []Line

	for scanner.Scan() {
		line := parseLine(scanner.Text())
		lines = append(lines, line)
	}

	fmt.Println("Part One: ", computeOverlap(lines, false))
	fmt.Println("Part Two: ", computeOverlap(lines, true))
}

func computeOverlap(lines []Line, includeDiags bool) int {
	points := make(map[Point]int)

	for _, line := range lines {
		linePoints := line.getPoints(includeDiags)

		for _, point := range linePoints {
			points[point]++
		}
	}

	countOverlap := 0
	for _, value := range points {
		if (value > 1) {
			countOverlap++
		}
	}

	return countOverlap
}