package main

import (
	"encoding/json"
	"os"
)

type AppConfig struct {
	AppName      string     `json:"appname",omitempty`
	DatabasePath string     `json:"dbpath",omitempty`
	DebugLevel   DebugLevel `json:"debug_level",omitempty`
	Version      string     `json:"version",omitempty`
	Host         string     `json:"host",omitempty`
	Port         int        `json:"port",omitempty`
}

func loadAppConfigDefaults(config AppConfig) AppConfig {
	config.AppName = APPNAME
	config.DatabasePath = ""
	config.DebugLevel = DebugLevelInfo
	config.Version = version()
	config.Host = "127.0.0.1"
	config.Port = 9091

	return config
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
