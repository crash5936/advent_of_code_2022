package days

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Day2() {
	fmt.Println("Day 2\n---------------------------------")
	scanner := bufio.NewScanner(os.Stdin)
	var opp_score int = 0
	var my_score int = 0
	for scanner.Scan() {
		line := scanner.Text()
		play := strings.Split(line, " ")
		switch opp := play[0]; opp {
			case "A": // opp plays rock
				switch mine := play[1]; mine {
				case "X": // I play rock
					my_score += 1 + 3 // draw
					opp_score += 1 + 3 // draw
				case "Y": // I play paper
					my_score += 2 + 6 // win
					opp_score += 1 + 0 // lose
				case "Z": // I play scissors
					my_score += 3 + 0 // lose
					opp_score += 1 + 6 // win
				}
			case "B": // opp plays paper
				switch mine := play[1]; mine {
				case "X": // I play rock
					my_score += 1 + 0 // lose
					opp_score += 1 + 6 // win
				case "Y": // I play paper
					my_score += 2 + 3 // draw
					opp_score += 2 + 3 // draw
				case "Z": // I play scissors
					my_score += 3 + 6 // win
					opp_score += 2 + 0 // lose
				}
			case "C": // opp plays scissors
				switch mine := play[1]; mine {
				case "X": // I play rock
					my_score += 1 + 6 // win
					opp_score += 3 + 0 // lose
				case "Y": // I play paper
					my_score += 2 + 0 // lose
					opp_score += 3 + 6 // win
				case "Z": // I play scissors
					my_score += 3 + 3 // draw
					opp_score += 3 + 3 // draw
			}
		}
	}
	fmt.Printf("My score: %d\n", my_score)
}

func Day2Part2() {
	fmt.Println("Day 2 Part 2\n---------------------------------")
	scanner := bufio.NewScanner(os.Stdin)
	var opp_score int = 0
	var my_score int = 0
	for scanner.Scan() {
		line := scanner.Text()
		play := strings.Split(line, " ")
		switch opp := play[0]; opp {
			case "A": // opp plays rock
				switch mine := play[1]; mine {
				case "X": // I need to lose - I play scissors
					my_score += 3 + 0
					opp_score += 1 + 6
				case "Y": // I need to draw - I play rock
					my_score += 1 + 3
					opp_score += 1 + 3
				case "Z": // I need to win - I play paper
					my_score += 2 + 6
					opp_score += 1 + 0
				}
			case "B": // opp plays paper
				switch mine := play[1]; mine {
					case "X": // I need to lose - I play rock
					my_score += 1 + 0
					opp_score += 2 + 6
				case "Y": // I need to draw - I play paper
					my_score += 2 + 3
					opp_score += 2 + 3
				case "Z": // I need to win - I play scissors
					my_score += 3 + 6
					opp_score += 2 + 0
				}
			case "C": // opp plays scissors
				switch mine := play[1]; mine {
				case "X": // I need to lose - I play paper
					my_score += 2 + 0
					opp_score += 3 + 6
				case "Y": // I need to draw - I play scissors
					my_score += 3 + 3
					opp_score += 3 + 3
				case "Z": // I need to win - I play rock
					my_score += 1 + 6
					opp_score += 3 + 0
			}
		}
	}
	fmt.Printf("My score: %d\n", my_score)
}