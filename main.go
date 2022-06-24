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

func version() string {
	return fmt.Sprintf("%d.%d.%d", MAJOR, MINOR, REVISION)
}

func api() {
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

func main() {
	CONFIG = loadAppConfigDefaults(CONFIG)

	var err error
	cfg := fmt.Sprintf("%s.conf", APPNAME)
	CONFIG, err = loadAppConfigFromFile(CONFIG, cfg)
	if err != nil {
		out(DebugLevelCritical, "main", fmt.Sprintf("loading configuration file '%s' failed: %s", cfg, err))
	}

	out(DebugLevelInfo, "main", fmt.Sprintf("starting version %s", version()))
	out(DebugLevelInfo, "main", fmt.Sprintf("current debug level is %s", CONFIG.DebugLevel))

	api()
}
