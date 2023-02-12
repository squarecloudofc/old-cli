package cmd

import (
	"squarecloud/internal/rest"

	"github.com/spf13/cobra"
)

var aboutmeCmd = &cobra.Command{
	Use:   "aboutme",
	Short: "View your Square Cloud user information",
	Long:  ``,

	Run: func(cmd *cobra.Command, args []string) {
		user, err := rest.GetSelfUser()
		if err != nil {
			cmd.PrintErrln(err)
			return
		}

		cmd.Printf("ID: %s\n", user.User.ID)
		cmd.Printf("TAG: %s\n", user.User.Tag)
		cmd.Printf("Plan Name: %s\n", user.User.Plan.Name)
		cmd.Printf("Plan Duration: %s\n", user.User.Plan.Duration.Formatted)
		cmd.Printf("Memory Total: %d\n", user.User.Plan.Memory.Limit)
		cmd.Printf("Memory Available: %d\n", user.User.Plan.Memory.Available)
		cmd.Printf("Memory Used: %d\n", user.User.Plan.Memory.Used)
		cmd.Printf("Is Blocked: %t\n", user.User.Blocklist)
	},
}

func init() {
	rootCmd.AddCommand(aboutmeCmd)
}
