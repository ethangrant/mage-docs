package main

import (
	md "github.com/nao1215/markdown"
)

type Mixins struct {

}

func (m *Mixins) Generate(cnf Config, markdown *md.Markdown) {
	requireJsConfig := RequireJsConfig{}
	areaMap := make(map[string][]byte)

	areaMap["base"] = []byte{}
	areaMap["frontend"] = []byte{}
	areaMap["adminhtml"] = []byte{}

	markdown.H2("Mixins")

	for area, _ := range areaMap {
		data := requireJsConfig.getRequireJsConfigContent(area, cnf.ModulePath)
		areaMap[area] = data

		requireJsConfig.UnmarshalJSON(data)

		// requireJsConfig.Config
	}
}