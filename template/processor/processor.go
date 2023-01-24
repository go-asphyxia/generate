package processor

import (
	"github.com/go-asphyxia/data/tree/general"
)

type (
	Configuration struct {
		Selector string `json:"Selector"`
	}

	Processor struct {
		configuration Configuration
	}
)

func Constructor(configuration *Configuration) *Processor {
	return &Processor{
		configuration: *configuration,
	}
}

func (processor *Processor) Process(tree general.Tree[any]) (err error) {
	return
}
