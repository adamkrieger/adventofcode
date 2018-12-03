package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	path := os.Args[1]

	fileHndl, err := os.Open(path)

	if err != nil {
		os.Exit(1)
	}

	rdr := bufio.NewReader(fileHndl)

	currentLine := []byte{}
	total := 0

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
			}
		}

	}

	log.Println(total)
}
