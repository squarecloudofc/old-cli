package cmd

import (
	"fmt"
	"os"
	"squarecloud/internal/build"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "squarecloud",
	Version: build.Version,
	Long:    "squarego is Square Cloud on the command line. Which allows you to manage all your applications from the command line!",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
