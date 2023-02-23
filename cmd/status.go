package cmd

import (
	"fmt"
	"squarecloud/internal/api"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status [appid]",
	Short: "Get your application status",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		app, err := api.ApplicationStatus(args[0])
		if err != nil {
			cmd.PrintErrln()
			return
		}

		w := tabwriter.NewWriter(cmd.OutOrStdout(), 10, 0, 0, ' ', tabwriter.TabIndent)
		defer w.Flush()

		fmt.Fprintln(w, "Application ID", "\t", "STATUS", "\t", "CPU", "\t", "RAM", "\t", "STORAGE", "\t", "NETWORK")
		fmt.Fprintln(w, args[0], "\t", app.Status, "\t", app.CPU, "\t", app.RAM, "\t", app.Storage, "\t", app.Network.Total)
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
