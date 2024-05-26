package main

import (
	md "github.com/nao1215/markdown"
)

type Module struct {
	Module struct {
		Text     string `xml:",chardata"`
		Name     string `xml:"name,attr"`
		Sequence struct {
			Text   string `xml:",chardata"`
			Module []struct {
				Text string `xml:",chardata"`
				Name string `xml:"name,attr"`
			} `xml:"module"`
		} `xml:"sequence"`
	} `xml:"module"`
}

func (m *Module) Generate(cnf Config, markdown *md.Markdown) {
	var deps []string
	xml := NewXml("module")
	areamap := xml.UnmarshalToMap(m, cnf)
	module := areamap["global"].(*Module).Module

	markdown.H1f("%s", md.Bold(module.Name))
	markdown.PlainTextf("{{ %s }}", md.Italic("Provide a brief description of the module here"))
	markdown.PlainText("")

	for _, m := range module.Sequence.Module {
		deps = append(deps, m.Name)
	}

	if len(deps) > 0 {
		markdown.PlainTextf("%s", md.Bold("Dependencies:"))

		for _, dep := range deps {
			markdown.PlainTextf("%s", md.Code(dep))
		}

		return
	}

	markdown.PlainTextf("%s", md.Bold("Module has no stated dependencies"))
}
