package source

import (
	"encoding/json"
	"sync"
	"fmt"
	"github.com/JDLK7/go-channels-example/model"
)

type JSONSource struct {
}

type XMLSource struct {
}

type Ingestor interface {
	Run(journeys []string) (out chan model.Journey, quit chan int)
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


func (s *JSONSource) Run(journeys []string) (out chan model.Journey, quit chan int) {
	out = make(chan model.Journey)
	quit = make(chan int)

	go func() {
		for _, journey := range journeys {
			var parsedJourney model.Journey
			if err := json.Unmarshal([]byte(journey), &parsedJourney); err != nil {
				fmt.Printf("Unmarshall error: %v\n", err)
			} else {
				out <- parsedJourney
			}
		}

		quit <- 0
	}()

	return
}

func (s *XMLSource) Run(journeys []string) (out chan model.Journey, quit chan int) {
	return
}
