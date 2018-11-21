package channel

import (
	"sync"
)

type Channel struct {
}

type Processor interface {
	Run()
}

var once sync.Once
var instance *Channel

// New returns the singleton instance
func New(channelAction ChannelAction) Processor {
	once.Do(func() {
		instance = new(Channel)
	})

	return instance
}

func (s *Channel) Run() {}