package cmd

import (
	"fmt"
	"os"

	cli "github.com/spf13/cobra"
)

var RootCmd = &cli.Command{
	Use:   "gophan",
	Short: "Go commandline wrapper for phantomjs ",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
