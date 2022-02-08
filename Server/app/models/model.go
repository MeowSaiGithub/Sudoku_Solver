package models

type Sudoku struct {
	Board [9][9]int `json:"Board"`
}

type ResponseSudoku struct {
	Status  string `json:"Status"`
	RSudoku Sudoku `json:"RSudoku"`
}
