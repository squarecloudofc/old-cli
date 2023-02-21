package cmd

import (
	"fmt"
	"squarecloud/internal/appconfig"

	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check your configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := appconfig.Get()
		if err != nil {
			cmd.PrintErrln(err)
			return
		}

		isValid, errs := appconfig.Validate(cfg)
		if !isValid || len(errs) > 0 {
			cmd.Println("Some errors were found in your configuration file")
			for _, err := range errs {
				cmd.Println(fmt.Sprintf(" - %s", err.Error()))
			}
			return
		}

		cmd.Println("Your configuration file is valid!")
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
