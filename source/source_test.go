package source

import (
	"testing"
	"reflect"
	"github.com/JDLK7/go-channels-example/model"
)

// TestRunReturnsActiveChannel checks that the channel returned by Run sends journeys.
func TestRunReturnsActiveChannel(t *testing.T) {
	journeys := []string{`{"id": 1, "journey_time": 100}`}
	source := New(JSON)

	isActive := false
	isOpen := true
	channel, quit := source.Run(journeys)

	for isOpen {
		select {
		case <- channel:
			isActive = true
		case <- quit:
			isOpen = false
		}
	}

	if !isActive {
		t.Errorf("Channel returned is inactive")
	}
}

func TestRunChannelSendsJourneys(t *testing.T) {
	journeys := []string{`{"id": 256, "time": 3600}`}
	expectedJourney := model.Journey{
		Id: 	256,
		Time:	3600,
	}

	source := New(JSON)

	isOpen := true
	var actualJourney model.Journey
	channel, quit := source.Run(journeys)

	for isOpen {
		select {
		case inboundJourney := <- channel:
			actualJourney = inboundJourney
		case <- quit:
			isOpen = false
		}
	}

	if actualJourney.Id != expectedJourney.Id || actualJourney.Time != expectedJourney.Time {
		t.Errorf("Channel returned by Run doesn't send journeys.\nExpected journey: %v\nActual journey: %v", expectedJourney, actualJourney)
	}
}

func TestRunChannelSendsEveryJourney(t *testing.T) {
	journeys := []string{`{"id": 256, "time": 3600}`, `{"id": 257, "time": 2400}`, `{"id": 258, "time": 1800}`}
	expectedJourneys := []model.Journey{
		model.Journey { Id: 256, Time: 3600 },
		model.Journey { Id: 257, Time: 2400 },
		model.Journey { Id: 258, Time: 1800 },
	}
	var actualJourneys []model.Journey

	isOpen := true
	source := New(JSON)
	channel, quit := source.Run(journeys)

	for isOpen {
		select {
		case inboundJourney := <- channel:
			actualJourneys = append(actualJourneys, inboundJourney)
		case <- quit:
			isOpen = false
		}
	}

	equalJourneys := true

	if len(expectedJourneys) == len(actualJourneys) {
		for i := 0; i < len(expectedJourneys); i++ {
			if !reflect.DeepEqual(expectedJourneys[i], actualJourneys[i]) {
				equalJourneys = false
			}
		}
	} else {
		equalJourneys = false
	}

	if !equalJourneys {
		t.Errorf("Channel doesn't send all journeys.\nExpected journeys: %v\nActual journeys: %v", expectedJourneys, actualJourneys)
	}
}
