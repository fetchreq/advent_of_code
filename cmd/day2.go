/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rjprice04/advent_of_code/util"
	"github.com/spf13/cobra"
)

// day2Cmd represents the day2 command
var day2Cmd = &cobra.Command{
	Use:   "day2",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		input := util.ReadFile("2015", "day2", false)
		fmt.Printf("Part 1 %d\n", day2Part1(input))
		fmt.Printf("Part 2 %d\n", day2Part2(input))
	},
}

func init() {
	fifteenCmd.AddCommand(day2Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day2Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type Dim struct {
	length int
	width int
	height int
}

func newDim(input string) *Dim {
	vals := strings.Split(input, "x");

	l, err := strconv.Atoi(vals[0])
	util.CheckErr(err)

	w, err := strconv.Atoi(vals[1])
	util.CheckErr(err)

	h, err := strconv.Atoi(vals[2])
	util.CheckErr(err)

	return &Dim{length: l, width: w, height: h}
}

func (d *Dim) surfaceArea() int {
	return (2 * d.length * d.width) + (2 * d.width * d.height) + (2 * d.height * d.length)
}
func (d *Dim) smallestArea() int {
	return util.Min((d.length * d.width), (d.width * d.height), (d.height * d.length))
}

func (d *Dim) smallestParameter() int {
	min1, min2 := util.Min2(d.length, d.width, d.height);

	return 2 * min1 + 2 * min2;
}

func (d *Dim) cubicVolume() int {
	return d.length * d.width * d.height;
}






func day2Part1(input string) int {
	dims := strings.Split(input, "\n")
	total := 0
	for _, dim := range dims {

		if len(strings.TrimSpace(dim)) == 0 {
			continue
		}
		curr := newDim(dim)

		total += curr.surfaceArea() + curr.smallestArea()
	}

	return total

}

func day2Part2(input string) int {

	dims := strings.Split(input, "\n")
	total := 0
	for _, dim := range dims {

		if len(strings.TrimSpace(dim)) == 0 {
			continue
		}

		curr := newDim(dim)

		total += curr.smallestParameter() + curr.cubicVolume()
	}

	return total

}
