package main

import (
	"encoding/json"

	lib "git.sr.ht/~u472892/libiscdhcpd"
	"github.com/gofiber/fiber/v2"
)

func apiPing(ctx *fiber.Ctx) error {
	json.NewEncoder(ctx).Encode(CONFIG)
	return nil
}

func apiView(ctx *fiber.Ctx) error {
	cfg, err := lib.LoadConfigFromFile(CONFIG.Instances[0].ConfigurationFile)
	if err != nil {
		json.NewEncoder(ctx).Encode(err)
		return err
	}
	json.NewEncoder(ctx).Encode(cfg)
	return nil
}
