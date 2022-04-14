/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/DionTech/cgrep/pckg/grep"

	"github.com/spf13/cobra"
)

// findCmd represents the find command
var findCmd = &cobra.Command{
	Use:   "find",
	Short: "find some files by expression",
	Long: `find some files by expression, for example:
	find -p ./ -e .go
`,
	Run: func(cmd *cobra.Command, args []string) {
		scan := &grep.Scan{
			Path:       path,
			Expression: expression,
			Mode:       "find",
			Threads:    threads}

		scan.Run()
	},
}

func init() {
	rootCmd.AddCommand(findCmd)

	findCmd.Flags().StringVarP(&expression, "expression", "e", "", "expression to search")
	findCmd.Flags().StringVarP(&path, "path", "p", "", "path to search")

	err := findCmd.MarkFlagRequired("expression")
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	err = findCmd.MarkFlagRequired("path")
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
}
