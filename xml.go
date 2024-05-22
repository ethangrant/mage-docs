package main

import (
	"fmt"
	"os"
	"io"
	"encoding/xml"
)

type XmlPath struct {
	Area string
	Path string
}

type Xml struct {
	Paths []XmlPath
}

func NewXml(file string) Xml {
	var x Xml
	x.Paths = []XmlPath{
		{"global", fmt.Sprintf("etc/%s.xml", file)},
		{"frontend", fmt.Sprintf("etc/frontend/%s.xml", file)},
		{"adminhtml", fmt.Sprintf("etc/adminhtml/%s.xml", file)},
	}

	return x
}

// pass a markdowngenerator to hydrate once a new instance of Xml has been created specifying files
// function will unmarshal the xml files and hydrate the passed type.
func (x *Xml) UnmarshalToMap(m MarkdownGenerator, cnf Config) (map[string]interface{}) {
	areaMap := make(map[string]interface{})

	for _, xmlPath := range x.Paths {
		area := xmlPath.Area
		path := xmlPath.Path

		file, err := os.Open(cnf.ModulePath + path)
		if err != nil {
			continue
		}
		defer file.Close()

		byts, err := io.ReadAll(file)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		err = xml.Unmarshal(byts, &m)
		if err != nil {
			fmt.Println("Unmarshal failing")
			continue
		}

		areaMap[area] = m
	}

	return areaMap
}