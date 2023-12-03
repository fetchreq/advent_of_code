/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package fifteen

import (
	"fmt"
	"strings"

	"github.com/rjprice04/advent_of_code/cmd"
	"github.com/rjprice04/advent_of_code/util"
	"github.com/spf13/cobra"
)

// day15Cmd represents the day15 command
var day15Cmd = &cobra.Command{
	Use:   "day15",
	Short: "Aoc 2015 day15",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day15 called")
		input := util.ReadFile("2015", "15", false)

		fmt.Println("Part 1: %d", day15Part1(input))
		fmt.Println("Part 2: %d", day15Part2(input))
	},
}

func init() {
	cmd.FifteenCmd.AddCommand(day15Cmd)
}



func day15Part1(input string) int {

	vals := [][]int{}

	for _, line := range(strings.Split(input, "\n")) {
		var name string
		var capacity, durability, flavor, texture, calories int
		fmt.Sscanf(line, "%s capacity %d, durability %d, flavor %d, texture %d, calories %d", &name, &capacity, &durability, &flavor, &texture, &calories)
		vals = append(vals, []int{capacity, durability, flavor, texture, calories})
	}
	best := 0
	for ing1 := 0; ing1 < 100; ing1++ {
		for ing2 := 0; ing2 < 100; ing2++ {

			for ing3 := 0; ing3 < 100; ing3++ {
				ing4 := 100 - ing1 - ing2 - ing3
				cap := util.Max(0, ing1 * vals[0][0] + ing2 * vals[1][0] + ing3 * vals[2][0] + ing4 * vals[3][0])
				dur := util.Max(0, ing1 * vals[0][1] + ing2 * vals[1][1] + ing3 * vals[2][1] + ing4 * vals[3][1])
				fla := util.Max(0, ing1 * vals[0][2] + ing2 * vals[1][2] + ing3 * vals[2][2] + ing4 * vals[3][2])
				tex := util.Max(0, ing1 * vals[0][3] + ing2 * vals[1][3] + ing3 * vals[2][3] + ing4 * vals[3][3])
				//cal := ing1 * vals[0][4] + ing2 * vals[1][4]

				total := cap * dur * fla * tex;
				if total > best {
					best = total

				}

			}
		}

	}
	return best
}

func day15Part2(input string) int {

	vals := [][]int{}

	for _, line := range(strings.Split(input, "\n")) {
		var name string
		var capacity, durability, flavor, texture, calories int
		fmt.Sscanf(line, "%s capacity %d, durability %d, flavor %d, texture %d, calories %d", &name, &capacity, &durability, &flavor, &texture, &calories)
		vals = append(vals, []int{capacity, durability, flavor, texture, calories})
	}
	best := 0
	for ing1 := 0; ing1 < 100; ing1++ {
		for ing2 := 0; ing2 < 100; ing2++ {

			for ing3 := 0; ing3 < 100; ing3++ {
				ing4 := 100 - ing1 - ing2 - ing3
				cap := util.Max(0, ing1 * vals[0][0] + ing2 * vals[1][0] + ing3 * vals[2][0] + ing4 * vals[3][0])
				dur := util.Max(0, ing1 * vals[0][1] + ing2 * vals[1][1] + ing3 * vals[2][1] + ing4 * vals[3][1])
				fla := util.Max(0, ing1 * vals[0][2] + ing2 * vals[1][2] + ing3 * vals[2][2] + ing4 * vals[3][2])
				tex := util.Max(0, ing1 * vals[0][3] + ing2 * vals[1][3] + ing3 * vals[2][3] + ing4 * vals[3][3])
				cal := ing1 * vals[0][4] + ing2 * vals[1][4] + ing3 * vals[2][4] +  ing4 * vals[3][4]

				total := cap * dur * fla * tex;
				if cal == 500 {
					best = util.Max(best, total)
				}

			}
		}

	}
	return best
}
