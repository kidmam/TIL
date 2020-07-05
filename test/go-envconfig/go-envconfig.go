package main

import (
	"context"
	"log"

	"github.com/sethvargo/go-envconfig"
)

type MyConfig struct {
	Port     int    `env:"PORT"`
	Username string `env:"USERNAME"`
}

func main() {
	ctx := context.Background()

	var c MyConfig
	if err := envconfig.Process(ctx, &c); err != nil {
		log.Fatal(err)
	}

	// c.Port = 5555
	// c.Username = "yoyo"
}
