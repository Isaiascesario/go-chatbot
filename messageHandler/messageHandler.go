package messagehandler

import (
	"log"
	"math/rand"
	"temp/disgo/utils"
	"temp/disgo/yaml"
)

var defaultContent map[string]yaml.MessageList

func init() {
	var err error
	defaultContent, err = yaml.GetDefaultContent()
	if err != nil {
		log.Fatalf("Could not get default content: %s", err)
	}
}

// Default !
func Default(msg []string) string {
	for _, v := range defaultContent {
		for _, lista := range v.ContainsList {
			if utils.ContainsAll(msg, lista) {
				return v.Responses[rand.Intn(len(v.Responses))]
			}
		}
	}
	return defaultContent["no_aswer"].Responses[rand.Intn(len(defaultContent["no_aswer"].Responses))]
}
