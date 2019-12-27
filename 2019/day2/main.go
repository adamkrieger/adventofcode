package main

import (
	"bufio"
	"fmt"
	"github.com/go-errors/errors"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strconv"
)


func main() {
	fileHandle := openOrPanic("./input.txt")
	fileReader := bufio.NewReader(fileHandle)

	valueStream := make(chan int)
	workDone := manageArr(valueStream)

	var currentValue []byte
	fullyRead := false
	for fullyRead == false {
		nextByte, readErr := fileReader.ReadByte()

		if readErr == nil {

			if nextByte == ',' {
				currentValue = convertAndDump(valueStream,currentValue)
			} else {
				currentValue = append(currentValue, nextByte)
			}
		} else if readErr == io.EOF {
			_ = convertAndDump(valueStream, currentValue)
			fullyRead = true
			close(valueStream)
		} else {
			panic(readErr)
		}
	}

	logrus.Info("Done Reading, starting to wait")

	_ = <-workDone
	_ = fileHandle.Close()
}

func convertAndDump(valueStream chan<- int, currentValue []byte) (newArr []byte){
	if len(currentValue) <= 0 {
		return currentValue
	}

	val, err := strconv.ParseInt(string(currentValue),10, 64)
	if err != nil {
		panic(err)
	}

	valueStream <- int(val)

	return []byte{}
}

func openOrPanic(fileLocation string) (*os.File) {
	var openErr error
	inputFile, openErr := os.Open(fileLocation)
	if openErr!=nil {
		panic(openErr)
	}
	return inputFile
}

func manageArr(valueStream <-chan int) <-chan bool {
	workDone := make(chan bool)

	go func() {
		var valueArr []int

		opCodeI := 0

		readStart := false
		readEnd := false
		procEnd := false

		for !procEnd || !readEnd {

			select {
			case nextValue, ok := <-valueStream:
				readStart = true
				if !ok {
					readEnd = true
					valueStream = nil
				} else {
					valueArr = append(valueArr, nextValue)
				}
			default:
				if readStart {
					var res int
					if valueArr, opCodeI, res = processMut(valueArr, opCodeI); res == exitNow {
						procEnd = true
					} else if res == waiting {
						if readEnd {
							procEnd = true
						} else {
							logrus.Info("Waiting for more text stream.")
						}
					} else if res == unrecognizableOpCode {
						panic(errors.New(fmt.Sprintf("Unrecognizable opCode at idx: %d", opCodeI)))
					}
				}
			}
		}

		logrus.Info("Closing and exiting. ", valueArr)

		close(workDone)
		return
	}()

	return workDone
}

const(
	nextOpCodeI    = 4
	opCodeRelIArg1 = 1
	opCodeRelIArg2 = 2
	opCodeRelIDest = 3

	unrecognizableOpCode = 3
	waiting = 2
	exitNow = 1
	continueProc = 0
)

func processMut(arr []int, nextI int) (retArr []int, nextOpCode, exitNowResult int) {
	if nextI < 0 || nextI >= len(arr) {
		return arr, nextI, waiting
	}

	retI := nextI
	result := waiting

	switch arr[nextI] {
	case 99:
		return arr, -1, exitNow
	case 1:
		nextIPlus := nextI + nextOpCodeI

		if (nextIPlus) > len(arr) {
			retI = nextI
			result = waiting
		} else {
			logrus.Info("Adding: ", nextI, len(arr))

			argi1 := arr[nextI + opCodeRelIArg1]
			argi2 := arr[nextI + opCodeRelIArg2]
			desti := arr[nextI + opCodeRelIDest]

			if desti >= len(arr) || argi1 >= len(arr) || argi2 >= len(arr) {
				retI = nextI
				result = waiting
			} else {
				arr[desti] = arr[argi1] + arr[argi2]

				//Part2
				arr[nextI] += 4

				retI = nextIPlus
				result = continueProc
			}
		}
	case 2:
		nextIPlus := nextI + nextOpCodeI

		if (nextIPlus) > len(arr) {
			retI = nextI
			result = waiting
		} else {
			logrus.Info("Multiplying: ", nextI, len(arr))

			argi1 := arr[nextI + opCodeRelIArg1]
			argi2 := arr[nextI + opCodeRelIArg2]
			desti := arr[nextI + opCodeRelIDest]

			if desti >= len(arr) || argi1 >= len(arr) || argi2 >= len(arr) {
				retI = nextI
				result = waiting
			} else {
				arr[desti] = arr[argi1] * arr[argi2]

				//Part2
				arr[nextI] += 4

				retI = nextIPlus
				result = continueProc
			}
		}
	default:
		retI = -1
		result = unrecognizableOpCode
	}

	return arr, retI, result
}
