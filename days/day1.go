package days

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func Day1() {
	fmt.Println("Hello from Advent of Code Day 1")
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	totals_slice := []int{}
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			totals_slice = append(totals_slice, sum)
			sum = 0
			continue
		}
		num, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("Error converting string to int")
		} else {
			sum += num
		}

	}

	fmt.Println("\n", totals_slice)

	max := 0.0
	var curr float64
	for _, val := range totals_slice {
		curr = float64(val)
		max = math.Max(max, curr)
	}

	fmt.Printf("\nThe elf with most calories has %d calories\n", int(math.Floor(max)))

	top3 := [3]int{0, 0, 0}
	for _, val := range totals_slice {
		if val > top3[0] {
			top3[2] = top3[1]
			top3[1] = top3[0]
			top3[0] = val
		} else if val > top3[1] {
			top3[2] = top3[1]
			top3[1] = val
		} else if val > top3[2] {
			top3[2] = val
		}
	}

	fmt.Printf("\nThe top 3 elves have %d, %d, and %d calories\n", top3[0], top3[1], top3[2])
	top3_total := top3[0] + top3[1] + top3[2]
	fmt.Printf("\nThe top 3 elves have %d calories in total\n", top3_total)

}
