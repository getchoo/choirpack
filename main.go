package main

import (
	"fmt"
	"os"

	"github.com/ryanccn/choirpack/cmd"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "choirpack",
}

func main() {
	rootCmd.AddCommand(cmd.UseCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
