package main

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

func apiPing(ctx *fasthttp.RequestCtx) {
	json.NewEncoder(ctx).Encode(CONFIG)
}
