package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type AppConfigInstance struct {
	Name              string `json:"name"`
	ConfigurationFile string `json:"configuration_file"`
}

type AppConfig struct {
	AppName      string              `json:"appname",omitempty`
	DatabasePath string              `json:"dbpath",omitempty`
	DebugLevel   DebugLevel          `json:"debug_level",omitempty`
	Host         string              `json:"host",omitempty`
	Agent        string              `json:"agent",omitempt`
	Port         int                 `json:"port",omitempty`
	Version      string              `json:"version",omitempty`
	Started      time.Time           `json:"started",omitempty`
	Instances    []AppConfigInstance `json:"instances",omitempt`
}

func loadAppConfigDefaults(config AppConfig) AppConfig {
	config.AppName = APPNAME
	config.DebugLevel = DebugLevelDebug
	config.Host = "127.0.0.1"
	config.Agent = "default"
	config.DatabasePath = fmt.Sprintf("./%s", config.Agent)
	config.Port = 9091
	config.Version = version()
	config.Started = time.Now()

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

	if cfg.Agent != "" {
		config.Agent = cfg.Agent
		config.DatabasePath = fmt.Sprintf("./%s", config.Agent)
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

	if len(cfg.Instances) > 0 {
		config.Instances = cfg.Instances
	}

	return config, nil
}
