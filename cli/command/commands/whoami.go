package commands

import (
	"fmt"

	squarecli "github.com/squarecloudofc/cli/cli"
	"github.com/urfave/cli/v2"
)

func NewWhoamiCommand(client *squarecli.SquareCloudCli) *cli.Command {
	cmd := &cli.Command{
		Name:  "whoami",
		Usage: "See which user are logged to Square Cloud CLI",
		Action: func(ctx *cli.Context) error {
			return runWhoamiCommand(client, ctx)
		},
	}

	return cmd
}

func runWhoamiCommand(client *squarecli.SquareCloudCli, ctx *cli.Context) error {
	if client.Config == nil {
		fmt.Println("User not logged")
		return nil
	}

	fmt.Printf("%s is currently logged", client.Config.Identity.Username)
	return nil
}
