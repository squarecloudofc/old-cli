package commands

import (
	squarecli "github.com/squarecloudofc/cli/cli"
	"github.com/urfave/cli/v2"
)

func NewInitCommand(scli *squarecli.SquareCloudCli) *cli.Command {
	cmd := &cli.Command{
		Name:  "login",
		Usage: "Setup your squarecloud.app config",
		Action: func(ctx *cli.Context) error {
			return runInitCommand(scli, ctx)
		},
	}

	return cmd
}

func runInitCommand(scli *squarecli.SquareCloudCli, ctx *cli.Context) error {

	return nil
}
