/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// 2023Cmd represents the 2023 command
var TwentyThreeCmd = &cobra.Command{
	Use:   "2023",
	Short: "AoC 2023",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("2023 called")
	},
}

func init() {
	rootCmd.AddCommand(TwentyThreeCmd)
}
