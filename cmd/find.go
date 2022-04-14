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
	find -p ./ -e "stored-expression-name"
	find ".go"

`,
	Run: func(cmd *cobra.Command, args []string) {
		if path == "" {
			path = "."
		}

		var expr string
		if len(args) > 0 {
			expr = args[0]
		}

		//we will load a template file
		if expression != "" {
			tmpl, err := grep.LoadExpression(expression)
			if err != nil {
				return
			}
			expr = tmpl
		}

		if expr == "" {
			fmt.Println("grep: missing expression to grep for")
			fmt.Println("Try 'grep --help' for more information.")
			return
		}

		scan := &grep.Scan{
			Path:       path,
			Expression: expr,
			Mode:       "find",
			Threads:    threads}

		scan.Run()
	},
}

func init() {
	rootCmd.AddCommand(findCmd)

	findCmd.Flags().StringVarP(&path, "path", "p", "", "path to search; when not set it will be the current directory")
	findCmd.Flags().IntVarP(&threads, "threads", "t", 1, "number of threads")
}
