package parser

import (
	"os"
	"path"

	"github.com/go-asphyxia/core/bytes"
	"github.com/go-asphyxia/data/tree/general"
	"github.com/go-asphyxia/generate/template/source"
)

type (
	Configuration struct {
		Path     string `json:"Path"`
		Selector string `json:"Selector"`
	}

	Parser struct {
		configuration Configuration
	}
)

func Constructor(configuration *Configuration) *Parser {
	return &Parser{
		configuration: *configuration,
	}
}

func (parser *Parser) Parse() (tree general.Tree[source.File], err error) {
	tree.Node, err = Read(parser.configuration.Path)
	return
}

func Read(name string) (node general.Node[source.File], err error) {
	file, err := os.Open(name)
	if err != nil {
		return
	}

	info, err := file.Stat()
	if err != nil {
		return
	}

	node.Data.Name = info.Name()

	if info.IsDir() {
		fileList, err := os.ReadDir(name)
		if err != nil {
			return node, err
		}

		for i := range fileList {
			inner, err := Read(path.Join(name, fileList[i].Name()))
			if err != nil {
				return node, err
			}

			node.Children.PushTail(inner)
		}
	} else {
		buffer := bytes.Buffer{}

		_, err = buffer.ReadFrom(file)
		if err != nil {
			return
		}

		node.Data.Data = source.FileData{
			Text: string(buffer),
		}
	}

	return
}
