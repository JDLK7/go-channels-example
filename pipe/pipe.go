package pipe

import (
	"github.com/JDLK7/go-channels-example/model"
	"sync"
)

type PrefixPipe struct {
}

type Processor interface {
	Run(inChannel, inQuit chan model.Journey) (chan model.Journey, chan model.Journey)
}

var once sync.Once
var instance Processor

// New returns the singleton instance
func New(pipeAction PipeAction) Processor {
	once.Do(func() {
		switch pipeAction {
		case Default:
		case AddPrefix:
			instance = new(PrefixPipe)
		}
	})

	return instance
}

func (s *PrefixPipe) Run(inChannel, inQuit chan model.Journey) (outChannel, outQuit chan model.Journey) {
	outChannel = make(chan model.Journey)
	outQuit = make(chan model.Journey)

	go func() {
		isOpen := true
		for isOpen {
			select {
			case journey := <- inChannel:
				journey.Destination = "cabify_" + journey.Destination
				outChannel <- journey
			case q := <- inQuit:
				isOpen = false
				outQuit <- q
			}
		}
	}()

	return
}