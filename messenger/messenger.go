package messenger

import (
	"temp/disgo/messenger/core"
	"temp/disgo/messenger/discord"
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
