package source

import "github.com/go-asphyxia/core/bytes"

type (
	Point struct {
		Name string
		Data any
	}

	Folder struct {
		Path string
	}

	File struct {
		Text bytes.Buffer
	}

	Package struct {
		Name string
	}
)
