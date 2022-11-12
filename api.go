package main

import (
	"encoding/json"

	lib "git.sr.ht/~u472892/libiscdhcpd"
	"github.com/valyala/fasthttp"
)

func apiPing(ctx *fasthttp.RequestCtx) {
	json.NewEncoder(ctx).Encode(CONFIG)
}

func apiView(ctx *fasthttp.RequestCtx) {
	cfg, err := lib.LoadConfigFromFile(CONFIG.Instances[0].ConfigurationFile)
	if err != nil {
		json.NewEncoder(ctx).Encode(err)
	}
	json.NewEncoder(ctx).Encode(cfg)
}
