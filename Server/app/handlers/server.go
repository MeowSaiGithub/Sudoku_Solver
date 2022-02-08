package handlers

import (
	pb "Sudoku_Solver/Server/app/micro"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedSolverServer
}

func (s *server) SolveSudoku(ctx context.Context, in *pb.SudokuRequest) (*pb.SudokuReply, error) {
	log.Printf("Recieved: %v", in.GetSudoku())

	var data [9][9]int
	if err := json.Unmarshal(in.GetSudoku(), &data); err != nil {
		log.Panic(err)
	}
	if backtrack(&data) {
		log.Println("The Sudoku was solved successfully")
		log.Printf("Converted into array : %v", data)
		printBoard(data)
		solved, _ := json.Marshal(data)
		return &pb.SudokuReply{Status: "Successfully Solved", Sudoku: solved}, nil
	}
	log.Println("The Sudoku was not abled to solved!!!")
	return &pb.SudokuReply{Status: "Cannot be solved", Sudoku: in.GetSudoku()}, nil
}

func StartServer() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSolverServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
