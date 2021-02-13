package main

import (
	"fmt"
	"os"
	"strconv"
)

type gameField [3][3]string

type playerSign struct {
	PlayerName string
	PlayerSign string
}

var gf = initGameField()
var p1 = playerSign{PlayerName: "Player1"}
var p2 = playerSign{PlayerName: "Player2"}

func menu() {
	var input string
	fmt.Println("Choose game type")
	fmt.Println("1. Single-player")
	fmt.Println("2. Multi-player")
	fmt.Println("0. Exit")
	fmt.Scanln(&input)
	val, err := strconv.Atoi(input)
	if err != nil {
		wrongOption()
	}
	switch val {
	case 0:
		os.Exit(0)
	case 1:
		singlePlayer()
	case 2:
		multiPlayer()
	default:
		wrongOption()
	}
}

func wrongOption() {
	fmt.Println("\n Wrong option")
	menu()
}

func multiPlayer() {
	fmt.Println("multi")
	var input string
	for {
		fmt.Println("Pick your sign (x or o)")
		fmt.Scanln(&input)
		if input == "x" {
			p1.PlayerSign = "x"
			p2.PlayerSign = "o"
			break
		} else if input == "o" {
			p1.PlayerSign = "o"
			p2.PlayerSign = "x"
			break
		} else {
			fmt.Println("Wrong sign, try again")
		}
	}
	fmt.Print("\n********************************************\n")
	fmt.Println("use cordinate to play the game")
	fmt.Println("row=1, col=1 | row=1, col=2 | row=1, col=3")
	fmt.Println("------------------------------------------")
	fmt.Println("row=2, col=1 | row=2, col=2 | row=2, col=3")
	fmt.Println("------------------------------------------")
	fmt.Println("row=3, col=1 | row=3, col=2 | row=3, col=3")
	fmt.Print("\n********************************************\n")

	c := 0
	for {
		var rS bool
		if c%2 == 0 {
			rS = gf.move(p1)
		} else {
			rS = gf.move(p2)
		}
		if !rS {
			continue
		}
		renderBoard(gf)
		// Detect final move and and play loop
		if gf.finalMoveCheck() {
			break
		}
		c++

	}

}

func singlePlayer() {
	fmt.Println("single")

}

func initGameField() gameField {
	var gf gameField
	for i := 0; i < len(gf); i++ {
		for j := 0; j < len(gf[0]); j++ {
			gf[i][j] = " "
		}
	}
	return gf
}

func renderBoard(gf gameField) {
	for i := 0; i < len(gf); i++ {
		for j := 0; j < len(gf[0]); j++ {
			fmt.Print(gf[i][j])
			if j < 2 {
				fmt.Print("|")
			}
		}
		fmt.Print("\n")
		if i < 2 {
			fmt.Println("------")
		}
	}
}

func (gf *gameField) move(p playerSign) bool {
	var input string
	var row, col int
	var err error
	fmt.Println(p.PlayerName, " move")
	for {
		fmt.Println("row:")
		fmt.Scanln(&input)
		row, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Incorrect row, please type 1 or 2 or 3")
		} else {
			break
		}
	}
	for {
		fmt.Println("column:")
		fmt.Scanln(&input)
		col, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Incorrect column, please type 1 or 2 or 3")
		} else {
			break
		}
	}
	if gf[row-1][col-1] == " " {
		gf[row-1][col-1] = p.PlayerSign
		return true
	}
	fmt.Println("Field already set, try again")
	return false
}

func (gf *gameField) finalMoveCheck() bool {
	// add draft???
	// check row
	for i := 0; i < len(gf); i++ {
		if gf[i][0] == gf[i][1] && gf[i][1] == gf[i][2] {
			if gf[i][0] == "x" {
				fmt.Println("x WON!")
				return true
			} else if gf[i][0] == "o" {
				fmt.Println("o WON!")
				return true
			}
		}
	}
	// check columns
	for i := 0; i < len(gf); i++ {
		if gf[0][i] == gf[1][i] && gf[1][i] == gf[2][i] {
			if gf[0][i] == "x" {
				fmt.Println("x WON!")
				return true
			} else if gf[0][i] == "o" {
				fmt.Println("o WON!")
				return true
			}
		}
	}
	// diagonal
	if gf[0][0] == gf[1][1] && gf[1][1] == gf[2][2] {
		if gf[0][0] == "x" {
			fmt.Println("x WON!")
			return true
		} else if gf[0][0] == "o" {
			fmt.Println("o WON!")
			return true
		}
	}
	if gf[0][2] == gf[1][1] && gf[1][1] == gf[2][0] {
		if gf[0][2] == "x" {
			fmt.Println("x WON!")
			return true
		} else if gf[0][2] == "o" {
			fmt.Println("o WON!")
			return true
		}
	}
	return false
}
