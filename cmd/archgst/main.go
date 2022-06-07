package main

import (
	"log"

	junocmd "github.com/forbole/juno/v3/cmd"
	parsecmdtypes "github.com/forbole/juno/v3/cmd/parse/types"
	"github.com/giansalex/archway-gastracker/database"
	"github.com/giansalex/archway-gastracker/modules"
	"github.com/giansalex/archway-gastracker/types/config"
)

func main() {
	// Setup the config
	parseCfg := parsecmdtypes.NewConfig().
		WithRegistrar(modules.NewModulesRegistrar()).
		WithEncodingConfigBuilder(config.MakeEncodingConfig).
		WithDBBuilder(database.Builder)

	cfg := junocmd.NewConfig("archgst").
		WithParseConfig(parseCfg)

	// Run the commands and panic on any error
	executor := junocmd.BuildDefaultExecutor(cfg)
	err := executor.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
