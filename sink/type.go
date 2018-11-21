package sink

type SinkType int

const (
	Log		SinkType = iota
	File	SinkType = iota
)