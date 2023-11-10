package util

import (
	"math"
	"os"
	"strings"
)

// Reads a file from the input directory
// The year is the sub year folder
// day is the input for the day
// test is if we should use test file or real input
func ReadFile(year string, day string, test bool) string {
	path := "./input/" + year + "/" + day
	if test {
		path += ".test"
	} else {
		path += ".txt"
	}

	data, err := os.ReadFile(path)
		
	CheckErr(err)
	strContent := string(data)
	return strings.TrimRight(strContent, "\n")

}

func CheckErr(e error) {
	if (e != nil) {
		panic(e)
	}
}


func Min(args ...int) int {
	min := args[0]

	for _, val := range args {
		if val < min {
			min = val
		}
	}
	return min
}

func Max(args ...int) int {
	max := args[0]

	for _, val := range args {
		if val > max {
			max = val
		}
	}
	return max
}

func Min2(args ...int) (int, int){
	min1 := math.MaxInt
	min2 := math.MaxInt

	for _, val := range args {
		if val < min1 {
			min2 = min1
			min1 = val
		} else if val < min2 {
			min2 = val
		}
	}

	return min1, min2
}
