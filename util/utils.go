package util

import (
	"math"
	"os"
)

func ReadFile(year string, day string, test bool) string {
	path := "./input/" + year + "/" + day
	if test {
		path += ".test"
	} else {
		path += ".txt"
	}

	data, err := os.ReadFile(path)
		
	CheckErr(err)

	return string(data)
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
