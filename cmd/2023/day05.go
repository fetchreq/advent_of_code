/*
Copyright © 23 NAME HERE <EMAIL ADDRESS>
*/
package twentythree

import (
	"fmt"
	"math"
	"strings"

	"github.com/rjprice04/advent_of_code/cast"
	"github.com/rjprice04/advent_of_code/cmd"
	"github.com/rjprice04/advent_of_code/util"
	"github.com/spf13/cobra"
)

// day05Cmd represents the day05 command
var day05Cmd = &cobra.Command{
	Use:   "day05",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day05 called")
		input := util.ReadFile("2023", "5", false)
		fmt.Printf("Part 1: %d", day5Part1(input))
		fmt.Printf("Part 2: %d", day5Part2(input))

	},
}

func init() {
	cmd.TwentyThreeCmd.AddCommand(day05Cmd)
}

type seedMapValues struct {
	srcStart int
	destStart int
	size int
}

var order = []string{
	"seed-to-soil",
	"soil-to-fertilizer",
	"fertilizer-to-water",
	"water-to-light",
	"light-to-temperature",
	"temperature-to-humidity",
	"humidity-to-location",
}

func day5Part1(input string) int {
	lowest := math.MaxInt
	seeds := []int{}
	maps := make(map[string][]seedMapValues)
	for idx, chunk := range strings.Split(input, "\n\n") {

		if idx == 0 {
			rawSeed := strings.TrimPrefix(chunk, "seeds: ")
			for _, seed := range strings.Fields(rawSeed) {
				seeds = append(seeds, cast.ToInt(seed))
			}
			continue
		}

		mapType, content, _:= strings.Cut(chunk, " map:")
		var seedMapVals []seedMapValues
		for _, row := range strings.Split(content, "\n") {
			if strings.TrimSpace(row) == "" {
				continue
			}
			var src, dest, size int
			fmt.Sscanf(row, "%d %d %d", &dest, &src, &size)
			//fmt.Printf("%s: src Start %d - dest Start %d size %d\n", mapType, src, dest ,size) 
			seedMapVals = append(seedMapVals, seedMapValues{srcStart: src, destStart: dest, size: size})
		}
		maps[mapType] = seedMapVals
	}
	
	for _, seed := range seeds {
		curr := seed
		for _, mapType := range order {
			for _, seedMapVal := range maps[mapType] {
				if seedMapVal.srcStart <= curr && curr < seedMapVal.srcStart + seedMapVal.size {
					curr += (seedMapVal.destStart - seedMapVal.srcStart)
					break
				}
			}
		}
		if curr < lowest {
			lowest = curr	
		}
	}
	return lowest;

} 

type seedRange struct  {
	start, size int
}

func day5Part2(input string) int {
	lowest := math.MaxInt
	seeds := []seedRange{}
	maps := make(map[string][]seedMapValues)
	for idx, chunk := range strings.Split(input, "\n\n") {

		if idx == 0 {
			raw := strings.TrimPrefix(chunk, "seeds: ")
			rawSeeds :=  strings.Split(raw, " ");
			for i := 0; i < len(rawSeeds) - 1; i+=2 {
				seeds = append(seeds, seedRange{start: cast.ToInt(rawSeeds[i]), size: cast.ToInt(rawSeeds[i+1])})
			}
			continue
		}

		mapType, content, _:= strings.Cut(chunk, " map:")
		var seedMapVals []seedMapValues
		for _, row := range strings.Split(content, "\n") {
			if strings.TrimSpace(row) == "" {
				continue
			}
			var src, dest, size int
			fmt.Sscanf(row, "%d %d %d", &dest, &src, &size)

			seedMapVals = append(seedMapVals, seedMapValues{srcStart: src, destStart: dest, size: size})
		}
		maps[mapType] = seedMapVals
	}

	for _, seedRange := range seeds {
		//fmt.Printf("Looking at seeds %d to %d\n", seedRange.start, seedRange.start + seedRange.size)
		for i := seedRange.start; i < seedRange.start + seedRange.size; i++ {
			curr := i
			//fmt.Printf("Starting seed %d\n", i)
			for _, mapType := range order {
				for _, seedMapVal := range maps[mapType] {
					if seedMapVal.srcStart <= curr && curr < seedMapVal.srcStart + seedMapVal.size {
						curr += (seedMapVal.destStart - seedMapVal.srcStart)
						break
					}
				}
			}
			if curr < lowest {
				lowest = curr	
			}
		}
	}	
	return lowest;

}
