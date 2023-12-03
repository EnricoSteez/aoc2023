package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// type trie struct {
// 	children map[byte]*trie
// 	isLeaf   bool
// }
//
// func (t *trie) contains(input string) bool {
// 	for i := 0; i < len(input); i++ {
// 		if _, contained := t.children[input[i]]; !contained {
// 			return false
// 		}
// 		t = t.children[input[i]]
// 	}
// 	return t.isLeaf
// }
//
// func (t *trie) add(input string) {
// 	for i := 0; i < len(input); i++ {
// 		ch := input[i]
// 		if t.children == nil {
// 			t.children = make(map[byte]*trie)
// 		}
// 		if _, contains := t.children[ch]; !contains {
// 			t.children[ch] = new(trie)
// 		}
// 		t = t.children[ch]
// 	}
// 	t.isLeaf = true
// }

func isDigit(c byte) bool {
	return c >= 48 && c <= 57
}

func byteToInt(c byte) int {
	return int(c) - 48
}

func main() {
	// Populate trie with spelled digits
	spelledDigits := map[string]int{
		"zero": 0, "one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9,
	}

	f, err := os.Open("day1_input.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var line string
	sum := 0

	for scanner.Scan() {
		line = scanner.Text()
		// line = "12sldfjsledjrfslkdflwsd"
		// line = "sonesevenlcmfefiodchh"
		// line = "two32ndxvvqpnn5"
		// line = "2gzqrfldtlpeight3fivencmlmffivevqkhncfm"
		// line = "dhpbgtkmjfourone6rsgnpvsbjtkfqsvrs9threethree"
		// fmt.Println(line)
		var first int
		var last int
		firstPos := -1
		lastPos := -1

		// First look for spelled digits
		for i := 0; i < len(line); i++ {
			for j := i; j < len(line) && math.Abs(float64(j-i)) <= 5; j++ {
				substringFirst := line[i : j+1]
				substringLast := line[len(line)-j-1 : len(line)-i]
				if val, in := spelledDigits[substringFirst]; in && firstPos == -1 {
					// We haven't found the first digit yet, so we must set the first
					first = val
					firstPos = i
					// fmt.Printf("found first spelled %s at %d\n", substringFirst, firstPos)
				}
				// fmt.Printf("Substring last %s\n", substringLast)
				if val, in := spelledDigits[substringLast]; in && lastPos == -1 {
					// We haven't found the second digit yet, so we must set the second
					// We haven't found the first digit yet, so we must set the first
					last = val
					lastPos = len(line) - j - 1
					// fmt.Printf("found last spelled %s at %d\n", substringLast, lastPos)
				}
				if firstPos >= 0 && lastPos >= 0 {
					break
				}
			}
			if firstPos >= 0 && lastPos >= 0 {
				break
			}
		}
		for i := 0; i < len(line); i++ {
			j := len(line) - i - 1
			// fmt.Println("Idx last is ", j)
			if (firstPos == -1 || i < firstPos) && isDigit(line[i]) {
				firstPos = i
				first = byteToInt(line[i])
				// fmt.Printf("foundFirstDigit %d at %d\n", first, i)
			}
			if (lastPos == -1 || j > lastPos) && isDigit(line[j]) {
				lastPos = j
				last = byteToInt(line[j])
				// fmt.Printf("foundLastDigit %d at %d\n", last, j)
			}
			if firstPos >= 0 && i > firstPos && lastPos >= 0 && j < lastPos {
				break
			}
		}
		n := first*10 + last
		// fmt.Printf("Line is %s, first is %d, last is %d, total is %d\n", line, first, last, n)
		sum = sum + n
	}

	fmt.Printf("Final value is %d\n", sum)
}
