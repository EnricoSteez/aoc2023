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
	var timeStr strings.Builder
	for _, s := range timesStr {
		timeStr.WriteString(s)
	}
	maxTime, _ := strconv.Atoi(timeStr.String())

	scanner.Scan()
	distancesStr := strings.Fields(strings.Split(scanner.Text(), ":")[1])
	var distanceStr strings.Builder
	for _, s := range distancesStr {
		distanceStr.WriteString(s)
	}
	distance, _ := strconv.Atoi(distanceStr.String())

	fmt.Println(maxTime, distance)
	var score int
	var firstWin int
	for i := 0; i < maxTime; i++ {
		score = (maxTime - i) * i
		if score > distance {
			firstWin = i
			break
		}
	}
	var lastWin int
	score = 0
	for i := maxTime; score <= distance; i-- {
		score = (maxTime - i) * i
		if score > distance {
			lastWin = i
			break
		}
	}

	fmt.Println(lastWin - firstWin + 1)
}
