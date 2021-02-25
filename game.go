package main

import (
	"fmt"
	"os"
	"strconv"
)

type gameField [3][3]string

type playerSign struct {
	PlayerName string
	Sign       string
	IsComputer bool
}
type cord struct {
	X int
	Y int
}

var gf = initGameField()
var p1 = playerSign{PlayerName: "Player1"}
var p2 = playerSign{PlayerName: "Player2"}
var winStrike []cord

const (
	cReset  = "\033[0m"
	cYellow = "\033[33m"
	cGreen  = "\033[32m"
	cPurple = "\033[35m"
)

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
	var input string
	for {
		fmt.Println("Pick your sign (x or o)")
		fmt.Scanln(&input)
		if input == "x" {
			p1.Sign = "x"
			p2.Sign = "o"
			break
		} else if input == "o" {
			p1.Sign = "o"
			p2.Sign = "x"
			break
		}
		fmt.Println("Wrong sign, try again")
	}

	CallClear()
	fmt.Print("\n********************************************\n")
	fmt.Println("use cordinate to play the game")
	fmt.Println("row=1, col=1 | row=1, col=2 | row=1, col=3")
	fmt.Println("------------------------------------------")
	fmt.Println("row=2, col=1 | row=2, col=2 | row=2, col=3")
	fmt.Println("------------------------------------------")
	fmt.Println("row=3, col=1 | row=3, col=2 | row=3, col=3")
	fmt.Print("********************************************\n\n")

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
		CallClear()
		renderBoard(gf)
		// Detect final move and stop play loop
		if finalMoveCheck(&gf) {
			// TODO:renderBoardFinal(gf, false)
			break
		}
		c++
	}

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

// add render color for win strike or draft
// func renderBoardFinal(gf gameField, isDraft bool) {
// 	for i := 0; i < len(gf); i++ {
// 		for j := 0; j < len(gf[0]); j++ {
// 			if isDraft {
// 				fmt.Print(gf[i][j])
// 			} else {
// 				fmt.Println(string(cGreen), gf[i][j], string(cReset))
// 			}
// 			if j < 2 {
// 				fmt.Print("|")
// 			}
// 		}
// 		fmt.Print("\n")
// 		if i < 2 {
// 			fmt.Println("------")
// 		}
// 	}
// }

func (gf *gameField) move(p playerSign) bool {
	var input string
	var row, col int
	var err error
	fmt.Println(p.PlayerName, string(cPurple), p.Sign, string(cReset), " move")
	for {
		fmt.Println("row:")
		fmt.Scanln(&input)
		row, err = strconv.Atoi(input)
		if err != nil || row < 1 || row > 3 {
			fmt.Println("Incorrect row, please type 1 or 2 or 3.")
		} else {
			break
		}
	}
	for {
		fmt.Println("column:")
		fmt.Scanln(&input)
		col, err = strconv.Atoi(input)
		if err != nil || col < 1 || col > 3 {
			fmt.Println("Incorrect column, please type 1 or 2 or 3.")
		} else {
			break
		}
	}
	if gf[row-1][col-1] == " " {
		gf[row-1][col-1] = p.Sign
		return true
	}
	fmt.Println("Field already set, try again.")
	return false
}

// finalMoveCheck check is someone won the game
func finalMoveCheck(gf *gameField) bool {
	// check row
	for i := 0; i < len(gf); i++ {
		if gf[i][0] == gf[i][1] && gf[i][1] == gf[i][2] {
			winStrike = append(winStrike,
				cord{X: i, Y: 0},
				cord{X: i, Y: 1},
				cord{X: i, Y: 2})
			if gf[i][0] == "x" {
				fmt.Println(string(cYellow), "x WON!", string(cReset))
				return true
			} else if gf[i][0] == "o" {
				fmt.Println(string(cYellow), "o WON!", string(cReset))
				return true
			}
		}
	}
	// check columns
	for i := 0; i < len(gf); i++ {
		if gf[0][i] == gf[1][i] && gf[1][i] == gf[2][i] {
			winStrike = append(winStrike,
				cord{X: 0, Y: i},
				cord{X: 1, Y: i},
				cord{X: 2, Y: i})
			if gf[0][i] == "x" {
				fmt.Println(string(cYellow), "x WON!", string(cReset))
				return true
			} else if gf[0][i] == "o" {
				fmt.Println(string(cYellow), "o WON!", string(cReset))
				return true
			}
		}
	}
	// check diagonal
	if gf[0][0] == gf[1][1] && gf[1][1] == gf[2][2] {
		winStrike = append(winStrike,
			cord{X: 0, Y: 0},
			cord{X: 1, Y: 1},
			cord{X: 2, Y: 2})
		if gf[0][0] == "x" {
			fmt.Println(string(cYellow), "x WON!", string(cReset))
			return true
		} else if gf[0][0] == "o" {
			fmt.Println(string(cYellow), "o WON!", string(cReset))
			return true
		}
	}

	// check diagonal
	if gf[0][2] == gf[1][1] && gf[1][1] == gf[2][0] {
		winStrike = append(winStrike,
			cord{X: 0, Y: 2},
			cord{X: 1, Y: 1},
			cord{X: 2, Y: 0})
		if gf[0][2] == "x" {
			fmt.Println(string(cYellow), "x WON!", string(cReset))
			return true
		} else if gf[0][2] == "o" {
			fmt.Println(string(cYellow), "o WON!", string(cReset))
			return true
		}
	}

	if !moveIsAvailable(gf) {
		fmt.Println("DRAW!")
		return true
	}

	return false
}

// moveIsAvailable check we have empty field on the board
// return true if exists empty filed else false
func moveIsAvailable(gf *gameField) bool {
	for i := 0; i < len(gf); i++ {
		for j := 0; j < len(gf[0]); j++ {
			if gf[i][j] == " " {
				return true
			}
		}
	}
	return false
}
