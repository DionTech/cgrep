/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/DionTech/cgrep/pckg/grep"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list your stored expressions",
	Long:  `list your stored expressions`,
	Run: func(cmd *cobra.Command, args []string) {
		grep.PrintTemplates()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
