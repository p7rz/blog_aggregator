package config

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func getConfigFilePath() (string, error) {
	home_dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home_dir, configFileName), nil
}

func Read() (Config, error) {
	path, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	jsonFile, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}
	defer jsonFile.Close()

	val, err := io.ReadAll(jsonFile)
	var config Config
	err = json.Unmarshal(val, &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}

func write(cfg Config) error {
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}

	dat, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	return os.WriteFile(path, dat, 0644)
}

func (c *Config) SetUser(name string) error {
	c.CurrentUserName = name
	return write(*c)
}
