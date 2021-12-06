package main

import (
    "bufio"
    "fmt"
	"strings"

	"../util"
)

type Finish struct {
	board int
	winningNumber int
}

type Space struct {
	value int
	filled bool
}

type Board struct {
	spaces [][]Space
}

func (board *Board) addLine(input string) {
	inputs := strings.Split(input, " ")
	var row []Space
	for _, value := range inputs {
		if (strings.TrimSpace(value) != "") {
			row = append(row, Space{util.RequireAtoI(value), false})
		}
	}
	board.spaces = append(board.spaces, row)
}

func (board *Board) markSpace(value int) (bool, int, int) {
	for row := 0; row < len(board.spaces); row++ {
		for col := 0; col < len(board.spaces[row]); col++ {
			if (board.spaces[row][col].value == value) {
				board.spaces[row][col].filled = true
				return true, row, col
			}
		}
	}

	return false, 0, 0
}

func (board *Board) checkRowForWin(row int) bool {
	for col := 0; col < len(board.spaces[row]); col++ {
		if (board.spaces[row][col].filled == false) {
			return false
		}
	}

	return true
}

func (board *Board) checkColForWin(col int) bool {
	for row := 0; row < len(board.spaces); row++ {
		if (board.spaces[row][col].filled == false) {
			return false
		}
	}

	return true
}

func (board *Board) sumUnfilledSpaces() int {
	sum := 0
	for row := 0; row < len(board.spaces); row++ {
		for col := 0; col < len(board.spaces[row]); col++ {
			if (board.spaces[row][col].filled == false) {
				sum += board.spaces[row][col].value
			}
		}
	}

	return sum
}

func containsBoard(finishers []Finish, boardNumber int) bool {
    for _, finish := range finishers {
        if finish.board == boardNumber {
            return true
        }
    }
    return false
}

func main() {
	file := util.RequireFile("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	drawnNumbersString := scanner.Text()

	var drawnNumbers []int

	for _, value := range strings.Split(drawnNumbersString, ",") {
		drawnNumbers = append(drawnNumbers, util.RequireAtoI(value))
	}

	var boards []Board

	for scanner.Scan() {
		scanner.Text()
		var board Board
		for i := 0; i < 5; i++ {
			scanner.Scan()
			inputLine  := scanner.Text()
			board.addLine(inputLine)
		}
		boards = append(boards, board)
	}

	orderOfFinish := findWinningBoards(boards, drawnNumbers)

	fmt.Println(orderOfFinish)
	sumWinner := boards[orderOfFinish[0].board].sumUnfilledSpaces()
	sumLast := boards[orderOfFinish[len(orderOfFinish) - 1].board].sumUnfilledSpaces()

	fmt.Println("Part 1: ", sumWinner * orderOfFinish[0].winningNumber)
	fmt.Println("Part 2: ", sumLast * orderOfFinish[len(orderOfFinish) - 1].winningNumber)
}

func findWinningBoards(boards []Board, drawnNumbers []int) []Finish {
	var finishers []Finish
	for _, number := range drawnNumbers {
		for boardNumber, board := range boards {
			if (containsBoard(finishers, boardNumber)) {
				continue
			}

			status, row, col := board.markSpace(number)
			if (status && board.checkRowForWin(row) || board.checkColForWin(col)) {
				finishers = append(finishers, Finish{boardNumber, number})
			}
		}
	}

	return finishers
}