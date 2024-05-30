package main

import (
	md "github.com/nao1215/markdown"
)

type Webapi struct {
	Title string
	Route []struct {
		Text    string `xml:",chardata"`
		URL     string `xml:"url,attr"`
		Method  string `xml:"method,attr"`
		Service struct {
			Text   string `xml:",chardata"`
			Class  string `xml:"class,attr"`
			Method string `xml:"method,attr"`
		} `xml:"service"`
		Resources struct {
			Text     string `xml:",chardata"`
			Resource struct {
				Text string `xml:",chardata"`
				Ref  string `xml:"ref,attr"`
			} `xml:"resource"`
		} `xml:"resources"`
	} `xml:"route"`
}

func (w *Webapi) Generate(cnf Config, markdown *md.Markdown) {
	areaMap := NewXml("webapi").UnmarshalToMap(Webapi{}, cnf)

	if len(areaMap) == 0 {
		return
	}

	w.Title = RenderTitle(w.Title, "API Routes", markdown)

	for area, webapi := range areaMap {
		webapi := webapi.(*Webapi)
		var rows [][]string

		for _, route := range webapi.Route {
			row := []string{md.Code(route.URL), route.Method}
			rows = append(rows, row)
		}

		RenderTable([]string{"Url", "Method"}, rows, area, markdown)
	}
}
