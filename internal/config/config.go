package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
)

type Config struct {
	Token string `json:"token"`
}

func init() {
	if !fileExists(ConfigPath()) {
		err := createConfig(Config{
			Token: "",
		})
		if err != nil {
			panic(err)
		}
	}
}

func fileExists(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}

func DataDir() string {
	var path string

	if a := os.Getenv("AppData"); runtime.GOOS == "windows" && a != "" {
		path = filepath.Join(a, "Square Cloud CLI")
	} else if b := os.Getenv("XDG_CONFIG_HOME"); b != "" {
		path = filepath.Join(b, "squarecloud")
	} else {
		dir, err := os.UserHomeDir()
		if err != nil {
			panic("cannot get a valid dir for square cloud config")
		}
		path = filepath.Join(dir, ".config", "squarecloud")
	}

	return path
}

func ConfigPath() string {
	return filepath.Join(DataDir(), "config.json")
}

func GetConfigBytes() ([]byte, error) {
	return os.ReadFile(ConfigPath())
}

func GetConfig() (cfg *Config, err error) {
	configBytes, err := GetConfigBytes()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(configBytes, &cfg)
	if err != nil {
		return nil, err
	}
	return
}

func createConfig(config Config) error {
	var cfg []byte
	var err error

	if cfg, err = json.Marshal(config); err != nil {
		return err
	}

	if err = writeFile(ConfigPath(), cfg); err != nil {
		return err
	}

	return nil
}

func UpdateConfig(updateConfig func(cf *Config)) error {
	var cfg []byte
	var err error
	oldConfig, err := GetConfig()
	if err != nil {
		panic(err)
	}

	updateConfig(oldConfig)

	if cfg, err = json.Marshal(oldConfig); err != nil {
		return err
	}

	if err = writeFile(ConfigPath(), cfg); err != nil {
		return err
	}

	return nil

}

func writeFile(filename string, data []byte) error {
	err := os.MkdirAll(filepath.Dir(filename), 0771)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(data)
	return err
}
