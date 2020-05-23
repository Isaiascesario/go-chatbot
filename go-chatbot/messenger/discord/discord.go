package discord

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	messagehandler "temp/disgo/messageHandler"
	"temp/disgo/messenger/core"
	"temp/disgo/utils"
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
	if m.Author.ID == s.State.User.ID || !mentionedMe(m.Mentions) {
		return
	}
	var msgHandled = utils.NormalizeString(strings.ReplaceAll(m.Content, fmt.Sprintf("<@!%s>", me), ""))
	var msgSlice = strings.Split(msgHandled, " ")
	var response = messagehandler.Default(msgSlice)
	log.Printf(`Responding "%s" to %s's message "%s" `, response, m.Author, m.Content)
	s.ChannelMessageSend(m.ChannelID, response)
}

func mentionedMe(mentions []*discordgo.User) bool {
	for _, v := range mentions {
		if v.ID == me {
			return true
		}
	}
	return false
}
