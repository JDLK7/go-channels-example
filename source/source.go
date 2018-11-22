package source

import (
	"github.com/JDLK7/go-channels-example/config"
	"github.com/JDLK7/go-channels-example/model"
	"encoding/json"
	"encoding/xml"
	"time"
	"sync"
	"fmt"
)

// shutdownDelay constant time (in miliseconds) that Run waits for sending a quit message
const shutdownDelay = 10

type Callable func(string) (model.Journey, error)

type Source struct {}
type Executor interface {
	exec([]string, Callable) (chan model.Journey, chan model.Journey)
}

type JSONSource struct {
	executor Executor
}

type XMLSource struct {
	executor Executor
}

type Ingestor interface {
	Run([]string) (chan model.Journey, chan model.Journey)
	parse(string) (model.Journey, error)
}

var once sync.Once
var instance Ingestor

// New returns the singleton instance
func New(configManager *config.ConfigManager) Ingestor {
	once.Do(func() {
		
		switch SourceType(configManager.IngestorType) {
		case Default:
		case JSON:
			sourceInstance := new(JSONSource)
			sourceInstance.executor = new(Source)
			instance = sourceInstance
		case XML:
			sourceInstance := new(XMLSource)
			sourceInstance.executor = new(Source)
			instance = sourceInstance
		}
	})

	return instance
}

func channelAfterDelay(channel chan model.Journey, subject model.Journey, delay int) {
	time.Sleep(time.Duration(delay) * time.Millisecond)
	channel <- subject
}

func (s *JSONSource) parse(subject string) (parsedSubject model.Journey, err error) {
	err = json.Unmarshal([]byte(subject), &parsedSubject)
	return
}

func (s *XMLSource) parse(subject string) (parsedSubject model.Journey, err error) {
	err = xml.Unmarshal([]byte(subject), &parsedSubject)
	return
}

func (s *JSONSource) Run(journeys []string) (chan model.Journey, chan model.Journey) {
	return s.executor.exec(journeys, s.parse)
}

func (s *XMLSource) Run(journeys []string) (chan model.Journey, chan model.Journey) {
	return s.executor.exec(journeys, s.parse)
}

func (s *Source) exec(journeys []string, parse Callable) (out chan model.Journey, quit chan model.Journey) {
	out = make(chan model.Journey)
	quit = make(chan model.Journey)

	go func() {
		var maxDelay int

		for _, journey := range journeys {
			parsedJourney, err := parse(journey)
			if err != nil {
				fmt.Printf("Unmarshall error: %v\n", err)
			} else {
				go channelAfterDelay(out, parsedJourney, parsedJourney.Time)

				if parsedJourney.Time > maxDelay {
					maxDelay = parsedJourney.Time
				}
			}
		}

		go channelAfterDelay(quit, model.Journey{}, maxDelay + shutdownDelay)
	}()

	return
}
