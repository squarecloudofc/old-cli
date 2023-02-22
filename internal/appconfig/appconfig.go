package appconfig

import (
	"errors"
	"os"
	"squarecloud/pkg/properties"
)

type AppConfig struct {
	DisplayName string `properties:"DISPLAY_NAME"`
	Description string `properties:"DESCRIPTION"`
	Main        string `properties:"MAIN"`
	Version     string `properties:"VERSION"`
	Memory      string `properties:"MEMORY"`
	Subdomain   string `properties:"SUBDOMAIN"`
	Start       string `properties:"START"`
}

func Get() (*AppConfig, error) {
	data, err := getFile()
	if err != nil {
		return nil, err
	}

	cfg, err := readFromBytes(data)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func getFile() ([]byte, error) {
	names := []string{"squarecloud.config", "squarecloud.app"}

	for _, name := range names {
		_, err := os.Lstat(name)
		if err != nil {
			continue
		}

		file, err := os.ReadFile(name)
		if err != nil {
			return nil, err
		}
		return file, nil
	}

	return nil, errors.New("squarecloud.app not found")
}

func readFromBytes(data []byte) (*AppConfig, error) {
	var d AppConfig
	if err := properties.Unmarshall(data, &d); err != nil {
		return nil, err
	}

	return &d, nil
}
