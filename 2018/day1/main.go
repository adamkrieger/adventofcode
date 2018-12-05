package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

var (
	freqMap            = make(map[int]interface{})
	firstDupeFreq      = 0
	firstDupeFreqFound = false
)

func main() {
	path := os.Args[1]

	total := 0

	for {
		total = processFile(total, path)
		if firstDupeFreqFound {
			break
		}
	}

	log.Println("Total: ", total)
	log.Println("FirstDupe: ", firstDupeFreq)
}

func processFile(total int, path string) (newTotal int) {
	fileHndl, err := os.Open(path)

	if err != nil {
		os.Exit(1)
	}

	rdr := bufio.NewReader(fileHndl)

	currentLine := []byte{}

	for {
		partial, isPrefix, err := rdr.ReadLine()

		if err != nil {
			break
		} else {
			currentLine = append(currentLine, partial...)
		}

		if isPrefix {
			continue
		} else if len(currentLine) == 0 {
			continue
		} else {
			currentVal, err := strconv.Atoi(string(currentLine))

			if err != nil {
				log.Println(currentLine)
				os.Exit(3)
			} else {
				currentLine = []byte{}
				total += currentVal

				if _, found := freqMap[total]; !found {
					freqMap[total] = nil
				} else {
					if !firstDupeFreqFound {
						firstDupeFreqFound = true
						firstDupeFreq = total
					}
				}
			}
		}

	}

	fileHndl.Close()
	return total
}
