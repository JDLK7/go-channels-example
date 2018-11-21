package sink

import (
	"github.com/sirupsen/logrus"
	"github.com/JDLK7/go-channels-example/model"
	"sync"
	"time"
)

type Formatter interface {
	Format(channel, quit chan model.Journey)
}

type LogSink struct {}

var once sync.Once
var instance Formatter

func New(sinkType SinkType) Formatter {
	once.Do(func() {
		switch sinkType {
		case Log:
			instance = new(LogSink)
		}
	})

	return instance
}

func (l *LogSink) Format(channel, quit chan model.Journey) {
	startTime := time.Now().UnixNano()

	isOpen := true
	for isOpen {
		select {
		case inboundJourney := <- channel:
			currentTime := (time.Now().UnixNano() - startTime) / int64(time.Millisecond)
			logrus.Printf("[T-%v] New journey arrived: %v\n", currentTime, inboundJourney)
		case <- quit:
			isOpen = false
			currentTime := (time.Now().UnixNano() - startTime) / int64(time.Millisecond)
			logrus.Printf("[T-%v] Quit\n", currentTime)
		}
	}
}