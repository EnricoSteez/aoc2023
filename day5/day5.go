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
	seeds := make([]int, len(seedsStr))
	for i, s := range seedsStr {
		seed, _ := strconv.Atoi(s)
		seeds[i] = seed
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
	for _, seed := range seeds {
		// fmt.Printf("%d ", seed)
		for _, m := range mainStruct {
			for _, entry := range m.mappings {
				if seed < entry.sourceStart {
					continue
				} else if seed >= entry.sourceStart && seed < entry.sourceStart+entry.length {
					// Since mapping entries are sorted by source start, the first
					// entry with source <= seed is what we are looking for
					offset := seed - entry.sourceStart
					seed = entry.destinationStart + offset
					break
				}
			}
			// fmt.Printf("--> %d ", seed)
		}
		// fmt.Println()
		// Use seed as final number in the mappings chain
		if location == 0 || seed < location {
			location = seed
		}
	}

	fmt.Println("min location: ", location)

}

func TestMain(t *testing.T) {
	main()
}
