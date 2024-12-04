/*
Copyright © 23 NAME HERE <EMAIL ADDRESS>
*/
package twentythree

import (
	"fmt"
	"sort"
	"strings"

	"github.com/fetchreq/advent_of_code/cast"
	"github.com/fetchreq/advent_of_code/cmd"
	"github.com/fetchreq/advent_of_code/util"
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
		fmt.Printf("Part 1: %d\n", day7Part1(input))
		fmt.Printf("Part 2: %d\n", day7Part2(input))
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
	cards    string
	bid      int
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

var cardOrder []string = []string{
	"A",
	"K",
	"Q",
	"T",
	"9",
	"8",
	"7",
	"6",
	"5",
	"4",
	"3",
	"2",
	"J",
}
var cardValsUpdated map[string]int = map[string]int{
	"A": 13,
	"K": 12,
	"Q": 11,
	"T": 10,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
	"J": 1,
}

var handOrders []string = []string{"highCard", "onePair", "twoPair", "threeOfKind", "fullHouse", "fourOfKind", "fiveOfKind"}

func day7Part1(input string) int {
	winnings := 0
	handMap := make(map[string][]hand)

	for _, row := range strings.Split(input, "\n") {

		rowFields := strings.Split(row, " ")
		bid := cast.ToInt(rowFields[1])
		cards := rowFields[0]
		handType := getHandTypeFromCards(cards)

		if _, ok := handMap[handType]; !ok {
			handMap[handType] = []hand{}
		}
		handMap[handType] = append(handMap[handType], hand{cards: cards, bid: bid})
	}

	rank := 1
	for _, handOrder := range handOrders {
		vals, _ := handMap[handOrder]

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
			winnings += (hand.bid * rank)
			rank += 1
		}

	}

	return winnings
}

func day7Part2(input string) int {
	winnings := 0
	handMap := make(map[string][]hand)
	rank := 1
	for _, row := range strings.Split(input, "\n") {

		rowFields := strings.Split(row, " ")
		bid := cast.ToInt(rowFields[1])
		cards := rowFields[0]
		handType := getHandTypeFromCardWithUpgrades(cards)

		if _, ok := handMap[handType]; !ok {
			handMap[handType] = []hand{}
		}
		handMap[handType] = append(handMap[handType], hand{cards: cards, bid: bid})
	}
	for _, handOrder := range handOrders {
		vals, _ := handMap[handOrder]

		sort.SliceStable(vals, func(i, j int) bool {
			currCards := strings.Split(vals[i].cards, "")
			nextCards := strings.Split(vals[j].cards, "")
			k := 0
			for currCards[k] == nextCards[k] {
				k++
			}

			return cardValsUpdated[currCards[k]] < cardValsUpdated[nextCards[k]]
		})

		for _, hand := range vals {
			winnings += (hand.bid * rank)
			rank += 1
		}

	}

	return winnings
}

func getHandTypeFromCards(hand string) string {
	cardCount := make([]int, 13)
	for _, key := range cardOrder {
		if strings.Contains(hand, key) {
			cardCount = append(cardCount, getCardCount(hand, key))
		} else {
			cardCount = append(cardCount, 0)
		}
	}

	return getHandType(cardCount)
}

func getHandTypeFromCardWithUpgrades(hand string) string {

	cardCount := make([]int, 13)
	for _, key := range cardOrder {
		if strings.Contains(hand, key) {
			cardCount = append(cardCount, getCardCount(hand, key))
		} else {
			cardCount = append(cardCount, 0)
		}
	}

	jokerCount := cardCount[len(cardCount)-1]
	if jokerCount == 5 || jokerCount == 4 {
		return "fiveOfKind"
	}

	handType := getHandType(cardCount[:len(cardCount)-1])

	if jokerCount > 0 {
		handType = upgradeHand(hand, handType, jokerCount)
	}

	return handType
}

func upgradeHand(inputHand, currHand string, jokerCount int) string {
	upgrades := map[string]string{
		"fourOfKind_1":  "fiveOfKind",
		"threeOfKind_2": "fiveOfKind",
		"threeOfKind_1": "fourOfKind",
		"onePair_3":     "fiveOfKind",
		"onePair_2":     "fourOfKind",
		"onePair_1":     "threeOfKind",
		"highCard_3":    "fourOfKind",
		"highCard_2":    "threeOfKind",
		"highCard_1":    "onePair",
		"twoPair_1":     "fullHouse",
	}

	newHand, ok := upgrades[fmt.Sprintf("%s_%d", currHand, jokerCount)]
	if !ok {
		fmt.Printf("%s upgrade not found for handtype %s with %d jokers\n", inputHand, currHand, jokerCount)
		return currHand
	}

	return newHand

}

func getHandType(cardCount []int) string {
	handType := "highCard"
	handTypeSet := false
	for i := 0; i < len(cardCount) && !handTypeSet; i++ {
		if cardCount[i] == 0 {
			continue
		}
		if cardCount[i] == 5 {
			handType = "fiveOfKind"
		} else if cardCount[i] == 4 {
			handType = "fourOfKind"
		}

		if cardCount[i] == 3 {
			for j := i; j < len(cardCount); j++ {
				if cardCount[j] == 0 {
					continue
				}
				if cardCount[j] == 2 {
					handType = "fullHouse"
					handTypeSet = true
				}
			}
			if !handTypeSet {
				handType = "threeOfKind"
				handTypeSet = true
			}

		} else if cardCount[i] == 2 {
			for j := i + 1; j < len(cardCount); j++ {
				if cardCount[j] == 0 {
					continue
				}
				if cardCount[j] == 3 {
					handType = "fullHouse"
					handTypeSet = true
				} else if cardCount[j] == 2 {
					handType = "twoPair"
					handTypeSet = true
				}
			}
			if !handTypeSet {
				handType = "onePair"
				handTypeSet = true
			}
		}

	}
	return handType
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
