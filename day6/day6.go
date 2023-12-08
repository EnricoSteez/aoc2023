package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// f, err := os.Open("test.txt")
	f, err := os.Open("input.txt")
	if err != nil {
		panic("cannot open input file")
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	scanner.Scan()
	timesStr := strings.Fields(strings.Split(scanner.Text(), ":")[1])
	times := make([]int, len(timesStr))
	for i, t := range timesStr {
		times[i], _ = strconv.Atoi(t)
	}

	scanner.Scan()
	distancesStr := strings.Fields(strings.Split(scanner.Text(), ":")[1])
	distances := make([]int, len(timesStr))
	for i, d := range distancesStr {
		distances[i], _ = strconv.Atoi(d)
	}

	var score int
	var waysThisGame int
	var startedWinning bool
	res := 1
	for i, maxTime := range times {
		waysThisGame = 0
		startedWinning = false
		for time := 0; time < maxTime; time++ {
			score = (maxTime - time) * time
			// if i == 0 {
			// 	fmt.Printf("hold: %d, score: %d\n", time, score)
			// }
			// score = time * (times[i] - time + 1)
			// Score starts decreasing
			if score <= distances[i] && startedWinning {
				break
			}
			if score > distances[i] {
				startedWinning = true
				waysThisGame++
			}
		}
		fmt.Printf("Ways to win game %d: %d\n", i+1, waysThisGame)
		res *= waysThisGame
	}

	fmt.Println(res)
}
