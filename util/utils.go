package util

import (
	"os"
)

func ReadFile(year string, day string) string {
	path := "./input/" + year + "/" + day + ".txt"
	data, err := os.ReadFile(path)

	if (err != nil) {
		panic(err)
	}

	return string(data)
}
