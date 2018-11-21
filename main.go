package main

import (
	"github.com/JDLK7/go-channels-example/source"
	"github.com/JDLK7/go-channels-example/pipe"
	"github.com/JDLK7/go-channels-example/sink"
)

func main() {
	journeys := []string{
		`{"id": 256, "time": 3600, "destination": "Madrid"}`,
		`{"id": 257, "time": 2400, "destination": "Madrid"}`,
		`{"id": 258, "time": 1800, "destination": "Madrid"}`,
		`{"id": 259, "time": 5000, "destination": "Madrid"}`,
		`{"id": 260, "time": 4000, "destination": "Madrid"}`,
	}
	source := source.New(source.JSON)
	pipe := pipe.New(pipe.AddPrefix)
	sink := sink.New(sink.Log)

	out, quit := pipe.Run(source.Run(journeys))

	sink.Format(out, quit)
}