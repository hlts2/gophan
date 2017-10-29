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
	Short: "",
	Run: func(cmd *cli.Command, args []string) {
		if err := run(args); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var (
	outPut string
	//filters []string
	jsFile string
)

func init() {
	runCmd.PersistentFlags().StringVarP(&outPut, "out", "o", "capture.png", "Specify the output file of the capture result")
	runCmd.PersistentFlags().StringVarP(&jsFile, "set", "s", "", "Set Custom Javascript file for phantomjs")
	//runCmd.PersistentFlags().StringArrayVarP(&filters, "filter", "f", []string{}, "Retrieve the filtered html element")
}

func run(args []string) error {
	if len(args) == 0 {
		return errors.New("Argument does not exist")
	}

	p := phantom.NewPhantom()
	p.Exec(append([]string{jsFile}, args...))

	return nil
}
