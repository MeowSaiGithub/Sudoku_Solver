This app is implemented with RESTAPI and microservice(grpc) using golang language to solve Sudokus.
RESTAPI model is to use `POST` method. URL is `localhost:4040/sudoku-solver`
The body is 
`{
    "Board": [[1,0,8,4,0,0,9,0,0],[4,0,6,0,0,0,0,0,0],[0,5,0,0,8,0,7,0,0],[0,9,0,0,3,0,0,0,2],[0,0,0,0,4,0,0,0,6],[2,6,0,5,0,0,0,0,1],[0,0,0,0,0,0,0,0,0],[0,8,0,0,0,1,0,0,9],[0,2,0,0,5,0,3,0,0]]
}
`
The server will respond with status which is `Sudoku was successfully` or `The sudoku cannot be solved!!!` along with solved sudoku or inputed sudoku.
Microservice model is run on port `50051` the input and output will be the same as RESTAPI one.
The server also log steps and show the solved sudoku at the terminal the server is running.
