package main

import (
	"Sudoku_Solver/Server/app"
	"Sudoku_Solver/Server/app/handlers"
	"fmt"
	"log"
)

func main() {

	fmt.Println("Select Mode")
	fmt.Println("Enter 1 for REST API Mode")
	fmt.Println("Enter 2 for Microservice Mode")
	fmt.Println("Enter 0 to Exit")
	var choice string
	fmt.Scanln(&choice)
	for choice != "0" {
		switch choice {
		case "1":
			log.Println("RESTAPI Mode")
			application := &app.App{}
			application.Initialize()
			application.Run(":4040")
		case "2":
			log.Println("Microservice Mode")
			handlers.StartServer()
		case "0":
			log.Println("Exiting")
		default:
			fmt.Println("Invalid inputs. Please enter again!! Enter 0 to Exit")
			fmt.Scanln(&choice)
		}

	}

}
