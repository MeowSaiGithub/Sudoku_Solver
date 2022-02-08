package handlers

import (
	"Sudoku_Solver/Server/app/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var response = models.ResponseSudoku{}

func Solve_Sudoku(w http.ResponseWriter, r *http.Request) {
	log.Print("Started solving sudoku")

	sudoku := models.Sudoku{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&sudoku)

	if err != nil {
		fmt.Print("Error after decoder")
		log.Fatal(err)
	}

	printBoard(sudoku.Board)
	if backtrack(&sudoku.Board) {
		log.Println("The Sudoku was solved successfully")
		printBoard(sudoku.Board)
		response.Status = "The Sudoku was solved successfully"
		response.RSudoku = sudoku
		respondJSON(w, http.StatusOK, response)
	} else {
		fmt.Printf("The Sudoku can't be solved!!!")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("The Sudoku can't be solved!!!"))
	}
}
