package source

import (
	"time"
	"encoding/json"
	"sync"
	"fmt"
	"github.com/JDLK7/go-channels-example/model"
)

// shutdownDelay constant time (in miliseconds) that Run waits for sending a quit message
const shutdownDelay = 10

type JSONSource struct {
}

type XMLSource struct {
}

type Ingestor interface {
	Run(journeys []string) (out chan model.Journey, quit chan model.Journey)
}

var once sync.Once
var instance Ingestor

// New returns the singleton instance
func New(sourceType SourceType) Ingestor {
	once.Do(func() {
		switch sourceType {
		case Default:
		case JSON:
			instance = new(JSONSource)
		case XML:
			instance = new(XMLSource)
		}
	})

	return instance
}

func channelAfterDelay(channel chan model.Journey, subject model.Journey, delay int) {
	time.Sleep(time.Duration(delay) * time.Millisecond)
	channel <- subject
}

func (s *JSONSource) Run(journeys []string) (out chan model.Journey, quit chan model.Journey) {
	out = make(chan model.Journey)
	quit = make(chan model.Journey)

	go func() {
		totalDelay := shutdownDelay

		for _, journey := range journeys {
			var parsedJourney model.Journey
			if err := json.Unmarshal([]byte(journey), &parsedJourney); err != nil {
				fmt.Printf("Unmarshall error: %v\n", err)
			} else {
				totalDelay += parsedJourney.Time
				go channelAfterDelay(out, parsedJourney, parsedJourney.Time)
			}
		}

		go channelAfterDelay(quit, model.Journey{}, totalDelay)
	}()

	return
}

func (s *XMLSource) Run(journeys []string) (out chan model.Journey, quit chan model.Journey) {
	return
}
