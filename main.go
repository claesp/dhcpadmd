package main

import (
	"fmt"
	"log"

	"git.mills.io/prologic/bitcask"
	"github.com/valyala/fasthttp"
)

var (
	APPNAME  = "dhcpdadmd"
	MAJOR    = 0
	MINOR    = 1
	REVISION = 220624
	CONFIG   AppConfig
	DATABASE *bitcask.Bitcask
)

func version() string {
	return fmt.Sprintf("%d.%d.%d", MAJOR, MINOR, REVISION)
}

func config() {
	CONFIG = loadAppConfigDefaults(CONFIG)

	var err error
	CONFIG, err = loadAppConfigFromFile(CONFIG, fmt.Sprintf("%s.conf", APPNAME))
	if err != nil {
		out(DebugLevelCritical, "main", fmt.Sprintf("loading configuration file '%s' failed: %s", fmt.Sprintf("%s.conf", APPNAME), err))
	}
}

func server() {
	rh := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/api/v1/ping":
			apiPing(ctx)
		default:
			ctx.Error("Unsupported path", fasthttp.StatusNotFound)
		}
	}

	out(DebugLevelInfo, "server", fmt.Sprintf("listening on %s:%d", CONFIG.Host, CONFIG.Port))
	log.Fatalln(fasthttp.ListenAndServe(fmt.Sprintf("%s:%d", CONFIG.Host, CONFIG.Port), rh))
}

func main() {
	config()

	out(DebugLevelInfo, "main", fmt.Sprintf("starting version %s", version()))
	out(DebugLevelInfo, "main", fmt.Sprintf("current debug level is %s", CONFIG.DebugLevel))

	var dbErr error
	DATABASE, dbErr = bitcask.Open(CONFIG.DatabasePath)
	if dbErr != nil {
		out(DebugLevelInfo, "main", fmt.Sprintf("unable to open database file '%s': %s", CONFIG.DatabasePath, dbErr))
	}
	defer DATABASE.Close()

	server()
}
