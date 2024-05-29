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
			row := []string{md.Code(mixin.Target), md.Code(mixin.Mixin), strconv.FormatBool(mixin.Status)}
			rows = append(rows, row)
		}

		m.Title = RenderTitle(m.Title, "Mixins", markdown)
		RenderTable([]string{"Target", "Mixin", "Status"}, rows, area, markdown)
	}
}
