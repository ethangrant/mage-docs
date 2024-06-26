package main

import (
	md "github.com/nao1215/markdown"
)

type Observers struct {
	Title string
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

func (o *Observers) Generate(cnf Config, markdown *md.Markdown) {
	areaMap := NewXml("events").UnmarshalToMap(Observers{}, cnf)

	if len(areaMap) == 0 {
		return
	}

	for area, observer := range areaMap {
		observer := observer.(*Observers)
		var rows [][]string
		for _, event := range observer.Event {
			name := event.Name
			for _, obs := range event.Observer {
				row := []string{md.Code(name), md.Code(obs.Name), md.Code(obs.Instance)}
				rows = append(rows, row)
			}
		}

		if len(observer.Event) == 0 {
			continue
		}

		o.Title = RenderTitle(o.Title, "Observers", markdown)
		RenderTable([]string{"Event", "Observer name", "Instance"}, rows, area, markdown)
	}
}
