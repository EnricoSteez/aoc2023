package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isDigit(t byte) bool {
	// 48 == "0"
	// 57 == "9"
	return t >= 48 && t <= 57
}
func isSymbol(t byte) bool {
	// 46 == "."
	// 13 == "\n"  --> not sure if this is needed tho
	return !isDigit(t) && t != 46 && t != 13
}

func isValidPart(prevLine, line, nextLine string, numStartIdx, numEndIdxPlusOne int) bool {
	// Start is one before the start of the part number,
	// or 0 if the part sits at the beginning of the line
	var start int
	if numStartIdx == 0 {
		start = 0
	} else {
		start = numStartIdx - 1
	}
	// End is one after the end of the part number,
	// or the last index of the line if the part ends at the end of the line
	var end int
	if numEndIdxPlusOne == len(line) {
		end = len(line) - 1
	} else {
		end = numEndIdxPlusOne
	}

	// Check one line above the number for a symbol that validates this part
	// prevLine might be empty if line is the first line in the file
	if prevLine != "" {
		for i := start; i <= end; i++ {
			if isSymbol(prevLine[i]) {
				return true
			}
		}
	}

	// Do the same for the line below
	// nextLine might be empty if line is the last line in the file
	if nextLine != "" {
		for i := 0; i <= end; i++ {
			if isSymbol(nextLine[i]) {
				return true
			}
		}
	}

	// Check right after the number on the same line
	if isSymbol(line[end]) {
		return true
	}
	// Check right before the number on the same line
	if isSymbol(line[start]) {
		return true
	}

	// If none of the previous checks were successful, this part is not valid
	return false
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic("cannot open input file")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	sum := 0
	var start int
	scanner.Scan()
	line := scanner.Text()
	var prevLine string
	var nextLine string
	var lineIdx int

	// Games LOOP
	for scanner.Scan() {
		lineIdx++
		nextLine = scanner.Text()
		//
		// prevLine = `...........772............&....345..93...465*................419......676...............-.@521.....-...........................399.662......`
		// line = `.................17+..2..531.......................79........*...589......198*734....534.........614..................109...................`
		// nextLine = `.......301............=............................&..321..895..*..........................344.................694............717...511*....`
		for idx := 0; idx < len(line); idx++ {
			if isDigit(line[idx]) {
				start = idx
				for idx < len(line) && isDigit(line[idx]) {
					idx++
				}

				if isValidPart(prevLine, line, nextLine, start, idx) {
					// fmt.Printf("startIdx: %d, endIdx: %d. ", start, idx-1)
					num, err := strconv.Atoi(line[start:idx])
					if err != nil {
						fmt.Println(lineIdx)
						fmt.Printf("StartIdx: %d, EndIdx: %d\n", start, idx)
						fmt.Println("-----")

						fmt.Println(prevLine[start-1 : idx+1])
						fmt.Println(line[start-1 : idx+1])
						fmt.Println(nextLine[start-1 : idx+1])
						fmt.Println("-----")
						panic("Not valid part \n")
					}
					sum += num
				}
			}
		}

		prevLine = line
		line = nextLine
	}

	nextLine = ""
	// Now line is the last line
	for idx := 0; idx < len(line); idx++ {
		if isDigit(line[idx]) {
			start = idx
			for idx < len(line) && isDigit(line[idx]) {
				idx++
			}
			if isValidPart(prevLine, line, nextLine, start, idx) {
				// fmt.Printf("startIdx: %d, endIdx: %d. ", start, idx-1)
				num, err := strconv.Atoi(line[start : idx-1])
				if err != nil {
					panic(fmt.Sprintf("Not valid part %s\n", line[start:idx-1]))
				}
				if start > 0 {
					start -= 1
				}
				if idx < len(line) {
					idx++
				}
				if prevLine != "" {
					fmt.Println(prevLine[start:idx])
				}
				fmt.Println(line[start:idx])
				if nextLine != "" {
					fmt.Println(nextLine[start:idx])
				}
				fmt.Println("-----")
				sum += num
			}
		}
	}

	fmt.Println("Total: ", sum)
}
