package main

import (
	"log"

	"github.com/danielgtaylor/huma/cli"
	"github.com/mrfrosty0430/photoflock/api"
	"github.com/mrfrosty0430/photoflock/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	app := cli.New(
		cli.WithTitle("PhotoFlock API"),
		cli.WithDescription("Photo sharing service with face recognition"),
		cli.WithVersion("1.0.0"),
	)

	api.RegisterRoutes(app)

	log.Printf("Server starting on %s", cfg.Server.Address)
	if err := app.Listen(cfg.Server.Address); err != nil {
		log.Fatal(err)
	}
}
