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
	markdown.H2("Plugins (Interceptors)")
	xml := NewXml("di")
	areamap := xml.UnmarshalToMap(p, cnf)

	for area, plugin := range areamap {
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

		markdown.H3(area)
		markdown.Table(
			md.TableSet{
				Header: []string{"Target", "Type", "Name"},
				Rows:   rows,
			},
		)
	}
}
