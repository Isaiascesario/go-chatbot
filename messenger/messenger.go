package messenger

import (
	"github.com/isaiascesario/go-chatbot/messenger/core"
	"github.com/isaiascesario/go-chatbot/messenger/discord"
)

func init() {
	discord.Register()
}

// Get !
var Get = core.Get

// DefaultMessageHandler !
func DefaultMessageHandler(message string) (string, error) {
	var responseMessage string
	return responseMessage, nil
}
