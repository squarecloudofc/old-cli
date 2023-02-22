package appconfig

import (
	"os"
)

// squarecloud.config file validation (hope to try to add more functionality in the future)
func ValidateFile() (valid bool, errs []string) {	
	cfg, err := Get()
	if err != nil {
		errs = append(errs, "squarecloud.app file not found")
		return
	}

	// lets verify if main file is valid
	if _, err := os.Lstat(cfg.Main); err != nil {
		errs = append(errs, "main file not found")
	}

	if len(errs) < 1 {
		valid = true
	}

	return
}