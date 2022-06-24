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

func config() error {
	CONFIG = loadAppConfigDefaults(CONFIG)

	var err error
	CONFIG, err = loadAppConfigFromFile(CONFIG, fmt.Sprintf("%s.conf", APPNAME))
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := config()
	if err != nil {
		out(DebugLevelCritical, "main", fmt.Sprintf("loading configuration file failed: %s", err))
	}

	out(DebugLevelInfo, "main", fmt.Sprintf("starting version %s", version()))
	out(DebugLevelInfo, "main", fmt.Sprintf("current debug level is %s", CONFIG.DebugLevel))

	api()
}
