package cmd

import (
	"fmt"
	"os"
	"squarecloud/internal/rest"
	"strconv"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var applistCmd = &cobra.Command{
	Use:   "apps",
	Short: "See all your active apps on Square Cloud",
	Run: func(cmd *cobra.Command, args []string) {
		user, err := rest.GetSelfUser()
		if err != nil {
			cmd.PrintErrln(err.Error())
			return
		}
		if len(user.Applications) < 1 {
			fmt.Println("You does not have any application active")
			return
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 0, ' ', tabwriter.TabIndent)
		fmt.Fprintln(w, "Application ID", "\t", "RAM", "\t", "CLUSTER", "\t", "LANG", "\t", "TYPE", "\t", "IS WEBSITE?", "\t")
		defer w.Flush()

		for _, app := range user.Applications {
			fmt.Fprintln(w,
				app.ID[:12], "\t",
				strconv.Itoa(app.RAM)+"mb", "\t",
				app.Cluster, "\t",
				app.Lang, "\t",
				app.Type, "\t",
				app.IsWesite, "\t",
			)
		}
	},
}

func init() {
	rootCmd.AddCommand(applistCmd)
}
