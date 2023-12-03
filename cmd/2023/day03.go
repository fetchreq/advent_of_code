/*
Copyright © 23 NAME HERE <EMAIL ADDRESS>
*/
package twentythree

import (
	"fmt"

	"github.com/rjprice04/advent_of_code/cmd"
	"github.com/spf13/cobra"
)

// day03Cmd represents the day03 command
var day03Cmd = &cobra.Command{
	Use:   "day03",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("day03 called")
	},
}

func init() {
	cmd.TwentyThreeCmd.AddCommand(day03Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// day03Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// day03Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
