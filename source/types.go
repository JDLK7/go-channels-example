package source

type SourceType int

const (
	// Default ingest method
	Default SourceType = iota
	JSON		SourceType = iota
	XML			SourceType = iota
)