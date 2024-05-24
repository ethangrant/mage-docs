package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type RequireJsConfig struct {
	Config struct {
		Mixins []Mixin
	}
}

type Mixin struct {
	Target string
	Mixin string
}

func (r *RequireJsConfig) getRequireJsConfigContent(area string, modulePath string) ([]byte)  {
	var data []byte

	data, err := os.ReadFile(fmt.Sprintf("%sview/%s/requirejs-config.js", modulePath, area))
	if err != nil {
		fmt.Println(err.Error())
		return []byte{}
	}

	re := regexp.MustCompile(`([a-zA-Z]+:)`)
	reDocBlock := regexp.MustCompile(`\/\*\*(.|[\r\n])*?\*\/`)

	// these are JS files we are stripping them down to JSON
	dataAsString := string(data)
	dataAsString = strings.ReplaceAll(dataAsString, "'", "\"")
	dataAsString = strings.ReplaceAll(dataAsString, ";", "")
	dataAsString = strings.ReplaceAll(dataAsString, "var config = ", "")
	dataAsString = re.ReplaceAllString(dataAsString, "\"$1\":")
	dataAsString = reDocBlock.ReplaceAllString(dataAsString, "")
	dataAsString = strings.TrimSpace(dataAsString)


	fmt.Println(dataAsString)

	return data
}

func (r *RequireJsConfig) UnmarshalJSON(data []byte) error {
	var v []interface{}

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}


	fmt.Println(v[0])

	return nil
}