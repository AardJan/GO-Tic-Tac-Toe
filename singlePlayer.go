package main

import (
	"fmt"
	"os"
)

// singlePlayer game run
func singlePlayer() {
	fmt.Println("single")
	var input string
	for {
		fmt.Println("Pick your sign (x or o)")
		fmt.Scanln(&input)
		if input == "x" {
			p1.Sign = "x"
			p2.Sign = "o"
			p2.IsComputer = true
			break
		} else if input == "o" {
			p1.Sign = "o"
			p2.Sign = "x"
			p2.IsComputer = true
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
			// pc move
			// rS = gf.move(p2)
			fmt.Println("pc move")
			bCord := findBestMove(&gf)
			gf[bCord.X][bCord.Y] = p2.Sign
			rS = true
			fmt.Println("pc move end")
		}
		if !rS {
			continue
		}
		// CallClear()
		renderBoard(gf)
		// Detect final move and stop play loop
		if finalMoveCheck(&gf) {
			// renderBoardFinal(gf, false)
			fmt.Println("Final move!")
			break
		}
		c++
	}
}

// check is final move, return points
func evaluate(gf *gameField) int {
	// check row
	for i := 0; i < len(gf); i++ {
		if gf[i][0] == gf[i][1] && gf[i][1] == gf[i][2] {
			if gf[i][0] == "x" {
				if p2.Sign == "x" {
					return 10
				}
				return -10
			} else if gf[i][0] == "o" {
				if p2.Sign == "o" {
					return 10
				}
				return -10
			}
		}
	}
	// check columns
	for i := 0; i < len(gf); i++ {
		if gf[0][i] == gf[1][i] && gf[1][i] == gf[2][i] {
			if gf[0][i] == "x" {
				if p2.Sign == "x" {
					return 10
				}
				return -10
			} else if gf[0][i] == "o" {
				if p2.Sign == "o" {
					return 10
				}
				return -10
			}
		}
	}
	// check diagonal
	if gf[0][0] == gf[1][1] && gf[1][1] == gf[2][2] {
		if gf[0][0] == "x" {
			if p2.Sign == "x" {
				return 10
			}
			return -10
		} else if gf[0][0] == "o" {
			if p2.Sign == "o" {
				return 10
			}
			return -10
		}
	}

	// check diagonal
	if gf[0][2] == gf[1][1] && gf[1][1] == gf[2][0] {
		if gf[0][2] == "x" {
			if p2.Sign == "x" {
				return 10
			}
			return -10
		} else if gf[0][2] == "o" {
			if p2.Sign == "o" {
				return 10
			}
			return -10
		}
	}
	return 0
}

func minmax(gf *gameField, depth int, isMax bool) int {
	score := evaluate(gf)

	f, err := os.OpenFile("./test.log", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error open file")
		os.Exit(-1)
	}
	fmt.Fprintf(f, fmt.Sprintln("Score ", score, " Depth: ", depth, " isMax: ", isMax))
	for i := 0; i < len(gf); i++ {
		for j := 0; j < len(gf[0]); j++ {
			fmt.Fprintf(f, fmt.Sprint(gf[i][j]))
			if j < 2 {
				fmt.Fprintf(f, "|")
			}
		}
		fmt.Fprintf(f, "\n")
		if i < 2 {
			fmt.Fprintf(f, "------\n")
		}
	}
	// fmt.Println("Score ", score, " Depth: ", depth)
	fmt.Fprintf(f, "\n")
	f.Close()

	if score == 10 {
		return score
	}
	if score == -10 {
		return score
	}
	if !moveIsAvailable(gf) {
		return 0
	}

	if isMax {
		best := -1000
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if gf[i][j] == " " {
					gf[i][j] = p2.Sign
					best = intMax(best, minmax(gf, depth+1, !isMax))
					gf[i][j] = " "
				}
			}
		}
		return best
	}

	// if !isMax
	best := 1000
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if gf[i][j] == " " {
				gf[i][j] = p1.Sign
				best = intMin(best, minmax(gf, depth+1, !isMax))
				gf[i][j] = " "
			}
		}
	}
	return best
}

func findBestMove(gf *gameField) cord {
	bestVal := -100
	bestMove := cord{X: -1, Y: -1}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if gf[i][j] == " " {

				gf[i][j] = p2.Sign

				moveVal := minmax(gf, 0, false)

				gf[i][j] = " "

				if moveVal >= bestVal {
					bestMove.X = i
					bestMove.Y = j
					bestVal = moveVal
				}

				gf[i][j] = " "
			}
		}
	}
	fmt.Println("bestVal: ", bestVal)
	fmt.Println("bestMove: ", bestMove)
	return bestMove
}

func intMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func intMin(x, y int) int {
	if x < y {
		return x
	}
	return y
}
