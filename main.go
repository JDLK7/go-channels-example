package main

import (
	"github.com/JDLK7/go-channels-example/pipe"
	"time"
	"fmt"
	"github.com/JDLK7/go-channels-example/source"
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

	out, quit := pipe.Run(source.Run(journeys))

	startTime := time.Now().UnixNano()

	isOpen := true
	for isOpen {
		select {
		case inboundJourney := <- out:
			currentTime := (time.Now().UnixNano() - startTime) / int64(time.Millisecond)
			fmt.Printf("[T-%v] New journey arrived: %v\n", currentTime, inboundJourney)
		case <- quit:
			isOpen = false
			currentTime := (time.Now().UnixNano() - startTime) / int64(time.Millisecond)
			fmt.Printf("[T-%v] Quit\n", currentTime)
		}
	}
}