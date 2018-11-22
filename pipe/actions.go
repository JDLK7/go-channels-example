package pipe

type PipeAction string

const (
	// AddPrefix adds a prefix to an incoming message
	AddPrefix PipeAction = "AddPrefix"
	// AddSufix adds a sufix to an incoming message
	AddSufix 	PipeAction = "AddSufix"
)