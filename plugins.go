package main

import (
	md "github.com/nao1215/markdown"
)

type Plugins struct {
	Type []struct {
		Text   string `xml:",chardata"`
		Name   string `xml:"name,attr"`
		Shared string `xml:"shared,attr"`
		Plugin []struct {
			Text string `xml:",chardata"`
			Name string `xml:"name,attr"`
			Type string `xml:"type,attr"`
		} `xml:"plugin"`
	} `xml:"type"`
}

func (p *Plugins) Generate(cnf Config, markdown *md.Markdown) {
	var titleRendered bool = false
	areaMap := NewXml("di").UnmarshalToMap(Plugins{}, cnf)

	if len(areaMap) == 0 {
		return
	}

	for area, plugin := range areaMap {
		var rows [][]string
		plugin := plugin.(*Plugins)
		types := plugin.Type

		for _, typeNode := range types {
			target := typeNode.Name
			pluginNodes := typeNode.Plugin

			if len(pluginNodes) == 0 {
				continue
			}

			for _, pluginNode := range pluginNodes {
				row := []string{target, pluginNode.Type, pluginNode.Name}
				rows = append(rows, row)
			}
		}

		if !titleRendered {
			titleRendered = true
			markdown.H2("Plugins (Interceptors)")
		}

		markdown.H3(area)
		markdown.Table(
			md.TableSet{
				Header: []string{"Target", "Type", "Name"},
				Rows:   rows,
			},
		)
	}
}
