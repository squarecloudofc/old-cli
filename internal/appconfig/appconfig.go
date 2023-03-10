package appconfig

import (
	"errors"
	"os"
	"squarecloud/pkg/properties"
)

type AppConfig struct {
	ID          string `properties:"ID"`
	DisplayName string `properties:"DISPLAY_NAME"`
	Description string `properties:"DESCRIPTION"`
	Main        string `properties:"MAIN"`
	Version     string `properties:"VERSION"`
	Memory      string `properties:"MEMORY"`
	Subdomain   string `properties:"SUBDOMAIN"`
	Start       string `properties:"START"`
	AutoRestart string `properties:"AUTO_RESTART"`
}

func Get() (*AppConfig, error) {
	bytes, err := getConfigBytes()
	if err != nil {
		return nil, err
	}

	cfg, err := readFromBytes(bytes)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func Save(cfg *AppConfig) error {
	file, err := getConfigFile()
	if err != nil {
		return err
	}

	data, err := properties.Marshal(cfg)
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func getConfigFile() (*os.File, error) {
	names := []string{"squarecloud.config", "squarecloud.app"}

	for _, name := range names {
		_, err := os.Lstat(name)
		if err != nil {
			continue
		}

		file, err := os.OpenFile(name, os.O_RDWR, 0644)
		if err != nil {
			return nil, err
		}
		return file, nil
	}

	return nil, errors.New("squarecloud.app not found")
}

func getConfigBytes() ([]byte, error) {
	file, err := getConfigFile()
	if err != nil {
		return nil, err
	}

	data := make([]byte, 0)
	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err != nil {
			break
		}
		data = append(data, buf[:n]...)
	}

	return data, nil
}

func readFromBytes(data []byte) (*AppConfig, error) {
	var d AppConfig
	if err := properties.Unmarshall(data, &d); err != nil {
		return nil, err
	}

	return &d, nil
}
