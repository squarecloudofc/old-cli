package cmd

import (
	"fmt"
	"squarecloud/internal/appconfig"
	"squarecloud/pkg/groupwriter"

	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check your configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		// will get the file based on cwd dir
		isValid, errs := appconfig.ValidateFile()
		w := groupwriter.NewWriter(cmd.OutOrStdout(), 4)
		defer w.Flush()

		if !isValid || len(errs) > 0 {
			w.Group("\x1b[33m[!]\x1b[0m squarecloud.config")
			for _, err := range errs {
				w.Println(fmt.Sprintf("\x1b[31mX\x1b[0m %s", err))
			}
			w.Endgroup()
			return
		}
		w.Println("\x1b[32m[•]\x1b[0m squarecloud.config")
		

		w.Println("\n\x1b[32m•\x1b[0m Is everythink okay!")
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
