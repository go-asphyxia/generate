package main

import (
	"os"

	"github.com/go-asphyxia/generate/template/parser"
	"github.com/go-asphyxia/generate/template/processor"

	"github.com/goccy/go-json"
)

type (
	Configuration struct {
		Parser    parser.Configuration    `json:"Parser"`
		Processor processor.Configuration `json:"Processor"`
	}
)

func main() {
	configuration := &Configuration{}

	if len(os.Args) == 2 {
		err := json.Unmarshal([]byte(os.Args[1]), configuration)
		if err != nil {
			panic(err)
		}
	} else {
		configuration.Parser = parser.Configuration{
			Path:     ".",
			Selector: "go.generate",
		}

		configuration.Processor = processor.Configuration{
			Selector: "go",
		}
	}

	Main(configuration)
}

func Main(configuration *Configuration) {
	tree, err := parser.Constructor(&configuration.Parser).Parse()
	if err != nil {
		panic(err)
	}

	err = processor.Constructor(&configuration.Processor).Process(tree)
	if err != nil {
		panic(err)
	}
}
