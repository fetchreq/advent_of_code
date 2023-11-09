/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// 2015Cmd represents the 2015 command
var FifteenCmd = &cobra.Command{
	Use:   "2015",
	Short: "Top Level Command for AoC 2015",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("2015 called")
	},
}

func init() {
	rootCmd.AddCommand(FifteenCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// 2015Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// 2015Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
