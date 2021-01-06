package main

import (
	"flag"

	"github.com/DionTech/cgrep/pckg/grep"
)

func main() {
	var path string
	flag.StringVar(&path, "path", "", "path where to search for")

	var threads int
	flag.IntVar(&threads, "threads", 1, "how many threads should be used")

	flag.Parse()

	expr := flag.Arg(0)

	if path == "" {
		grep.Grep(expr, threads)

		return
	}

	scan := &grep.Scan{
		Path:       path,
		Expression: expr,
		Threads:    threads}

	scan.Run()
}
