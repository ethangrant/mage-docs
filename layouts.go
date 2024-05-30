package main

import (
	"io/fs"
	"path/filepath"
	"strings"

	md "github.com/nao1215/markdown"
)

type Layouts struct {
	Title string
}

func (l *Layouts) Generate(cnf Config, markdown *md.Markdown) {
	var layouts []string
	markdown.PlainText("This module interacts with the following layout handles in frontend")

	layoutPath := cnf.ModulePath + "view/frontend/layout"
	filepath.WalkDir(layoutPath, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		path = filepath.Base(path)
		layouts = append(layouts, md.Code(strings.TrimSuffix(path, ".xml")))

		return nil
	})

	if len(layouts) == 0 {
		return
	}

	l.Title = RenderTitle(l.Title, "Layouts", markdown)

	for _, layout := range layouts {
		markdown.PlainTextf("%s", layout)
	}
}
