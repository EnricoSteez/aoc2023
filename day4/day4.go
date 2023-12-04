package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic("cannot open input file")
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var line string
	var totalScore int

	for scanner.Scan() {
		line = scanner.Text()
		// fmt.Println(line)
		lr := strings.Split(line, ":")
		right := lr[1]
		numbers := strings.Split(right, "|")
		winning := strings.Split(numbers[0], " ")
		played := strings.Split(numbers[1], " ")
		winningNumbers := make(map[string]struct{})

		cardScore := 0
		for _, nStr := range winning {
			nStr = strings.Trim(nStr, " ")
			if nStr == "" {
				continue
			}
			winningNumbers[nStr] = struct{}{}
			fmt.Printf("%s ", nStr)
		}
		fmt.Println()
		for _, nStr := range played {
			nStr = strings.Trim(nStr, " ")
			if nStr == "" {
				continue
			}
			if _, in := winningNumbers[nStr]; in {
				fmt.Printf("%s, ", nStr)
				cardScore++
			}
		}

		var thisScore int
		if cardScore > 0 {
			thisScore = int(math.Pow(2, float64(cardScore)-1))
		}
		totalScore += thisScore
		fmt.Printf("winning numbers: score %d\n", thisScore)
	}

	fmt.Println("total score ", totalScore)

}
