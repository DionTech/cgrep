/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/DionTech/cgrep/pckg/grep"
	"github.com/spf13/cobra"
)

var name string

// saveCmd represents the save command
var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save an expression you will often use.",
	Long: `Save an expression you will often use. For example: 
	save -n formAction -e "form(.?)action="`,
	Run: func(cmd *cobra.Command, args []string) {
		grep.SaveExpression(name, expression)
	},
}

func init() {
	rootCmd.AddCommand(saveCmd)

	saveCmd.Flags().StringVarP(&expression, "expression", "e", "", "expression to save")
	saveCmd.Flags().StringVarP(&name, "name", "n", "", "name of the expression")

	err := saveCmd.MarkFlagRequired("expression")

	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	err = saveCmd.MarkFlagRequired("name")

	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}
}
