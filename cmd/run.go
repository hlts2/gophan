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
	output string
	jsFile string
)

func init() {
	runCmd.PersistentFlags().StringVarP(&output, "out", "o", "capture.png", "Specify the output file of the capture result")
	runCmd.PersistentFlags().StringVarP(&jsFile, "set", "s", "", "Set Custom Javascript file for phantomjs")
}

func run(args []string) error {
	if len(args) == 0 {
		return errors.New("Argument does not exist")
	}

	p := phantom.NewPhantom()
	err := p.Exec(append([]string{jsFile, args[0]}, output))
	if err != nil {
		return err
	}

	return nil
}
