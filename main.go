package main

import "fmt"

func main() {
	board := [9][9]int{
		{0, 0, 0, 5, 0, 0, 4, 2, 0},
		{0, 5, 0, 0, 0, 9, 6, 0, 0},
		{6, 8, 7, 0, 0, 0, 0, 1, 5},
		{0, 0, 9, 6, 5, 8, 1, 3, 2},
		{0, 0, 2, 0, 4, 0, 0, 0, 8},
		{0, 0, 0, 0, 9, 1, 0, 6, 4},
		{3, 0, 0, 0, 0, 2, 0, 0, 0},
		{7, 2, 0, 0, 1, 0, 3, 4, 9},
		{8, 9, 1, 0, 0, 7, 0, 5, 0},
	}

	history := [][9][9]int{}

	backtracking(&board, &history)

	fmt.Printf("%v", board)
}

func backtracking(board *[9][9]int, history *[][9][9]int) bool {
	// Store each step to visualise the process of the algorithm
	*history = append(*history, *board)

	// Base case
	if !hasEmptyCell(board) {
		return true
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				candidates := getCandidates(board, i, j)

				if len(candidates) != 0 {
					for _, value := range candidates {
						board[i][j] = value

						if backtracking(board, history) {
							return true
						} else {
							board[i][j] = 0
						}
					}
				}

				// Current search path failed
				return false
			}
		}
	}

	// Current search path failed
	return false
}

func getCandidates(board *[9][9]int, rowIndex int, columnIndex int) []int {
	candidates := []int{}

	for i := 1; i < 10; i++ {
		// Check for duplicate in row
		if contains(board[rowIndex], i) {
			continue
		}

		// Check for duplicate in column
		if isColumnDuplicate(board, columnIndex, i) {
			continue
		}

		// Check for duplicate in 3x3 subgrid
		if isSubgridDuplicate(board, rowIndex, columnIndex, i) {
			continue
		}

		candidates = append(candidates, i)
	}

	return candidates
}

func isColumnDuplicate(board *[9][9]int, columnIndex int, i int) bool {
	for j := 0; j < 9; j++ {
		v := board[j][columnIndex]

		if v != 0 && v == i {
			return true
		}
	}

	return false
}

func isSubgridDuplicate(board *[9][9]int, rowIndex int, columnIndex int, v int) bool {
	rowBound := (rowIndex / 3) * 3
	columnBound := (columnIndex / 3) * 3

	for i := rowBound; i < rowBound+3; i++ {
		for j := columnBound; j < columnBound+3; j++ {
			if board[i][j] == v {
				return true
			}
		}
	}

	return false
}

func hasEmptyCell(board *[9][9]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				return true
			}
		}
	}

	return false
}

func contains(i [9]int, int int) bool {
	for _, v := range i {
		if v == int {
			return true
		}
	}

	return false
}
