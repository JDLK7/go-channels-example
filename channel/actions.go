package channel

type ChannelAction string

const (
	// Default ingest method
	Default 	ChannelAction = "default"
	AddPrefix ChannelAction = "add_prefix"
	AddSufix 	ChannelAction = "add_sufix"
)