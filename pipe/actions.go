package pipe

type PipeAction int

const (
	// Default ingest method
	Default 	PipeAction = iota
	AddPrefix PipeAction = iota
	AddSufix 	PipeAction = iota
)