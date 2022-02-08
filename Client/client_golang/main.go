package main

import (
	pb "Sudoku_Solver/Client/client_golang/sudoku-solver"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	Board := [9][9]int{{1, 0, 8, 4, 0, 0, 9, 0, 0}, {4, 0, 6, 0, 0, 0, 0, 0, 0}, {0, 5, 0, 0, 8, 0, 7, 0, 0}, {0, 9, 0, 0, 3, 0, 0, 0, 2}, {0, 0, 0, 0, 4, 0, 0, 0, 6}, {2, 6, 0, 5, 0, 0, 0, 0, 1}, {0, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 8, 0, 0, 0, 1, 0, 0, 9}, {0, 2, 0, 0, 5, 0, 3, 0, 0}}
	printBoard(Board)
	SBoard, _ := json.Marshal(Board)
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSolverClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()

	r, err := c.SolveSudoku(ctx, &pb.SudokuRequest{Sudoku: SBoard})
	if err != nil {
		log.Fatalf("Could not send sudoku: %v", err)
	}
	log.Printf("%s", r.GetStatus())
	var RBoard [9][9]int
	if err := json.Unmarshal(r.GetSudoku(), &RBoard); err != nil {
		log.Panic(err)
	}
	log.Printf("The answers is %v", RBoard)
	printBoard(RBoard)
}

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
