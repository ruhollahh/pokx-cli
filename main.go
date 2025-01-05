package main

import (
	"github.com/ruhollahh/pokx-cli/internal/clients/pokeapi"
	"time"
)

func main() {
	cfg := &config{
		pokeapiClient: pokeapi.NewClient(5*time.Second, 5*time.Minute),
	}
	startRepl(cfg)
}
