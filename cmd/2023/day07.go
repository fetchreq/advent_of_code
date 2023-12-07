/*
Copyright © 23 NAME HERE <EMAIL ADDRESS>
*/
package twentythree

import (
	"fmt"
	"sort"
	"strings"

	"github.com/rjprice04/advent_of_code/cast"
	"github.com/rjprice04/advent_of_code/cmd"
	"github.com/rjprice04/advent_of_code/util"
	"github.com/spf13/cobra"
)

// day07Cmd represents the day07 command
var day07Cmd = &cobra.Command{
	Use:   "day07",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day07 called")
		input := util.ReadFile("2023", "7", false)
		fmt.Printf("Part 1: %d", day7Part1(input))
	},
}

func init() {
	cmd.TwentyThreeCmd.AddCommand(day07Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day07Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day07Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type hand struct {
	cards string
	bid int
	handType string
}
var cardVals map[string]int = map[string]int{
	"A": 13,
	"K": 12,
	"Q": 11, 
	"J": 10, 
	"T": 9, 
	"9": 8, 
	"8": 7, 
	"7": 6, 
	"6": 5, 
	"5": 4, 
	"4": 3, 
	"3": 2, 
	"2": 1,
}
var handOrders []string = []string{"highHand", "onePair", "twoPair", "threeOfKind", "fullHouse", "fourOfKind", "fiveOfKind"}

func day7Part1(input string) int {
	winnings := 0;
	handMap := make(map[string][]hand)
	for _, row := range strings.Split(input, "\n") {

		rowFields := strings.Split(row, " ")
		bid := cast.ToInt(rowFields[1])
		handType := getHandTypeFromCards(rowFields[0])
		
		if _, ok := handMap[handType]; !ok {
			handMap[handType] = []hand{}
		}
		handMap[handType] = append(handMap[handType], hand{cards: rowFields[0], bid: bid})
	}
	rank := 1
	for _,handOrder := range handOrders {
		vals, _ := handMap[handOrder]
		if len(vals) == 0 {
			continue	
		}

		sort.SliceStable(vals, func(i, j int) bool {
			currCards := strings.Split(vals[i].cards, "")
			nextCards := strings.Split(vals[j].cards, "")
			k := 0
			for currCards[k] == nextCards[k] {
				k++
			}

			return cardVals[currCards[k]] < cardVals[nextCards[k]]
		})

		//fmt.Printf("Looking at %s with hands %v\n", handOrder, vals)
		for _, hand := range vals {
			//fmt.Printf("Bid %d rank %d\n", hand.bid, rank)
			winnings += hand.bid * rank
			rank++
		}

	}

	return winnings
}
func getHandTypeFromCards(hand string) string {
	cardCount := make([]int, 13)
	idx := 0
	for key := range cardVals {
		if strings.Contains(hand, key) {
			cardCount[idx] = getCardCount(hand, key)
		} else {
			cardCount[idx] = 0
		}
		idx++
	}
	for i := 0; i < len(cardCount); i++ {
		if cardCount[i] == 5 {
			return "fiveOfKind"
		} else if cardCount[i] == 4 {
			return "fourOfKind"
		} 

		if cardCount[i] == 3 {
			for j := i + 1; j < len(cardCount); j++ {
				if cardCount[j] == 2 {
					return "fullHouse"
				}
			}
			return "threeOfKind"
		} else if cardCount[i] == 2 {
			for j := i + 1; j < len(cardCount); j++ {
				if cardCount[j] == 3 {
					return "fullHouse"
				} else if cardCount[j] == 2 {
					return "twoPair"
				}
			}
			return "onePair"
		}

	}	

	

	return "highCard"
}

func getCardCount(hand string, searchCard string) int {
	count := 0
	for _, card := range strings.Split(hand, "") {
		if card == searchCard {
			count++
		}
	}

	return count
}
