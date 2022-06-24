package main

import "log"

type DebugLevel int

func (d DebugLevel) String() string {
	switch d {
	case DebugLevelUndefined:
		return "UNDEFINED"
	case DebugLevelDebug:
		return "DEBUG"
	case DebugLevelInfo:
		return "INFO"
	case DebugLevelWarning:
		return "WARNING"
	case DebugLevelCritical:
		return "CRITICAL"
	}
	return "UNKNOWN"
}

const (
	DebugLevelUndefined DebugLevel = iota
	DebugLevelDebug
	DebugLevelInfo
	DebugLevelWarning
	DebugLevelCritical
)

func out(level DebugLevel, section string, text string) {
	if CONFIG.DebugLevel <= level {
		log.Printf("%s: %s: %s\n", CONFIG.AppName, section, text)
	}
}
