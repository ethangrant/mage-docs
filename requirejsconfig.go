package main

import (
	"encoding/json"
	"errors"
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
	Mixin  string
	Status bool
}

// requirejs-config.js files are javascript files with a single object defined
// By stripping down and modifying the file we can turn this into valid JSON.
func (r *RequireJsConfig) getRequireJsConfigContent(area string, modulePath string) []byte {
	var data []byte

	data, err := os.ReadFile(fmt.Sprintf("%sview/%s/requirejs-config.js", modulePath, area))
	if err != nil {
		fmt.Println(err.Error())
		return []byte{}
	}

	re := regexp.MustCompile(`([a-zA-Z]+):`)
	reDocBlock := regexp.MustCompile(`\/\*\*(.|[\r\n])*?\*\/`)

	dataAsString := string(data)
	dataAsString = strings.ReplaceAll(dataAsString, "'", "\"")
	dataAsString = strings.ReplaceAll(dataAsString, ";", "")
	dataAsString = strings.ReplaceAll(dataAsString, "var config = ", "")
	dataAsString = re.ReplaceAllString(dataAsString, "\"$1\":")
	dataAsString = reDocBlock.ReplaceAllString(dataAsString, "")
	dataAsString = strings.TrimSpace(dataAsString)

	return []byte(dataAsString)
}

// Once we have requirejs-config.js as a byte slice of JSON we can unmarshal
// and extract the mixin data we want.
func (r *RequireJsConfig) ExtractMixins(data []byte) ([]Mixin, error) {
	var mixinSlice []Mixin
	var result interface{}

	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	dataMap, ok := result.(map[string]interface{})
	if !ok {
		return nil, errors.New("problem with type assertion")
	}

	config, ok := dataMap["config"].(map[string]interface{})
	if !ok {
		return nil, errors.New("problem with type assertion could not find config")
	}

	mixins, ok := config["mixins"].(map[string]interface{})
	if !ok {
		return nil, errors.New("problem with type assertion could not find mixin")
	}

	for key, value := range mixins {
		var m Mixin

		mixinMap, ok := value.(map[string]interface{})
		if !ok {
			return nil, errors.New("problem with type assertion")
		}

		for key, value := range mixinMap {
			m.Status = value.(bool)
			m.Mixin = key
		}

		m.Target = key

		mixinSlice = append(mixinSlice, m)
	}

	return mixinSlice, nil
}
