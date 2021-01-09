package main

import (
	"flag"

	"github.com/DionTech/cgrep/pckg/grep"
)

func main() {
	var help bool
	flag.BoolVar(&help, "help", false, "show help")

	var save bool
	flag.BoolVar(&save, "save", false, "save the actual expression")
	flag.BoolVar(&save, "s", false, "short for save")

	var show bool
	flag.BoolVar(&show, "list-templates", false, "list the available templates")

	var name string
	flag.StringVar(&name, "name", "", "name of the expresion to save")
	flag.StringVar(&name, "n", "", "short for name")

	var expression string
	flag.StringVar(&expression, "expression", "", "name of the stored expression to use")
	flag.StringVar(&expression, "e", "", "short for expression")

	var path string
	flag.StringVar(&path, "path", "", "path where to search for")
	flag.StringVar(&path, "p", "", "short for path")

	var threads int
	flag.IntVar(&threads, "threads", 1, "how many threads should be used")
	flag.IntVar(&threads, "t", 1, "short for threads")

	flag.Parse()

	if help {
		flag.PrintDefaults()
		return
	}

	if show {
		grep.PrintTemplates()

		return
	}

	expr := flag.Arg(0)

	if expression != "" {
		tmpl, err := grep.LoadExpression(expression)
		if err != nil {
			return
		}
		expr = tmpl
	}

	if save {
		grep.SaveExpression(name, expr)

		return
	}

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
