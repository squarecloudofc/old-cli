package app

import (
	squarecli "github.com/squarecloudofc/cli/cli"
	"github.com/urfave/cli/v2"
)

func NewStatusCommand(scli *squarecli.SquareCloudCli) *cli.Command {
	cmd := &cli.Command{
		Name:  "status",
		Usage: "Manage your Square Cloud applications",
		Action: func(ctx *cli.Context) error {
			return runStatusCommand(scli)
		},
	}

	return cmd
}

func runStatusCommand(scli *squarecli.SquareCloudCli) error {
	scli.Client.AppStatus("11ae97ef5512423089336d041fced46c")

	return nil
}
