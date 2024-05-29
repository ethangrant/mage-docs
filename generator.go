package main

import (
	md "github.com/nao1215/markdown"
)

type MarkdownGenerator interface {
	Generate(cnf Config, markdown *md.Markdown)
}

func RenderTitle(currentTitle string, title string, markdown *md.Markdown) string {
	if currentTitle == "" {
		currentTitle = title
		markdown.H2(title)
	}

	return currentTitle
}

func RenderTable(header []string, rows [][]string, area string, markdown *md.Markdown) {
	if area != "" {
		markdown.H3(area)
	}

	markdown.Table(
		md.TableSet{
			Header: []string{"Target", "Mixin", "Status"},
			Rows:   rows,
		},
	)
}