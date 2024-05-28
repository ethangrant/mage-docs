package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"reflect"
)

type XmlPath struct {
	Area string
	Path string
}

type Xml struct {
	Paths []XmlPath
}

func NewXml(file string, area ...bool) Xml {
	var x Xml
	var useArea bool = true

	if len(area) > 0 {
		useArea = area[0]
	}

	if !useArea {
		x.Paths = []XmlPath{{"global", fmt.Sprintf("etc/%s.xml", file)}}
		return x
	}

	x.Paths = []XmlPath{
		{"global", fmt.Sprintf("etc/%s.xml", file)},
		{"frontend", fmt.Sprintf("etc/frontend/%s.xml", file)},
		{"adminhtml", fmt.Sprintf("etc/adminhtml/%s.xml", file)},
		{"webapi_rest", fmt.Sprintf("etc/webapi_rest/%s.xml", file)},
		{"webapi_soap", fmt.Sprintf("etc/webapi_soap/%s.xml", file)},
	}

	return x
}

// Takes any struct type and creates a map of them with unmarshalled XML
func (x Xml) UnmarshalToMap(dataStructure any, cnf Config) map[string]any {
	areaMap := make(map[string]any)

	for _, xmlPath := range x.Paths {
		file, err := os.Open(cnf.ModulePath + xmlPath.Path)
		if err != nil {
			continue
		}
		defer file.Close()

		byts, err := io.ReadAll(file)
		if err != nil {
			continue
		}

		t := reflect.TypeOf(dataStructure)
		if t.Kind() != reflect.Struct {
			continue
		}

		v := reflect.New(t).Interface()
		err = xml.Unmarshal(byts, &v)
		if err != nil {
			continue
		}

		areaMap[xmlPath.Area] = v
	}

	return areaMap
}
