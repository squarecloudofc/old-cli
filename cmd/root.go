package cmd

import (
	"fmt"
	"os"

	"github.com/richaardev/squarecloud-cli/internal/build"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "squarecloud",
	Version: build.Version,
	Long:    "Square Cloud CLI is a powerful tool that allows you to easily manage your applications hosted on Square Cloud, saving time and streamlining the process.",
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
