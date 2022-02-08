package app

import (
	"Sudoku_Solver/Server/app/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.setRouter()
}

func (a *App) setRouter() {
	a.Post("/sudoku-solver", a.solve_Sudoku)
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

func (a *App) solve_Sudoku(w http.ResponseWriter, r *http.Request) {
	handlers.Solve_Sudoku(w, r)
}
