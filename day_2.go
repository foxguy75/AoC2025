package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
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

func SplitDigits(theNumberToSplit int, theSplitBy int) []string {
	numStr := strconv.Itoa(theNumberToSplit)
	ret := make([]string, 0, 10)

	if len(numStr)%theSplitBy == 0 {
		for i := 0; i < len(numStr); i += theSplitBy {
			ret = append(ret, numStr[i:i+theSplitBy])
		}
	}

	return ret
}

func SumSlice(theSlice []int) int {
	total := 0

	for _, number := range theSlice {
		total += number
	}

	return total
}

func main() {
	data, err := LoadTestData("test_data.txt")

	if err != nil {
		fmt.Println("Failed to load data with error: ", err.Error())
	}

	data = strings.ReplaceAll(strings.ReplaceAll(data, "\n", ""), "\r", "")

	codeRanges := strings.Split(data, ",")

	var passedValues []int

	for _, codeRange := range codeRanges {
		codeRageStartEnd := strings.Split(codeRange, "-")

		start, err := strconv.Atoi(codeRageStartEnd[0])

		if err != nil {
			fmt.Println("Failed to parse start string: ", codeRageStartEnd[0], "into a int. Error", err.Error())
			return
		}

		end, err := strconv.Atoi(codeRageStartEnd[1])

		if err != nil {
			fmt.Println("Failed to parse end string: ", codeRageStartEnd[1], "into a int. Error", err.Error())
			return
		}

		for i := start; i <= end; i++ {
			for j := 1; j < len(strconv.Itoa(i)); j++ {
				testDigits := strconv.Itoa(i)[0:j]

				temp := SplitDigits(i, j)

				pass := false
				for _, item := range temp {
					if testDigits != item {
						pass = false
						break
					} else {
						pass = true
					}
				}

				if pass {
					passedValues = append(passedValues, i)
					break
				}
			}
		}
	}
	fmt.Println(SumSlice(passedValues))
}
