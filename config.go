package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type AppConfig struct {
	AppName      string     `json:"appname",omitempty`
	DatabasePath string     `json:"dbpath",omitempty`
	DebugLevel   DebugLevel `json:"debug_level",omitempty`
	Host         string     `json:"host",omitempty`
	InstanceName string     `json:"instance",omitempt`
	Port         int        `json:"port",omitempty`
	Version      string     `json:"version",omitempty`
}

func loadAppConfigDefaults(config AppConfig) AppConfig {
	config.AppName = APPNAME
	config.DebugLevel = DebugLevelDebug
	config.Host = "127.0.0.1"
	config.InstanceName = "default"
	config.DatabasePath = fmt.Sprintf("./%s", config.InstanceName)
	config.Port = 9091
	config.Version = version()

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

	if cfg.InstanceName != "" {
		config.InstanceName = cfg.InstanceName
	}

	if cfg.DatabasePath != "" {
		config.DatabasePath = cfg.DatabasePath
	}

	if cfg.DebugLevel != 0 {
		if cfg.DebugLevel > DebugLevelCritical {
			config.DebugLevel = DebugLevelCritical
		} else {
			config.DebugLevel = cfg.DebugLevel
		}
	}

	if cfg.Host != "" {
		config.Host = cfg.Host
	}

	if cfg.Port != 0 {
		config.Port = cfg.Port
	}

	return config, nil
}
