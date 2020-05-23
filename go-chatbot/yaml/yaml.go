package yaml

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// Default !
type Default struct {
	Greetings   MessageList `yaml:"greetings"`
	RandomFacts MessageList `yaml:"random_facts"`
	NoAnswer    MessageList `yaml:"no_aswer"`
}

// MessageList !
type MessageList struct {
	ContainsList [][]string `yaml:"contains_list"`
	Responses    []string   `yaml:"responses"`
}

const defaultYAML = "yaml/default.yaml"

// GetDefaultContent !
func GetDefaultContent() (map[string]MessageList, error) {
	var data, err = ioutil.ReadFile(defaultYAML)
	if err != nil {
		return nil, fmt.Errorf("Could not read data: %v", err)
	}
	var defaultContent = make(map[string]MessageList)
	err = yaml.Unmarshal([]byte(data), &defaultContent)
	if err != nil {
		return nil, fmt.Errorf("Could not unmarshal data: %v", err)
	}
	return defaultContent, nil
}
