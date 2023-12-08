package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"testing"
)

type mapping struct {
	sourceStart      int
	destinationStart int
	length           int
}

type mappingBlock struct {
	name     string
	mappings []mapping
}

type seedChunk struct {
	start int
	end   int
}

func (mb mappingBlock) Print() {
	fmt.Println(mb.name)
	for _, m := range mb.mappings {
		fmt.Printf("%d %d %d\n", m.sourceStart, m.destinationStart, m.length)
	}
	fmt.Println()
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic("cannot open input file")
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var line string
	var split []string
	scanner.Scan()
	seedsStr := strings.Fields(strings.Split(scanner.Text(), ":")[1])
	seeds := make([]seedChunk, len(seedsStr))
	for i := 0; i < len(seeds); i++ {
		start, _ := strconv.Atoi(seedsStr[i])
		i++
		len, _ := strconv.Atoi(seedsStr[i])
		seeds[i] = seedChunk{start: start, end: start + len - 1}
	}

	mainStruct := make([]mappingBlock, 7)
	currentIdx := 0

	scanner.Scan()       // Go to blank line
	for scanner.Scan() { // Go to title
		if currentIdx > 0 {
			// Sort the previously populated mapping
			sort.Slice(mainStruct[currentIdx-1].mappings, func(i, j int) bool {
				if mainStruct[currentIdx-1].mappings[i].sourceStart < mainStruct[currentIdx-1].mappings[j].sourceStart {
					return true
				} else if mainStruct[currentIdx-1].mappings[i].sourceStart > mainStruct[currentIdx-1].mappings[j].sourceStart {
					return false
				} else {
					if mainStruct[currentIdx-1].mappings[i].destinationStart < mainStruct[currentIdx-1].mappings[j].destinationStart {
						return true
					} else if mainStruct[currentIdx-1].mappings[i].destinationStart > mainStruct[currentIdx-1].mappings[j].destinationStart {
						return false
					} else {
						return mainStruct[currentIdx-1].mappings[i].length < mainStruct[currentIdx-1].mappings[j].length
					}
				}
			})
		}
		if currentIdx == len(mainStruct) {
			break
		}
		name := scanner.Text()
		mainStruct[currentIdx] = mappingBlock{name: name}
		mainStruct[currentIdx].mappings = make([]mapping, 0)
		scanner.Scan() // First number row of block
		line = scanner.Text()
		for line != "" {
			split = strings.Fields(line)
			dest, _ := strconv.Atoi(split[0])
			source, _ := strconv.Atoi(split[1])
			len, _ := strconv.Atoi(split[2])
			mainStruct[currentIdx].mappings = append(mainStruct[currentIdx].mappings, mapping{sourceStart: source, destinationStart: dest, length: len})
			scanner.Scan()
			line = scanner.Text() // Eventually go to blank
		}
		currentIdx++
	}

	// for _, m := range mainStruct {
	// 	m.Print()
	// }

	location := 0
	for _, sc := range seeds {

		// Use seed as final number in the mappings chain
		if location == 0 || sc < location {
			location = sc
		}
	}

	fmt.Println("min location: ", location)

}

func TestMain(t *testing.T) {
	main()
}
