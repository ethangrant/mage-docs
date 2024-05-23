package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	md "github.com/nao1215/markdown"
)

type Config struct {
	ModulePath string
}

func main() {
	const markdownTitle = "Generated by Mage-docs"

	var (
		path       string
		generators []MarkdownGenerator
	)

	flag.StringVar(&path, "module-path", "some path", "Path to the module location")
	flag.Parse()

	if !strings.HasSuffix(path, "/") {
		path = path + "/"
	}

	cnf := Config{
		ModulePath: path,
	}

	file, err := os.Create(cnf.ModulePath + "README.md")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	markdown := md.NewMarkdown(file).H1f("%s", md.BoldItalic(markdownTitle))

	generators = append(generators,
		new(Module),
		new(Routes),
		new(Webapi),
		new(Layout),
		new(Schema),
		new(Observer),
		new(Preference),
	)

	for _, generator := range generators {
		generator.Generate(cnf, markdown)
	}

	err = markdown.Build()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
