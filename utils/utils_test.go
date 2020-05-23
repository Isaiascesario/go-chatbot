package utils_test

import (
	"log"
	"github.com/isaiascesario/go-chatbot/utils"
	"testing"
)

func TestContainsEmpty(t *testing.T) {
	var emptySlice = []string{}
	var filledSlice = []string{"a", "b", "v"}
	log.Println(utils.ContainsAll(emptySlice, filledSlice))
}
