package processor

import (
	"github.com/go-asphyxia/data/tree/general"
	"github.com/go-asphyxia/generate/template/source"
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

func (processor *Processor) Process(tree *general.Node[source.File]) error {
	return nil
}
