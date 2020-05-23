package discord

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	messagehandler "github.com/isaiascesario/go-chatbot/messageHandler"
	"github.com/isaiascesario/go-chatbot/messenger/core"
	"github.com/isaiascesario/go-chatbot/utils"

	"github.com/bwmarrin/discordgo"
)

// Name !
const Name core.Name = "discord"

var (
	token string
	me    string
)

// Register !
func Register() {
	core.Register(Name, &impl{})
}

type impl struct {
}

var dg *discordgo.Session

func (i *impl) StartMessageConsumer() error {

	var exists bool
	if token, exists = os.LookupEnv("DISCORD_BOT_TOKEN"); !exists {
		return fmt.Errorf("Could not get DISCORD_BOT_TOKEN env")
	}
	var err error
	dg, err = discordgo.New("Bot " + token)
	if err != nil {
		return fmt.Errorf("error creating session: %s", err)
	}

	dg.AddHandler(messageCreate)

	err = dg.Open()
	if err != nil {
		return fmt.Errorf("error opening connection: %s", err)
	}
	me = dg.State.User.ID
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	return nil
}

func (i *impl) CloseConnection() error {
	return dg.Close()
}
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	var msgHandled = utils.NormalizeString(m.Content)
	var msgSlice = strings.Split(msgHandled, " ")
	if m.Author.ID == s.State.User.ID || (len(msgSlice) < 2 && (msgSlice[0] != "!c" || msgSlice[0] != "!i")) {
		return
	}
	var response, err = messagehandler.Default(msgSlice)
	if err != nil {
		log.Printf("error on messagehandler.Default %s", err)
	}
	log.Printf(`Responding "%s" to %s's message "%s" `, response, m.Author, m.Content)
	s.ChannelMessageSend(m.ChannelID, response)
}
