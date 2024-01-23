package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	emptyCell = " "
	playerX   = "X"
	playerO   = "O"
)

type gameBoard struct {
	board  [3][3]string
	player string
}

func playMove(gb gameBoard, row, col int) gameBoard {
	gb.board[row-1][col-1] = gb.player
	return gb
}

func ttteg() {
	gb := initializeGame()

	for {
		printBoard(gb)

		if gb.player == playerX {
			// Human player's turn
			fmt.Printf("Player %s, enter your move (row and column, e.g., 1 2): ", gb.player)
			var row, col int
			fmt.Scan(&row, &col)

			if isValidMove(gb, row, col) {
				gb = playMove(gb, row, col)
				if checkWin(gb) {
					printBoard(gb)
					fmt.Printf("Player %s wins!\n", gb.player)
					break
				} else if checkDraw(gb) {
					printBoard(gb)
					fmt.Println("It's a draw!")
					break
				}

				switchPlayer(&gb)
			} else {
				fmt.Println("Invalid move. Try again.")
			}
		} else {
			// Computer player's turn
			fmt.Println("Computer's turn...")
			time.Sleep(time.Duration(float64(0.5) * float64(time.Second))) // Simulate a delay for computer's move

			gb = computerPlay(gb)
			if checkWin(gb) {
				printBoard(gb)
				fmt.Printf("Player %s wins!\n", gb.player)
				break
			} else if checkDraw(gb) {
				printBoard(gb)
				fmt.Println("It's a draw!")
				break
			}

			switchPlayer(&gb)
		}
	}
}

func computerPlay(gb gameBoard) gameBoard {
	// Randomly choose an empty cell for the computer's move
	for {
		if checksetup(gb) == false {
		row := rand.Intn(3) + 1
		col := rand.Intn(3) + 1
		}
		else pcblock(gb)

		if isValidMove(gb, row, col) {
			fmt.Printf("Computer plays at %d %d\n", row, col)
			return playMove(gb, row, col)
		}
	}
}

// func pcblock (
// 	for {
// 		if checksetup(gb) == true {
// 			if setupcol(gb) == true {
				
// 			}
// 		}
// 	}
// )

func checksetup(gb gameBoard) bool {
	return setupcol(gb) || setupdiagon(gb) || setuprow(gb)
}

	
func setuprow(gb gameBoard) bool {
	for _, row := range gb.board {
		if (row[0] != emptyCell && row[0] == row[1]) || (row[1] != emptyCell && row[1] == row[2]) ||  (row[2] != emptyCell && row[2] == row[0]) {
			return true
		}
	}
	return false
}
	
func setupcol(gb gameBoard) bool {
	for i := 0; i < 3; i++ {
		if (gb.board[0][i] != emptyCell && gb.board[0][i] == gb.board[1][i]) || (gb.board[1][i] != emptyCell && gb.board[1][i] == gb.board[2][i]) || (gb.board[2][i] != emptyCell && gb.board[2][i] == gb.board[0][i]) {
			return true
		}
	}
	return false
}

func setupdiagon(gb gameBoard) bool {
	if (gb.board[0][0] != emptyCell && gb.board[0][0] == gb.board[1][1]) || (gb.board[1][1] != emptyCell && gb.board[1][1] == gb.board[2][2]) || (gb.board[2][2] != emptyCell && gb.board[2][2] == gb.board[0][0]) {
		return true
	}
	if (gb.board[0][2] != emptyCell && gb.board[0][2] == gb.board[1][1]) || (gb.board[1][1] != emptyCell && gb.board[1][1] == gb.board[2][0]) || (gb.board[2][0] != emptyCell && gb.board[2][0] == gb.board[0][2]) {
		return true
	}
	return false 
}

func initializeGame() gameBoard {
	// Initialize the board with empty cells
	var board [3][3]string
	for i := range board {
		for j := range board[i] {
			board[i][j] = emptyCell
		}
	}

	// Player X starts the game
	player := playerX
	return gameBoard{
		board:  board,
		player: player,
	}
}

func printBoard(gb gameBoard) {
	fmt.Println("  1 2 3")
	for i, row := range gb.board {
		fmt.Printf("%d ", i+1)
		for _, cell := range row {
			fmt.Printf("%s ", cell)
		}
		fmt.Println()
	}
	fmt.Println("Current player:", gb.player)
	fmt.Println()
}

func isValidMove(gb gameBoard, row, col int) bool {
	return row >= 1 && row <= 3 && col >= 1 && col <= 3 && gb.board[row-1][col-1] == emptyCell
}

func checkWin(gb gameBoard) bool {
	// Check rows, columns, and diagonals for a win
	return checkRows(gb) || checkColumns(gb) || checkDiagonals(gb)
}

func checkRows(gb gameBoard) bool {
	for _, row := range gb.board {
		if row[0] != emptyCell && row[0] == row[1] && row[1] == row[2] {
			return true
		}
	}
	return false
}

func checkColumns(gb gameBoard) bool {
	for i := 0; i < 3; i++ {
		if gb.board[0][i] != emptyCell && gb.board[0][i] == gb.board[1][i] && gb.board[1][i] == gb.board[2][i] {
			return true
		}
	}
	return false
}

func checkDiagonals(gb gameBoard) bool {
	if gb.board[0][0] != emptyCell && gb.board[0][0] == gb.board[1][1] && gb.board[1][1] == gb.board[2][2] {
		return true
	}
	if gb.board[0][2] != emptyCell && gb.board[0][2] == gb.board[1][1] && gb.board[1][1] == gb.board[2][0] {
		return true
	}
	return false
}

func checkDraw(gb gameBoard) bool {
	// Check if the board is full
	for _, row := range gb.board {
		for _, cell := range row {
			if cell == emptyCell {
				return false
			}
		}
	}
	return true
}

func switchPlayer(gb *gameBoard) {
	if gb.player == playerX {
		gb.player = playerO
	} else {
		gb.player = playerX
	}
}
