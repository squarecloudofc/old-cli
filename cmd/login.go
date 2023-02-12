package cmd

import (
	"fmt"
	"regexp"
	"squarecloud/internal/config"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in for full SquareCloud CLI access",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		token := ""
		question := []*survey.Question{
			{
				Name: "token",
				Prompt: &survey.Password{
					Message: "Insert your API/CLI token:",
				},
				Validate: func(val interface{}) error {
					if match, _ := regexp.MatchString("\\d{17,21}-[\\w-]{60,80}", val.(string)); !match {
						return fmt.Errorf("your token seems to be invalid")
					}

					// nothing was wrong
					return nil
				},
			},
		}
		survey.Ask(question, &token)

		config.UpdateConfig(func(config *config.Config) {
			config.Token = token
		})
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
