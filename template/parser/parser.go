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

func (parser *Parser) Parse() (tree general.Node[source.Point], err error) {
	tree, err = Read(parser.configuration.Path)
	if err != nil {
		return
	}

	err = Scan(&tree)
	return
}

func Read(name string) (node general.Node[source.Point], err error) {
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
		text := bytes.Buffer{}

		_, err = text.ReadFrom(file)
		if err != nil {
			return
		}

		return general.Node[source.Point]{
			Data: source.Point{
				Name: infoName,
				Data: source.File{
					Text: text,
				},
			},
		}, nil
	}

	fileList, err := os.ReadDir(name)
	if err != nil {
		return node, err
	}

	node.Data = source.Point{
		Name: infoName,
		Data: source.Folder{
			Path: name,
		},
	}

	for i := range fileList {
		fileNode, err := Read(path.Join(name, fileList[i].Name()))
		if err != nil {
			return node, err
		}

		fileNode.PushTail(&node.Children)
	}

	return
}

func Scan(tree *general.Node[source.Point]) (err error) {
	// iterator := tree.Iterator()

	return
}
