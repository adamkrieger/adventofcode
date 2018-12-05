package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

var(
	twoMatchTotal = 0
	threeMatchTotal = 0
)

func calcChecksum(twoMatchTotal, threeMatchTotal int) (checksum int) {
	return twoMatchTotal * threeMatchTotal
}

func isMatch(line string) (twoMatch, threeMatch bool) {
	arr := strings.Split(line, "")

	sorted := make(map[string]int)

	for _, each := range arr {
		if _, found := sorted[each]; found {
			sorted[each]++
		} else {
			sorted[each] = 1
		}
	}

	two := false
	three := false

	for _, count := range sorted {
		switch count {
		case 2:
			two = true
		case 3:
			three = true
		}
	}

	return two, three
}

type fileReader struct {
	rdr *bufio.Reader
}

func main() {
	fileRdr := openFile(os.Args[1])

	for {
		line, err := fileRdr.line()

		if err == io.EOF {
			break
		} else if err != nil {
			log.Println(err)
			os.Exit(2)
		}

		twoMatch, threeMatch := isMatch(line)

		if twoMatch {
			twoMatchTotal++
		}

		if threeMatch {
			threeMatchTotal++
		}
	}

	checkSum := calcChecksum(twoMatchTotal, threeMatchTotal)
	log.Println(checkSum)
}

func openFile(path string) *fileReader {
	fileHndl, err := os.Open(path)

	if err != nil {
		os.Exit(1)
	}

	rdr := bufio.NewReader(fileHndl)

	return &fileReader{
		rdr:rdr,
	}
}

func (fileRdr *fileReader) line() (line string, err error) {
	currentLine := []byte{}

	for {
		partial, isPrefix, err := fileRdr.rdr.ReadLine()

		if err != nil {
			return "", err
		} else {
			currentLine = append(currentLine, partial...)
		}

		if isPrefix {
			continue
		} else if len(currentLine) == 0 {
			continue
		} else {
			return string(currentLine), nil
		}

	}
}
