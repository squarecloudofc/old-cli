package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/squarecloudofc/cli/cli/config/store"
)

const (
	CONFIG_FILENAME = "squarecloud.json"
)

func getFileDir() string {
	datadir := store.GetDataDir()

	return filepath.Join(datadir, CONFIG_FILENAME)
}

func Load() (*Config, error) {
	file := getFileDir()
	bytes, err := os.ReadFile(file)
	if err != nil {
		if _, err = os.Stat(file); err != nil {
			data, err := json.Marshal(
				Config{
					Identity: Identity{},
				},
			)
			if err != nil {
				return nil, err
			}

			err = os.WriteFile(file, data, 0777)
			if err != nil {
				return nil, err
			}

			bytes = data
		} else {
			return nil, err
		}
	}

	var data Config
	json.Unmarshal(bytes, &data)

	return &data, nil
}

func Update(c *Config) error {
	file, err := os.OpenFile(getFileDir(), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	var cfg []byte
	if cfg, err = json.Marshal(c); err != nil {
		return err
	}

	_, err = file.Write(cfg)
	return err
}
