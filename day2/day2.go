package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	sum := 0
	var minCubesPerColor map[string]int

	// Games LOOP
	for scanner.Scan() {
		line = scanner.Text()
		fmt.Println(line)
		lr := strings.Split(line, ":")
		right := strings.Trim(lr[1], " ")
		// Draws Loop
		minCubesPerColor = make(map[string]int)
		for _, draw := range strings.Split(right, ";") {
			// Handful of cubes per color drawn in one single draw
			for _, amountColor := range strings.Split(strings.Trim(draw, " "), ",") {
				ac := strings.Split(strings.Trim(amountColor, " "), " ")
				count, _ := strconv.Atoi(ac[0])
				color := ac[1]
				if cur, in := minCubesPerColor[color]; !in || count < cur {
					minCubesPerColor[color] = count
				}
			}
		}
		prod := 1
		for color, count := range minCubesPerColor {
			prod *= count
			fmt.Printf("%s: %d, ", color, count)
		}
		sum += prod
		fmt.Printf("\n PROD IS %d, new sum is %d\n", prod, sum)
	}

	fmt.Println("Total: ", sum)
}
