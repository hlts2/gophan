package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/hlts2/gophan/phantom"
	cli "github.com/spf13/cobra"
)

var runCmd = &cli.Command{
	Use:   "run",
	Short: "Run binary of phantomjs internally",
	Run: func(cmd *cli.Command, args []string) {
		if err := run(args); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var (
	output string
	jsFile string
	query  string
)

func init() {
	RootCmd.AddCommand(runCmd)

	runCmd.PersistentFlags().StringVarP(&output, "out", "o", "capture.png", "Specify the output file of the capture result")
	runCmd.PersistentFlags().StringVarP(&jsFile, "set", "s", "", "Set Custom Javascript file for phantomjs")
	runCmd.PersistentFlags().StringVarP(&query, "query", "q", "", "Set CSS selector")
}

func run(args []string) error {
	if len(args) == 0 {
		return runUsage()
	}

	p, err := phantom.NewPhantom()
	if err != nil {
		return err
	}

	err = p.Exec(append([]string{jsFile, args[0]}, output, query))
	if err != nil {
		return err
	}

	return nil
}

func runUsage() error {
	return errors.New(`gophan run [URL or HTML] <option>

Available Options:
  -o, --out   Set the output file of the capture result (available extensions are png, jpg, pdf etc)
              $ gophan run [URL or HTML] -o output/capture.png

  -s, --set   Set custom javascript for phantomjs
              $ gophan run [URL or HTML] -s custom/load.js

  -q, --query Set CSS selector
              $ gophan run [URL or HTML] -q "#main"
	`)
}
