package main

import (
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
)

var (
	APPNAME  = "dhcpdadmd"
	MAJOR    = 0
	MINOR    = 1
	REVISION = 220624
	CONFIG   AppConfig
)

type DebugLevel int

const (
	DebugLevelDebug DebugLevel = iota
	DebugLevelInfo
	DebugLevelWarning
	DebugLevelCritical
)

func version() string {
	return fmt.Sprintf("%d.%d.%d", MAJOR, MINOR, REVISION)
}

func out(level DebugLevel, section string, text string) {
	if CONFIG.DebugLevel <= level {
		log.Printf("%s: %s: %s\n", CONFIG.AppName, section, text)
	}
}

func main() {
	CONFIG = loadAppConfigDefaults(CONFIG)
	out(DebugLevelInfo, "main", fmt.Sprintf("starting version %s", version()))

	var cfgFileErr error
	cfgFilename := fmt.Sprintf("%s.conf", APPNAME)
	CONFIG, cfgFileErr = loadAppConfigFromFile(CONFIG, cfgFilename)
	if cfgFileErr != nil {
		out(DebugLevelCritical, "main", fmt.Sprintf("configuration file '%s' loading failed: %s", cfgFilename, cfgFileErr))
	}

	rh := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/api/v1/ping":
			apiPing(ctx)
		default:
			ctx.Error("Unsupported path", fasthttp.StatusNotFound)
		}
	}

	out(DebugLevelInfo, "main", fmt.Sprintf("listening on %s:%d", CONFIG.Host, CONFIG.Port))
	log.Fatalln(fasthttp.ListenAndServe(fmt.Sprintf("%s:%d", CONFIG.Host, CONFIG.Port), rh))
}
