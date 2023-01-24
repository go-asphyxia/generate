package parser

import (
	"os"
	"path"

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

func (parser *Parser) Parse() (tree general.Tree[any], err error) {
	err = Read((*general.Node[any])(&tree), parser.configuration.Path)
	return
}

func Read(node *general.Node[any], name string) error {
	file, err := os.Open(name)
	if err != nil {
		return err
	}

	info, err := file.Stat()
	if err != nil {
		return err
	}

	if info.IsDir() {
		node.Data = source.FlieList{Name: info.Name()}

		fileList, err := os.ReadDir(name)
		if err != nil {
			return err
		}

		for i := range fileList {
			node := node.Children.PushTail(general.Node[any]{})

			err = Read(&node.Data, path.Join(name, fileList[i].Name()))
			if err != nil {
				return err
			}
		}
	} else {
		node.Data = source.File{Name: info.Name()}
	}

	return nil
}
