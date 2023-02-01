package main

import (
	"log"
	"os"

	"github.com/go-asphyxia/generate/template/parser"
	"github.com/go-asphyxia/generate/template/processor"
	"github.com/go-asphyxia/generate/template/source"

	"github.com/goccy/go-json"
)

type (
	Configuration struct {
		Parser    parser.Configuration    `json:"Parser"`
		Processor processor.Configuration `json:"Processor"`
	}
)

func main() {
	configuration := &Configuration{
		Parser: parser.Configuration{
			Path:     ".",
			Selector: "go.generate",
		},
		Processor: processor.Configuration{
			Selector: "go",
		},
	}

	if len(os.Args) == 2 {
		err := json.Unmarshal([]byte(os.Args[1]), configuration)
		if err != nil {
			log.Println(err)
		}
	}

	Main(configuration)
}

func Main(configuration *Configuration) {
	tree, err := parser.Constructor(&configuration.Parser).Parse()
	if err != nil {
		log.Panic(err)
	}

	println(tree.String(debug))

	err = processor.Constructor(&configuration.Processor).Process(&tree)
	if err != nil {
		log.Panic(err)
	}
}

func debug(data source.Point) string {
	switch data.Data.(type) {
	case source.Folder:
		return `"` + data.Name + `"`
	case source.File:
		return `"` + data.Name + `"`
	default:
		return `unknown`
	}
}
