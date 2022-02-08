package handlers

import (
	"Sudoku_Solver/Server/app/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func printBoard(board [9][9]int) {
	fmt.Println()
	fmt.Println("+-------+-------+-------+")
	for row := 0; row < 9; row++ {
		fmt.Print("| ")
		for col := 0; col < 9; col++ {
			if col == 3 || col == 6 {
				fmt.Print("| ")
			}
			fmt.Printf("%d ", board[row][col])
			if col == 8 {
				fmt.Print("|")
			}
		}
		if row == 2 || row == 5 || row == 8 {
			fmt.Println("\n+-------+-------+-------+")
		} else {
			fmt.Println()
		}
	}
}

func isBoardValid(board *[9][9]int) bool {
	//check duplicate by row
	for row := 0; row < 9; row++ {
		counter := [10]int{}
		for column := 0; column < 9; column++ {
			counter[board[row][column]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}

	// check duplicates by column
	for row := 0; row < 9; row++ {
		counter := [10]int{}
		for column := 0; column < 9; column++ {
			counter[board[column][row]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}

	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			counter := [10]int{}
			for row := i; row < i+3; row++ {
				for column := j; column < j+3; column++ {
					counter[board[row][column]]++
				}
				if hasDuplicates(counter) {
					return false
				}
			}
		}
	}
	//no duplicates return true
	return true
}

func hasDuplicates(counter [10]int) bool {
	for i, count := range counter {
		if i == 0 {
			continue
		}
		if count > 1 {
			return true
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

func respondJSON(w http.ResponseWriter, status int, payload models.ResponseSudoku) {
	respond, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(respond))
}
