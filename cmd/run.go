package cmd

import (
	"fmt"
	"os"

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

func run(args []string) error {
	return nil
}
