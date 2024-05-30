package main

import (
	md "github.com/nao1215/markdown"
)

type Preferences struct {
	Title string
}

type Preference struct {
	Preference []struct {
		Text string `xml:",chardata"`
		For  string `xml:"for,attr"`
		Type string `xml:"type,attr"`
	} `xml:"preference"`
}

func (p *Preferences) Generate(cnf Config, markdown *md.Markdown) {
	areaMap := NewXml("di").UnmarshalToMap(Preference{}, cnf)

	if len(areaMap) == 0 {
		return
	}

	for area, preference := range areaMap {
		preference := preference.(*Preference)
		var rows [][]string
		for _, p := range preference.Preference {
			row := []string{md.Code(p.For), md.Code(p.Type)}
			rows = append(rows, row)
		}

		if len(preference.Preference) == 0 {
			continue
		}

		p.Title = RenderTitle(p.Title, "Preferences", markdown)
		RenderTable([]string{"For", "Type"}, rows, area, markdown)
	}
}
