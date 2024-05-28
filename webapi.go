package main

import (
	md "github.com/nao1215/markdown"
)

type Webapi struct {
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
	xml := NewXml("webapi")
	areamap := xml.UnmarshalToMap(w, cnf)

	if len(areamap) == 0 {
		return
	}

	markdown.H2("API Routes")

	for area, webapi := range areamap {
		webapi := webapi.(*Webapi)
		var rows [][]string

		for _, route := range webapi.Route {
			row := []string{route.URL, route.Method}
			rows = append(rows, row)
		}

		markdown.H3(area)
		markdown.Table(
			md.TableSet{
				Header: []string{"Url", "Method"},
				Rows:   rows,
			},
		)
	}
}
