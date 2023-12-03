/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package fifteen

import (
	"fmt"
	"strings"

	"github.com/rjprice04/advent_of_code/cast"
	"github.com/rjprice04/advent_of_code/cmd"
	"github.com/rjprice04/advent_of_code/util"
	"github.com/spf13/cobra"
)

// day14Cmd represents the day14 command
var day14Cmd = &cobra.Command{
	Use:   "day14",
	Short: "Aoc 2015 Day 14",
	Run: func(cmd *cobra.Command, args []string) {

		input := util.ReadFile("2015", "14", false)
		fmt.Printf("Part 1: %d\n", day14Part1(input))
	},
}

func init() {
	cmd.FifteenCmd.AddCommand(day14Cmd)
}

type Reindeer struct {
	distanceTraveled int
	startTravelTime int 
	endTravelTime int
	speed int
	travalInterval int
	resetInterval int
	resting bool
}

func newReindeer(parts []string) *Reindeer{
		// Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds
	speed := cast.ToInt(parts[3])
	travelInterval := cast.ToInt(parts[6])
	restInterval := cast.ToInt(parts[13])

	return &Reindeer{
		distanceTraveled: 0,
		startTravelTime: 0,
		endTravelTime: travelInterval,
		speed: speed,
		travalInterval: travelInterval,
		resetInterval: restInterval,
		resting: true,
	}
}

func (r *Reindeer) updatePosition(name string, currTime int) {
	if !r.resting && currTime > r.endTravelTime {
		fmt.Printf("%s starts resting at %d (current distance: %d) ", name, currTime, r.distanceTraveled)
		r.resting = true
		r.startTravelTime = r.resetInterval + currTime
		fmt.Printf(" will start flying again at %d\n", r.startTravelTime)
		return
	}  else if r.resting && currTime == r.startTravelTime {
		fmt.Printf("%s starts flying at %d", name, currTime)
		r.resting = false
		r.endTravelTime = r.travalInterval + r.startTravelTime
		fmt.Printf("will start resting after %d seconds (timestamp: %d)\n", r.travalInterval, r.endTravelTime)
		return
	}

	if !r.resting {
		r.distanceTraveled += r.speed
	}
}

func (r Reindeer) printReindeer(name string) {
	fmt.Printf("%s traveled %d ", name, r.distanceTraveled)
	if r.resting {
		fmt.Println("and is currently resting")
	} else {
		fmt.Println("and is NOT currently resting")
	}
}


func day14Part1(input string) int {
	reindeerMap := make(map[string]*Reindeer)
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, " ")
		name := parts[0]
		reindeerMap[name] = newReindeer(parts)
	}

	for i := 0; i < 2503; i++ {
		for k := range reindeerMap {
			reindeerMap[k].updatePosition(k, i)
		}
	}
	maxDistance := 0;
	for k := range reindeerMap {
		maxDistance = util.Max(maxDistance, reindeerMap[k].distanceTraveled)	
		reindeerMap[k].printReindeer(k);
	}



	return maxDistance

}
