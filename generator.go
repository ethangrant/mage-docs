package main

import (
	md "github.com/nao1215/markdown"
)

type MarkdownGenerator interface {
	Generate(cnf Config, markdown *md.Markdown)
}