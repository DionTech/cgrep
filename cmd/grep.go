/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/DionTech/cgrep/pckg/grep"
	"github.com/spf13/cobra"
)

// grepCmd represents the grep command
var grepCmd = &cobra.Command{
	Use:   "grep",
	Short: "grep files or stdin for a pattern",
	Long: `grep files or stdin for a pattern, defining amount of threads you are willing to use. Some examples: 
	grep --path ./ --threads 10 "^open"
  	grep --path ./ "^open"

	portscan scan -i 127.0.0.1 -t 20 | {your-alias-here} grep --threads=20 ^open
`,
	Run: func(cmd *cobra.Command, args []string) {
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

		if path == "" {
			grep.Grep(expr, threads)
			return
		}

		scan := &grep.Scan{
			Path:       path,
			Expression: expr,
			Filter:     filter,
			Threads:    threads}

		scan.Run()
	},
}

func init() {
	rootCmd.AddCommand(grepCmd)

	grepCmd.Flags().IntVarP(&threads, "threads", "t", 1, "number of threads")
	grepCmd.Flags().StringVarP(&path, "path", "p", ".", "path to search")
	grepCmd.Flags().StringVarP(&expression, "expression", "e", "", "expression to grep for")
	grepCmd.Flags().StringVarP(&filter, "filter", "f", "", "filtering the path to decide whether to search or not")
}
