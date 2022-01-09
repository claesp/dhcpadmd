package main

import (
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
)

var (
	APP_NAME = "dhcpdadmd"
	MAJOR    = 0
	MINOR    = 1
	REVISION = 2201091308
	CONFIG   AppConfig
)

func version() string {
	return fmt.Sprintf("%d.%d.%d", MAJOR, MINOR, REVISION)
}

func main() {
	log.Printf("started")

	rh := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/api/v1/ping":
			apiPing(ctx)
		default:
			ctx.Error("Unsupported path", fasthttp.StatusNotFound)
		}
	}

	var cfgDefaultErr error
	CONFIG, cfgDefaultErr = loadAppConfigDefaults(CONFIG)
	if cfgDefaultErr != nil {
		log.Fatalln("configuration loading failed:", cfgDefaultErr)
	}

	var cfgFileErr error
	cfgFilename := fmt.Sprintf("%s.conf", APP_NAME)
	CONFIG, cfgFileErr = loadAppConfigFromFile(CONFIG, cfgFilename)
	if cfgFileErr != nil {
		log.Println(fmt.Sprintf("configuration file '%s' loading failed:", cfgFilename), cfgFileErr)
	}

	log.Printf("listening on port %d\n", CONFIG.Port)
	log.Fatalln(fasthttp.ListenAndServe(fmt.Sprintf(":%d", CONFIG.Port), rh))
}
