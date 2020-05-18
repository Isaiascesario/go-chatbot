package dao

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var apiStr = "https://uselessfacts.jsph.pl"

// GetRandomFact !
func GetRandomFact() string {
	var resp, err = http.Get(fmt.Sprintf("%s/random.json?language=en", apiStr))
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}

	var factAc fact
	err = json.Unmarshal(contents, &factAc)
	if err != nil {
		fmt.Printf("aaaaaaa %s", err)
	}
	return factAc.Text
}

type fact struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}
