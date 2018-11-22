package main

import (
	"github.com/JDLK7/go-channels-example/config"
	"github.com/JDLK7/go-channels-example/source"
	"github.com/JDLK7/go-channels-example/pipe"
	"github.com/JDLK7/go-channels-example/sink"
)

func init() {
	config.NewConfigManager()
}

func main() {
	journeys := []string{
		`{"id": 256, "time": 3600, "destination": "Madrid"}`,
		`{"id": 257, "time": 2400, "destination": "Madrid"}`,
		`{"id": 258, "time": 1800, "destination": "Madrid"}`,
		`{"id": 259, "time": 5000, "destination": "Madrid"}`,
		`{"id": 260, "time": 4000, "destination": "Madrid"}`,
	}

	configManager := config.ConfigManagerInstance

	source := source.New(configManager)
	pipe := pipe.New(configManager)
	sink := sink.New(configManager)

	sink.Format(
		pipe.Run(
			source.Run(journeys)))
}