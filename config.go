package main

import (
	"encoding/json"
	"os"
)

type AppConfig struct {
	AppName      string `json:"appname",omitempty`
	DatabasePath string `json:"dbpath",omitempty`
	Version      string `json:"version",omitempty`
	Port         int    `json:"port",omitempty`
}

func loadAppConfigDefaults(config AppConfig) (AppConfig, error) {
	config.AppName = APP_NAME
	config.DatabasePath = ""
	config.Version = version()
	config.Port = 9091

	return config, nil
}

func loadAppConfigFromFile(config AppConfig, filename string) (AppConfig, error) {
	d, err := os.ReadFile(filename)
	if err != nil {
		return config, err
	}

	var cfg AppConfig
	jsonErr := json.Unmarshal(d, &cfg)
	if jsonErr != nil {
		return config, jsonErr
	}

	config.Port = cfg.Port
	config.DatabasePath = cfg.DatabasePath

	return config, nil
}
