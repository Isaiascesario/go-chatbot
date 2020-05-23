package messagehandler

import (
	"fmt"
	"strings"

	"github.com/isaiascesario/go-chatbot/dao"
)

func init() {

}

// Default !
func Default(msg []string) (string, error) {
	var argument = strings.Join(msg[1:], " ")
	switch msg[0] {
	case "!c":
		var id, err = dao.GetCreature(argument)
		if err != nil {
			return "", fmt.Errorf("Error GetCreature: %s", err)
		}
		if id == 0 {
			return fmt.Sprintf(`Me parece que "%s" não é uma criatura, se aprume seu lixo`, argument), nil
		}
		return fmt.Sprintf("http://pt.wowhead.com/npc=%v", id), nil
	case "!i":
		var id, err = dao.GetItem(argument)
		if err != nil {
			return "", fmt.Errorf("Error GetItem: %s", err)
		}
		if id == 0 {
			return fmt.Sprintf(`Me parece que "%s" não é um item, seu merdinha corrige ai`, argument), nil
		}
		return fmt.Sprintf("http://pt.wowhead.com/item=%v", id), nil
	}
	return "", nil
}
