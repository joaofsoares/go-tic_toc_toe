package main

import (
	"fmt"
)

func main() {
	var p1 string
	fmt.Println("Enter the name of the first player:")
	fmt.Scanf("%s", &p1)

	var p2 string
	fmt.Println("Enter the name of the second player:")
	fmt.Scanf("%s", &p2)

	space := [3][3]int{{'1', '2', '3'}, {'4', '5', '6'}, {'7', '8', '9'}}

	var row int
	var column int
	var tie bool
	token := 'X'

	for !checkWinner(space, &tie) {
		printBoard(space)
		assignPosition(p1, p2, &space, &token, &row, &column)
		checkWinner(space, &tie)
	}

	if token == 'X' && tie == false {
		fmt.Printf("Player 1 - %s Wins", p1)
	} else if token == '0' && tie == false {
		fmt.Printf("Player 2 - %s Wins", p2)
	} else {
		fmt.Println("It's a draw!")
	}
}

func printBoard(space [3][3]int) {
	fmt.Println("     |     |     ")
	fmt.Printf("  %c  |  %c  |  %c  \n", space[0][0], space[0][1], space[0][2])
	fmt.Println("_____|_____|_____")
	fmt.Println("     |     |     ")
	fmt.Printf("  %c  |  %c  |  %c  \n", space[1][0], space[1][1], space[1][2])
	fmt.Println("_____|_____|_____")
	fmt.Println("     |     |     ")
	fmt.Printf("  %c  |  %c  |  %c  \n", space[2][0], space[2][1], space[2][2])
	fmt.Println("     |     |     ")
	fmt.Println()
}

func assignPosition(p1 string, p2 string, space *[3][3]int, token *rune, row *int, column *int) {
	var digit int

	if *token == 'X' {
		fmt.Printf("Player 1 - %s - insert digit: ", p1)
		fmt.Scanf("%d", &digit)
	}

	if *token == '0' {
		fmt.Printf("Player 2 - %s - insert digit: ", p2)
		fmt.Scanf("%d", &digit)
	}

	if digit == 1 {
		*row = 0
		*column = 0
	}

	if digit == 2 {
		*row = 0
		*column = 1
	}

	if digit == 3 {
		*row = 0
		*column = 2
	}

	if digit == 4 {
		*row = 1
		*column = 0
	}

	if digit == 5 {
		*row = 1
		*column = 1
	}

	if digit == 6 {
		*row = 1
		*column = 2
	}

	if digit == 7 {
		*row = 2
		*column = 0
	}

	if digit == 8 {
		*row = 2
		*column = 1
	}

	if digit == 9 {
		*row = 2
		*column = 2
	}

	if digit < 1 || digit > 9 {
		fmt.Println("Invalid!")
	} else if *token == 'X' && (*space)[*row][*column] != 'X' && (*space)[*row][*column] != '0' {
		(*space)[*row][*column] = 'X'
		*token = '0'
	} else if *token == '0' && (*space)[*row][*column] != 'X' && (*space)[*row][*column] != '0' {
		(*space)[*row][*column] = '0'
		*token = 'X'
	} else {
		fmt.Println("There is no empty space")
		assignPosition(p1, p2, space, token, row, column)
	}

	printBoard(*space)
}

func checkWinner(space [3][3]int, tie *bool) bool {
	for i := 0; i < 3; i += 1 {
		if (space[i][0] == space[i][1] && space[i][0] == space[i][2]) ||
			(space[0][i] == space[1][i] && space[0][i] == space[2][i]) {
			return true
		}
	}

	if (space[0][0] == space[1][1] && space[1][1] == space[2][2]) ||
		(space[0][2] == space[1][1] && space[1][1] == space[2][0]) {
		return true
	}

	for i := 0; i < 3; i += 1 {
		for j := 0; j < 3; j += 1 {
			if space[i][j] != 'X' && space[i][j] != '0' {
				return false
			}
		}
	}

	*tie = true

	return false
}
