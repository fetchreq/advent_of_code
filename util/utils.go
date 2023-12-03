package util

import (
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// Reads a file from the input directory
// The year is the sub year folder
// day is the input for the day
// test is if we should use test file or real input
func ReadFile(year string, day string, test bool) string {
	path := "./input/" + year + "/day" + day
	if test {
		path += ".test"
	} else {
		path += ".txt"
	}

	data, err := os.ReadFile(path)
	
	// if we get an error make a request to get the data from AoC site
	if err != nil {
		fmt.Println("Input file Not Found")
		fmt.Println("Downloading..." )
		// probably need to download the file
		viper.SetConfigFile(".env")
		err := viper.ReadInConfig()
		CheckErr(err)

		session := viper.Get("SESSION").(string)
		url := fmt.Sprintf("https://adventofcode.com/%s/day/%s/input", year, day)

		client := &http.Client{}
		req, err := http.NewRequest("GET", url, nil)
		CheckErr(err)
		req.AddCookie(&http.Cookie{Name: "session", Value: session})
		

		res, err := client.Do(req) 
		CheckErr(err)

		resBody, err := io.ReadAll(res.Body)
		CheckErr(err)
		
		fmt.Println("Creating File..." )
		file, err := os.Create(path)
		CheckErr(err)
		defer file.Close()

		num, err := file.Write(resBody)
		CheckErr(err)
		fmt.Printf("Wrote %d bytes\n", num)
		data = resBody


	} else {
		fmt.Println("Using Local File")
	}
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
