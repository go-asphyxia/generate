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

func Constructor(configuration *Configuration) (parser *Parser) {
	return &Parser{
		configuration: *configuration,
	}
}

func (parser *Parser) Parse() (tree general.Node[source.File], err error) {
	return Read(parser.configuration.Path)
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

	infoName := info.Name()

	if info.IsDir() == false {
		buffer := bytes.Buffer{}

		_, err = buffer.ReadFrom(file)
		if err != nil {
			return
		}

		return general.Node[source.File]{
			Data: source.File{
				Name: infoName,
				Data: source.FileData{
					Text: string(buffer),
				},
			},
		}, nil
	}

	infoList, err := os.ReadDir(name)
	if err != nil {
		return node, err
	}

	node.Data = source.File{
		Name: infoName,
	}

	for i := range infoList {
		infoNode, err := Read(path.Join(name, infoList[i].Name()))
		if err != nil {
			return node, err
		}

		infoNode.PushTail(&node.Children)
	}

	return
}
