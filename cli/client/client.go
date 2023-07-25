package client

import "github.com/squarecloudofc/cli/cli/config"

type APIClient struct {
	Config *config.Config
}

func NewClient(config *config.Config) *APIClient {
	return &APIClient{
		Config: config,
	}
}
