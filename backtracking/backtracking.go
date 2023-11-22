package backtracking

const dimension = 9

func solveSudokuBacktrackPlain(board [dimension][dimension]int) {
	backtrack(board)
}

func backtrack(board [dimension][dimension]int) bool {
	row, col, empty := findNextEmpty(board)
	if !empty {
		return true
	}

	for num := 1; num <= dimension; num++ {
		if !isValid(board, row, col, num) {
			continue
		}

		board[row][col] = num
		if backtrack(board) {
			return true
		}
		board[row][col] = 0
	}

	return false
}

func findNextEmpty(board [dimension][dimension]int) (int, int, bool) {
	for row := 0; row < dimension; row++ {
		for col := 0; col < dimension; col++ {
			if board[row][col] == 0 {
				return row, col, true
			}
		}
	}

	return 0, 0, false
}

func isValid(board [dimension][dimension]int, row, col, num int) bool {
	return isValidRow(board, row, num) &&
		isValidCol(board, col, num) &&
		isValidSubGrid(board, row, col, num)
}

func isValidRow(board [dimension][dimension]int, row, num int) bool {
	for col := 0; col < dimension; col++ {
		if board[row][col] == num {
			return false
		}
	}

	return true
}

func isValidCol(board [dimension][dimension]int, col, num int) bool {
	for row := 0; row < dimension; row++ {
		if board[row][col] == num {
			return false
		}
	}

	return true
}

func isValidSubGrid(board [dimension][dimension]int, row, col, num int) bool {
	rowStart, colStart := row-row%3, col-col%3

	for row := rowStart; row < rowStart+3; row++ {
		for col := colStart; col < colStart+3; col++ {
			if board[row][col] == num {
				return false
			}
		}
	}

	return true
}

func solveSudokuBacktrackConstraintPropagationNakedSingle(board [dimension][dimension]int) {
	backtrackConstraintPropagation(board)
}

func backtrackConstraintPropagation(board [dimension][dimension]int) bool {
	row, col, empty := findNextEmpty(board)
	if !empty {
		return true
	}

	possible := getPossibleValues(board, row, col)
	for _, num := range possible {
		board[row][col] = num
		if backtrackConstraintPropagation(board) {
			return true
		}
		board[row][col] = 0
	}

	return false
}

func getPossibleValues(board [dimension][dimension]int, row, col int) []int {
	possible := make([]bool, 0, dimension+1)
	for i := 0; i <= dimension; i++ {
		possible = append(possible, true)
	}

	// Eliminate values based on row and column.
	for i := 0; i < dimension; i++ {
		possible[board[row][i]] = false
		possible[board[i][col]] = false
	}

	// Eliminate values based on subgrid.
	rowStart, colStart := row-row%3, col-col%3
	for row := rowStart; row < rowStart+3; row++ {
		for col := colStart; col < colStart+3; col++ {
			possible[board[row][col]] = false
		}
	}

	// Add possible values to slice.
	result := make([]int, 0, dimension)
	for i := 0; i <= dimension; i++ {
		if possible[i] {
			result = append(result, i)
		}
	}

	return result
}

func solveSudokuBacktrackConstraintPropagationHiddenSingles(board [dimension][dimension]int) {
	backtrackConstraintPropagationHiddenSingles(board)
}

func backtrackConstraintPropagationHiddenSingles(board [dimension][dimension]int) bool {
	row, col, empty := findNextEmpty(board)
	if !empty {
		return true
	}

	possible := getPossibleValues(board, row, col)
	if len(possible) == 1 {
		board[row][col] = possible[0]
		return backtrackConstraintPropagationHiddenSingles(board)
	}

	for _, num := range possible {
		board[row][col] = num
		if backtrackConstraintPropagationHiddenSingles(board) {
			return true
		}
		board[row][col] = 0
	}

	return false
}

func isSolved(board [dimension][dimension]int) bool {
	for row := 0; row < dimension; row++ {
		for col := 0; col < dimension; col++ {
			if board[row][col] == 0 {
				return false
			}
		}
	}

	return true
}

func applyHiddenSingles(board [dimension][dimension]int) bool {
	changed := false

	for row := 0; row < dimension; row++ {
		for col := 0; col < dimension; col++ {
			if board[row][col] != 0 {
				continue
			}

			possible := getPossibleValues(board, row, col)
			if len(possible) == 1 {
				board[row][col] = possible[0]
				changed = true
			}
		}
	}

	return changed
}
