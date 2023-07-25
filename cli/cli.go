package squarecli

import (
	"github.com/squarecloudofc/cli/cli/client"
	"github.com/squarecloudofc/cli/cli/config"
)

type SquareCloudCli struct {
	Client *client.APIClient
	Config *config.Config
}

func NewSquareCloudCli() *SquareCloudCli {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	rest := client.NewClient(cfg)

	return &SquareCloudCli{
		Client: rest,
	}
}
