package main

import (
	md "github.com/nao1215/markdown"
)

type Preferences struct {
}

type Preference struct {
	Preference []struct {
		Text string `xml:",chardata"`
		For  string `xml:"for,attr"`
		Type string `xml:"type,attr"`
	} `xml:"preference"`
}

func (p *Preferences) Generate(cnf Config, markdown *md.Markdown) {
	var titleRendered bool = false
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
