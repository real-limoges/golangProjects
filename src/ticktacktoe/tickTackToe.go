package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
    "math"
)

func main() {
    const boardSize int = 3
    board := initializeBoard(boardSize)
    printBoard(board, boardSize)

    turn := 0

    for {
        fmt.Println("Enter X Coordinate for Piece")
        fmt.Println("Please Enter X Coordinate")
        x := receiveInput(boardSize)
        fmt.Println("Please Enter Y Coordinate")
        y := receiveInput(boardSize)

        var piece string
        if turn % 2 == 0  {
            piece = "X"
        } else {
            piece = "O"
        }

        valid := placePiece(board, piece, x, y)

        if valid {
            board[x][y] = piece
            victory := checkVictory(board, boardSize, piece)
            if victory {
                fmt.Println("Congrats Piece ", piece)
            }
            turn += 1
            break
        }
    }
    tie := checkTie(board, boardSize)
    if tie {
        fmt.Println("Tie Game. Bummer")
    }
    printBoard(board, boardSize)
}

func initializeBoard(boardSize int) ([][]string){
    board := make([][]string, boardSize)
    for i := 0; i < boardSize; i++ {
        board[i] = make([]string, boardSize)
        for j := 0; j < boardSize; j++ {
            board[i][j] = " "
        }
    }
    return board
}

func printBoard(board [][]string, boardSize int){

    for i:= 0; i < boardSize*2 + 1; i++ {
        fmt.Print("-")
    }
    fmt.Println()
    for i := 0; i < boardSize; i++ {
        for j := 0; j < boardSize; j++ {
            fmt.Print("|", board[i][j])
        }
        fmt.Print("|\n")
    }
    for i := 0; i < boardSize*2 + 1; i++ {
        fmt.Print("-")
    }
    fmt.Println()
}

func placePiece(board [][]string, piece string, x int, y int) (bool) {
    if board[x][y] == " " {
        return true
    }
    return false
}

func receiveInput(boardSize int) (int) {
    for {
        in := bufio.NewScanner(os.Stdin)
        in.Scan()
        val, err := strconv.Atoi(in.Text())
        if err != nil || val < 0 || val > boardSize - 1 {
            fmt.Println("Please enter a valid input")
        } else {
            return val
        }
    }
}

func checkVictory(board [][]string, boardSize int, piece string) (bool){
    for i := 0; i < boardSize; i++ {
        h_sum := 0
        v_sum := 0
        for j := 0; j < boardSize; j++ {
            if board[i][j] == piece {
                h_sum += 1
            }
            if board[j][i] == piece {
                v_sum += 1
            }
        }
        if h_sum == boardSize || v_sum == boardSize{
            return true
        }
    }

    d1_sum := 0
    for i := 0; i < boardSize; i++ {
        if board[i][i] == piece {
            d1_sum += 1
        }
    }
    if d1_sum == boardSize {
        return true
    }

    d2_sum := 0
    for i := 0; i < boardSize; i++ {
        j := boardSize -1 - i
        if board[i][j] == piece {
            d2_sum += 1
        }
    }
    if d2_sum == boardSize {
        return true
    }
    return false
}

func checkTie(board [][]string, boardSize int) (bool) {
    sum := 0
    board_2 := int(math.Pow(float64(boardSize), 2))
    for i := 0; i < boardSize; i++ {
        for j := 0; j < boardSize; j++ {
            if board[i][j] != " " {
                sum += 1
            }
        }
    }
    if sum == board_2 {
        return true
    }
    return false
}
