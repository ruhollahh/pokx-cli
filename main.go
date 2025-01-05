package main

import (
	"github.com/ruhollahh/pokx-cli/clients/pokeapi"
	"time"
)

func main() {
	cfg := &config{
		pokeapiClient: pokeapi.NewClient(5 * time.Second),
	}
	startRepl(cfg)
}
