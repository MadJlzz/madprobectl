// Package parser provides a simple way to parse yaml configuration files.
package parser

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"reflect"
)

// Parent structure of the YAML file for HTTP probes.
type HttpProbe struct {
	Http []Http `yaml:"http"`
}

// Children structure representing a HTTP probe.
type Http struct {
	Name  string `yaml:"name"`
	URL   string `yaml:"url"`
	Delay uint   `yaml:"delay"`
}

// ParseYAML is used to read a .yml file and unmarshal it's content into the given structure.
func ParseYAML(src string, out interface{}) error {
	data, err := ioutil.ReadFile(src)
	if err != nil {
		return fmt.Errorf("error reading file: %s. %v\n", src, err)
	}
	err = yaml.Unmarshal(data, out)
	if err != nil {
		return fmt.Errorf("error unmarshaling yaml into: %s. %v\n", reflect.TypeOf(out), err)
	}
	return nil
}
