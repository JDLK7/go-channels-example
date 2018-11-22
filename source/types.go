package source

type SourceType string

const (
	// Default ingest method
	Default SourceType = "Default"
	JSON		SourceType = "JSON"
	XML			SourceType = "XML"
)