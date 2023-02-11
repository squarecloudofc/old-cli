package auth

import "squarecloud/internal/config"

func IsAuthenticated() bool {
	config, err := config.GetConfig()
	if err != nil {
		return false
	}

	return config.Token != ""
}
