package source

import (
	"github.com/JDLK7/go-channels-example/config"
	"github.com/JDLK7/go-channels-example/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type SourceTestSuite struct {
	suite.Suite
	config *config.ConfigManager
}

func (suite *SourceTestSuite) SetupTest() {
	suite.config = &config.ConfigManager{
		IngestorType:			"JSON",
	}
}

// TestRunReturnsActiveChannel checks that the channel returned by Run sends journeys.
func (suite *SourceTestSuite) TestRunReturnsActiveChannel() {
	journeys := []string{`{"id": 1, "time": 100}`}
	source := New(suite.config)

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

	assert.True(suite.T(), isActive)
}

func (suite *SourceTestSuite) TestRunChannelSendsJourneys() {
	journeys := []string{`{"id": 256, "time": 3600}`}
	expectedJourney := model.Journey{
		Id: 	256,
		Time:	3600,
	}

	source := New(suite.config)

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

	assert.Equal(suite.T(), expectedJourney, actualJourney)
}

func (suite *SourceTestSuite) TestRunChannelSendsEveryJourney() {
	journeys := []string{`{"id": 256, "time": 360}`, `{"id": 257, "time": 240}`, `{"id": 258, "time": 180}`}
	expectedJourneys := []model.Journey{
		model.Journey { Id: 256, Time: 360 },
		model.Journey { Id: 257, Time: 240 },
		model.Journey { Id: 258, Time: 180 },
	}
	var actualJourneys []model.Journey

	isOpen := true
	source := New(suite.config)
	channel, quit := source.Run(journeys)

	for isOpen {
		select {
		case inboundJourney := <- channel:
			actualJourneys = append(actualJourneys, inboundJourney)
		case <- quit:
			isOpen = false
		}
	}

	assert.Equal(suite.T(), len(expectedJourneys), len(actualJourneys))
}

func (suite *SourceTestSuite) TestRunChannelSendsJourneysSortedByTime() {
	journeys := []string{`{"id": 256, "time": 360}`, `{"id": 257, "time": 240}`, `{"id": 258, "time": 180}`}
	expectedJourneys := []model.Journey{
		model.Journey { Id: 258, Time: 180 },
		model.Journey { Id: 257, Time: 240 },
		model.Journey { Id: 256, Time: 360 },
	}
	var actualJourneys []model.Journey

	isOpen := true
	source := New(suite.config)
	channel, quit := source.Run(journeys)

	for isOpen {
		select {
		case inboundJourney := <- channel:
			actualJourneys = append(actualJourneys, inboundJourney)
		case <- quit:
			isOpen = false
		}
	}

	assert.Equal(suite.T(), expectedJourneys, actualJourneys)
}

// TestSourceTestSuite runs the suite
func TestSourceTestSuite(t *testing.T) {
	suite.Run(t, new(SourceTestSuite))
}
