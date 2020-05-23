package disgo_test

import (
	"log"
	"github.com/isaiascesario/go-chatbot/utils"
	"testing"
)

func TestContainsAll(t *testing.T) {
	var slice1 = []string{"a", "b", "c"}
	var slice2 = []string{"a", "b", "c"}
	log.Printf("should be true: %v", utils.ContainsAll(slice1, slice2))
	slice1 = []string{"a", "b", "c"}
	slice2 = []string{"a", "c"}
	log.Printf("should be false: %v", utils.ContainsAll(slice1, slice2))
	slice1 = []string{"a", "d", "k"}
	slice2 = []string{"a", "c", "d", "f", "j", "k"}
	log.Printf("should be true: %v", utils.ContainsAll(slice1, slice2))
}

func TestNormalize(t *testing.T) {
	log.Println(utils.NormalizeString("Não há razão"))
}
