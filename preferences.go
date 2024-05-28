package main

import (
	md "github.com/nao1215/markdown"
)

type Preference struct {
	Preference []struct {
		Text string `xml:",chardata"`
		For  string `xml:"for,attr"`
		Type string `xml:"type,attr"`
	} `xml:"preference"`
}

func (p *Preference) Generate(cnf Config, markdown *md.Markdown) {
	var titleRendered bool = false
	xml := NewXml("di")
	areamap := xml.UnmarshalToMap(p, cnf)

	if len(areamap) == 0 {
		return
	}

	for area, preference := range areamap {
		preference := preference.(*Preference)
		var rows [][]string
		for _, p := range preference.Preference {
			row := []string{p.For, p.Type}
			rows = append(rows, row)
		}

		if len(preference.Preference) == 0 {
			continue
		}

		if !titleRendered {
			titleRendered = true
			markdown.H2("Preferences")
		}

		markdown.H3(area)
		markdown.Table(
			md.TableSet{
				Header: []string{"For", "Type"},
				Rows:   rows,
			},
		)
	}
}
