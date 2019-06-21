package main

import (
  "fmt"
)

func main() {
    boardSize := 9
    board := initializeBoard(boardSize)
    printBoard(board, boardSize)
}

func initializeBoard(boardSize int)([][]string) {
    board := make([][]string, boardSize)
    for i := 0; i < boardSize; i++ {
        board[i] = make([]string, boardSize)
        for j := 0; j < boardSize; j++ {
            board[i][j] = " "
        }
    }
    return board
}


func printBoard(board [][]string, boardSize int) {
    for i := 0; i < boardSize*2 + 1; i++ {
        fmt.Print("-")
    }
    fmt.Print("\n")
    for i := 0; i < boardSize; i++ {
        for j := 0; j < boardSize; j++ {
            fmt.Print("|", board[i][j])
        }
        fmt.Print("|\n")
    }
    for i := 0; i < boardSize*2 + 1; i++ {
        fmt.Print("-")
    }
    fmt.Print("\n")
}

