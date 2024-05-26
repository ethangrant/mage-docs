package main

import (
	"io/fs"
	"path/filepath"
	"strings"

	md "github.com/nao1215/markdown"
)

type Layout struct {
}

func (l *Layout) Generate(cnf Config, markdown *md.Markdown) {
	markdown.H2("Layouts")
	markdown.PlainText("This module interacts with the following layout handles in frontend")

	layoutPath := cnf.ModulePath + "view/frontend/layout"
	filepath.WalkDir(layoutPath, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		path = filepath.Base(path)
		markdown.PlainTextf("%s", md.Code(strings.TrimSuffix(path, ".xml")))

		return nil
	})
}
