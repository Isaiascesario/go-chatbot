package yaml_test

import (
	"fmt"
	"io/ioutil"
	"log"
	discy "temp/disgo/yaml"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestGetYAML(t *testing.T) {

	var data, err = ioutil.ReadFile("./default.yaml")
	if err != nil {
		log.Fatalf("cannot read data: %v", err)
	}
	// toInterface(data)
	toType(data)
}

func toInterface(data []byte) {
	var yamlStruct interface{}
	var err = yaml.Unmarshal([]byte(data), &yamlStruct)
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
	}
	fmt.Println(yamlStruct)
}

func toType(data []byte) {
	var yamlStruct map[string]discy.MessageList
	var err = yaml.Unmarshal([]byte(data), &yamlStruct)
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
	}
	fmt.Println(yamlStruct)
}
