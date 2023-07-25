package commands

import (
	"github.com/squarecloudofc/cli/cli"
	"github.com/urfave/cli/v2"
)

func NewAppsCommand(client *squarecli.SquareCloudCli) *cli.Command {
	cmd := &cli.Command{
		Name:  "apps",
		Usage: "List your Square Cloud applications",
		Action: func(ctx *cli.Context) error {
			return runAppsCommand(client, ctx)
		},
	}

	return cmd
}

func runAppsCommand(client *squarecli.SquareCloudCli, ctx *cli.Context) error {
	return nil
}
