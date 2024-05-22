package main

import (
	md "github.com/nao1215/markdown"
)

// @TODO -refactor similart XML types observer,preference etc

type Observer struct {
	Event []struct {
		Text     string `xml:",chardata"`
		Name     string `xml:"name,attr"`
		Observer []struct {
			Text     string `xml:",chardata"`
			Name     string `xml:"name,attr"`
			Instance string `xml:"instance,attr"`
		} `xml:"observer"`
	} `xml:"event"`
}

func (o *Observer) Generate(cnf Config, markdown *md.Markdown) {
	markdown.H2("Observers")
	xml := NewXml("events")
	areamap := xml.UnmarshalToMap(o, cnf)

	for area, observer := range areamap {
		observer := observer.(*Observer)
		var rows [][]string
		for _, event := range observer.Event {
			name := event.Name
			for _, obs := range event.Observer {
				row := []string{name, obs.Name, obs.Instance}
				rows = append(rows, row)
			}
		}

		if len(observer.Event) == 0 {
			continue
		}

		markdown.H3(area)
		markdown.Table(
			md.TableSet{
				Header: []string{"Event", "Observer name", "Instance"},
				Rows:   rows,
			},
		)
	}
}
