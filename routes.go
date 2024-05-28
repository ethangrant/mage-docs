package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	md "github.com/nao1215/markdown"
)

type Controllers struct {
}

type Routes struct {
	Router struct {
		Text  string `xml:",chardata"`
		ID    string `xml:"id,attr"`
		Route struct {
			Text      string `xml:",chardata"`
			ID        string `xml:"id,attr"`
			FrontName string `xml:"frontName,attr"`
			Module    struct {
				Text string `xml:",chardata"`
				Name string `xml:"name,attr"`
			} `xml:"module"`
		} `xml:"route"`
	} `xml:"router"`
}

// @todo - ignore abstract classes
// @todo - feasible to get param information?
// @todo - check if we can see what kind of result is returned?
// @todo - handle other router types
// @todo - rework to output a table
func (r *Routes) Generate(cnf Config, markdown *md.Markdown) {
	areaMap := NewXml("routes").UnmarshalToMap(Routes{}, cnf)

	for area, route := range areaMap {
		var uris []string
		router := route.(*Routes).Router
		id := router.ID
		frontName := router.Route.FrontName

		// Only handle stadard and admin for now
		if id != "standard" && id != "admin" {
			continue
		}

		controllerDir := cnf.ModulePath + "Controller/"

		filepath.WalkDir(controllerDir, func(path string, d os.DirEntry, err error) error {

			if d.IsDir() {
				return nil
			}

			// Skip adminhtml for FE area
			if area == "frontend" && strings.Contains(path, "Adminhtml") {
				return nil
			}

			// just the path to the action relative to "Controller/", remove .php file extension also
			relativePath := strings.TrimSuffix(strings.TrimPrefix(path, controllerDir), ".php")

			if area == "adminhtml" {
				relativePath = strings.TrimPrefix(relativePath, "Adminhtml/")
			}

			uriElements := strings.Split(relativePath, "/")

			// get action then remove it from slice
			action := uriElements[len(uriElements)-1]
			uriElements = uriElements[:len(uriElements)-1]

			// files that sit directly in "Controller/" match this condition
			if len(uriElements) == 0 {
				return nil
			}

			controllerName := uriElements[0]
			if len(uriElements) > 1 {
				strings.Join(uriElements, "_")
			}

			uris = append(uris, strings.ToLower(fmt.Sprintf("%s/%s/%s", frontName, controllerName, action)))

			return nil
		})

		markdown.H3(area)
		for key, uri := range uris {
			if key == 0 {
				markdown.H2("Routes")
			}

			markdown.PlainTextf("%s", md.Code(uri))
		}
	}
}
