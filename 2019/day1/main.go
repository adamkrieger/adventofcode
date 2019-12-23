package main

import (
	"bufio"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"strconv"
)

func main() {
	log.Println("hello")

	var openErr error
	inputFile, openErr := os.Open("./input.txt")
	defer inputFile.Close()

	if openErr != nil {
		log.Error(openErr)
		os.Exit(1)
	}

	fileReader := bufio.NewReader(inputFile)
	var readErr error
	var nextByte byte

	var currentArr []byte

	nextByte, readErr = fileReader.ReadByte()

	var moduleTotal int64

	for readErr == nil {

		if nextByte != '\n' {
			currentArr = append(currentArr, nextByte)
		} else {
			fuel,_ := convertAndCompute(currentArr)
			moduleTotal += fuel
			currentArr = []byte{}
		}

		nextByte, readErr = fileReader.ReadByte()

		if readErr == io.EOF {
			fuel,_ := convertAndCompute(currentArr)
			moduleTotal += fuel
			currentArr = []byte{}
		}

	}

	log.Print("Module Fuel total is: ", moduleTotal)


}

func convertAndCompute(eachInput []byte) (fuel int64, parseErr error) {
	distance, parseErr := strconv.ParseInt(string(eachInput),10, 32)

	if parseErr==nil {
		fuel = computeIncludingFuel(distance)
	}

	log.Printf("in: %s - out: %d - err: %s\n", eachInput, fuel, parseErr)

	return fuel, parseErr
}

func computeIncludingFuel(moduleMass int64) (fuel int64) {
	moduleFuel := compute(moduleMass)

	totalFuel := moduleFuel

	for moduleFuel > 0 {
		next := compute(moduleFuel)

		if next <= 0 {
			moduleFuel = 0
		} else {
			totalFuel += next
			moduleFuel = next
		}
	}

	return totalFuel
}

func compute(mass int64) (fuel int64) {

	fuel = mass / 3 - 2

	return fuel
}

