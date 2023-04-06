package cmd

import (
	"fmt"

	"github.com/richaardev/squarecloud-cli/internal/build"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "get the cli version",
	Run: func(cmd *cobra.Command, args []string) {
		version := build.Version

		fmt.Printf("Square Cloud CLI version %s\n", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
