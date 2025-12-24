package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func LoadTestData(theFileName string) (string, error) {
	_, file, _, ok := runtime.Caller(0)

	if !ok {
		return "", fmt.Errorf("Failed to open test data file.")
	}

	testDataFilePath := filepath.Join(filepath.Dir(file), "test_data", theFileName)

	fileData, err := os.ReadFile(testDataFilePath)

	if err != nil {
		return "", fmt.Errorf("Failed to Read Data File with error: %s", err.Error())
	}

	return string(fileData), nil
}

func CountFalse(theBools []bool) int {
	result := 0
	for _, value := range theBools {
		if !value {
			result++
		}
	}
	return result
}

func MakeTrueSlice(theSize int) []bool {
	result := make([]bool, theSize)
	for index := range theSize {
		result[index] = true
	}
	return result
}

func NumToBoolSlice(n int, theSliceLenght int) []bool {
	// Pre-allocate a slice of exactly 15 booleans
	res := make([]bool, theSliceLenght)

	for i := range theSliceLenght {
		// We want the 15th bit first (index 14), then 13, etc.
		// If (n & mask) is 0, it means the bit is 0 -> true
		res[i] = (n & (1 << (theSliceLenght - 1 - i))) == 0
	}
	return res
}

func main() {
	dataStr, err := LoadTestData("test_data.txt")

	if err != nil {
		panic("Failed to load test data: " + err.Error())
	}

	data := strings.Split(dataStr, "\n")

	for _, line := range data {

		for index, start := range line {
			numStart :=
				fmt.Print(string(line[index]), " ")

			tail := line[index:]

			if len(tail) < 12 {
				break
			}

			for i := range 1 << len(tail) {
				digitsToUse := NumToBoolSlice(i, len(tail))

				if CountFalse(digitsToUse) == 3 {
					newDigit := start
					// fmt.Print(start, " ")

					for index, value := range digitsToUse {
						if value {
							fmt.Print(tail[index], " ")
							temp := tail[index]
							newDigit += rune(temp)
						}
					}
					fmt.Println(newDigit)
				}
			}

		}

	}
}
