package main

import (
	"github.com/DionTech/cgrep/pckg/grep"
	"github.com/devfacet/gocmd"
)

func main() {
	flags := struct {
		Help bool `short:"h" long:"help" description:"Display usage" global:"false"`
		Scan struct {
			Path       string `short:"p" long:"path" description:"path to the root folder" required:"true" nonempty:"false"`
			Expression string `short:"e" long:"expr" description:"expression to search for" required:"true" nonempty:"false"`
			Threads    int    `short:"t" long:"threads" description:"amount of threads being used"  required:"true" nonempty:"false"`
		} `command:"scan" description:"search an expression" nonempty:"false"`
	}{}

	gocmd.HandleFlag("Scan", func(cmd *gocmd.Cmd, args []string) error {

		scan := &grep.Scan{
			Path:       flags.Scan.Path,
			Expression: flags.Scan.Expression,
			Threads:    flags.Scan.Threads}

		scan.Run()

		return nil
	})

	gocmd.New(gocmd.Options{
		Name:        "cgrep",
		Version:     "1.0.0",
		Description: "concurrent find",
		Flags:       &flags,
		ConfigType:  gocmd.ConfigTypeAuto,
	})
}
