package main

import (
	md "github.com/nao1215/markdown"
	"strconv"
)

type Mixins struct {
	Title string
}

func (m *Mixins) Generate(cnf Config, markdown *md.Markdown) {
	requireJsConfig := RequireJsConfig{}
	areaMap := make(map[string][]byte)

	areaMap["base"] = []byte{}
	areaMap["frontend"] = []byte{}
	areaMap["adminhtml"] = []byte{}

	for area, _ := range areaMap {
		data := requireJsConfig.getRequireJsConfigContent(area, cnf.ModulePath)
		areaMap[area] = data

		mixins, err := requireJsConfig.ExtractMixins(data)
		if err != nil {
			// File likely does not have any mixins
			continue
		}

		var rows [][]string
		for _, mixin := range mixins {
			row := []string{mixin.Target, mixin.Mixin, strconv.FormatBool(mixin.Status)}
			rows = append(rows, row)
		}

		if m.Title == "" {
			m.Title = "Mixins"
			markdown.H2("Mixins")
		}

		markdown.H3(area)
		markdown.Table(
			md.TableSet{
				Header: []string{"Target", "Mixin", "Status"},
				Rows:   rows,
			},
		)
	}
}
