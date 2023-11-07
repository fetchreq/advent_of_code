/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"crypto/md5"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// day4Cmd represents the day4 command
var day4Cmd = &cobra.Command{
	Use:   "day4",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day4 called")
		input := "yzbqklnj"
		fmt.Printf("Part 1: %d", day4Part1(input));
		fmt.Printf("Part 2: %d", day4Part2(input));
	},
}

func init() {
	fifteenCmd.AddCommand(day4Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day4Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day4Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func day4Part1(input string) int {
	current := 0;
	prefix := strings.Repeat("0", 5);
	for {
		value := fmt.Sprintf("%s%d", input, current)
		result := fmt.Sprintf("%x", md5.Sum([]byte(value)));
		if strings.HasPrefix(result, prefix) {
			return current

		} 
		current += 1

	}
}

func day4Part2(input string) int {
	current := 0;
	prefix := strings.Repeat("0", 6);
	for {
		value := fmt.Sprintf("%s%d", input, current)
		result := fmt.Sprintf("%x", md5.Sum([]byte(value)));
		if strings.HasPrefix(result, prefix) {
			return current

		} 
		current += 1

	}
}

