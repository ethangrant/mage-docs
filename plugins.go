package main

import (
	md "github.com/nao1215/markdown"
)

type Plugins struct {
	Title string
	Type  []struct {
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
				row := []string{md.Code(target), md.Code(pluginNode.Type), md.Code(pluginNode.Name)}
				rows = append(rows, row)
			}
		}

		p.Title = RenderTitle(p.Title, "Plugins (Interceptors)", markdown)
		RenderTable([]string{"Target", "Type", "Name"}, rows, area, markdown)
	}
}
