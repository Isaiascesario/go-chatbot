package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"github.com/isaiascesario/go-chatbot/messenger"
	"github.com/isaiascesario/go-chatbot/messenger/discord"
)

func main() {
	var dis, err = messenger.Get(discord.Name)
	if err != nil {
		log.Fatal(err)
	}
	err = dis.StartMessageConsumer()
	defer dis.CloseConnection()
	if err != nil {
		log.Fatal(err)
	}
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

}
