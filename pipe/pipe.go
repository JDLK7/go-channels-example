package pipe

import (
	"github.com/JDLK7/go-channels-example/config"
	"github.com/JDLK7/go-channels-example/model"
	"sync"
)

type Callable func(model.Journey) model.Journey

type Pipe struct {}
type Executor interface {
	exec(chan model.Journey, chan model.Journey, Callable) (chan model.Journey, chan model.Journey)
}

type PrefixPipe struct {
	executor Executor
}
type SufixPipe struct {
	executor Executor
}

type Processor interface {
	Run(chan model.Journey, chan model.Journey) (chan model.Journey, chan model.Journey)
	process(model.Journey) model.Journey
}

var once sync.Once
var instance Processor

// New returns the singleton instance
func New(configManager *config.ConfigManager) Processor {
	once.Do(func() {
		switch PipeAction(configManager.ProcessorAction) {
		case AddPrefix:
			pipeInstance := new(PrefixPipe)
			pipeInstance.executor = new(Pipe)
			instance = pipeInstance
		case AddSufix:
			pipeInstance := new(SufixPipe)
			pipeInstance.executor = new(Pipe)
			instance = pipeInstance
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
	return s.executor.exec(inChannel, inQuit, s.process)
}

func (s *SufixPipe) Run(inChannel, inQuit chan model.Journey) (chan model.Journey, chan model.Journey) {
	return s.executor.exec(inChannel, inQuit, s.process)
}

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
