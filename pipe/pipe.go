package pipe

import (
	"github.com/JDLK7/go-channels-example/config"
	"github.com/JDLK7/go-channels-example/model"
	"sync"
)

type Pipe struct {}
type Executor interface {
	exec(inChannel, inQuit chan model.Journey, process func(model.Journey) model.Journey) (outChannel, outQuit chan model.Journey)
}

type PrefixPipe struct {
	Executor
}
type SufixPipe struct {
	Executor
}

type Processor interface {
	Run(inChannel, inQuit chan model.Journey) (chan model.Journey, chan model.Journey)
	process(journey model.Journey) model.Journey
}

var once sync.Once
var instance Processor

// New returns the singleton instance
func New(configManager *config.ConfigManager) Processor {
	once.Do(func() {
		switch PipeAction(configManager.ProcessorAction) {
		case AddPrefix:
			instance = new(PrefixPipe)
		case AddSufix:
			instance = new(SufixPipe)
		}
	})

	return instance
}

func (s *PrefixPipe) process(journey model.Journey) model.Journey {
	journey.Destination = "cabify_" + journey.Destination
	return journey
}

func (s *SufixPipe) process(journey model.Journey) model.Journey {
	journey.Destination = journey.Destination + "_cabify"
	return journey
}

func (s *PrefixPipe) Run(inChannel, inQuit chan model.Journey) (chan model.Journey, chan model.Journey) {
	return s.Executor.exec(inChannel, inQuit, s.process)
}

func (s *SufixPipe) Run(inChannel, inQuit chan model.Journey) (chan model.Journey, chan model.Journey) {
	return s.Executor.exec(inChannel, inQuit, s.process)
}

type Callable func(model.Journey) model.Journey

func (p *Pipe) exec(inChannel, inQuit chan model.Journey, process Callable) (outChannel, outQuit chan model.Journey) {
	outChannel = make(chan model.Journey)
	outQuit = make(chan model.Journey)

	go func() {
		isOpen := true
		for isOpen {
			select {
			case journey := <- inChannel:
				outChannel <- process(journey)
			case q := <- inQuit:
				isOpen = false
				outQuit <- q
			}
		}
	}()

	return
}
