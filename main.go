package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Scrip7/nestjs-discord-utility-bot/cache"
	"github.com/Scrip7/nestjs-discord-utility-bot/core"
	"github.com/Scrip7/nestjs-discord-utility-bot/handlers"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

func main() {
	core.InitLogger()
	core.LoadConfig()

	cache.Init()
	cache.BootstrapContent()

	// Create a new Discord session
	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
	if err != nil {
		logrus.Fatalf("error creating Discord session: %v", err)
	}

	// Discord event handlers
	dg.AddHandler(handlers.MessageCreate)

	// we only care about receiving message events
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		logrus.Fatalf("error opening connection: %v", err)
	}

	user, err := dg.User("@me")
	if err != nil {
		logrus.Fatalf("error getting @me user: %v", err)
	}
	logrus.WithField("id", user.ID).Info("Bot is running")

	if err = dg.UpdateStatusComplex(discordgo.UpdateStatusData{
		Activities: []*discordgo.Activity{
			{
				Name: handlers.COMMAND_PREFIX,
				Type: discordgo.ActivityTypeListening,
			},
		},
	}); err != nil {
		logrus.Fatalf("error updating status: %v", err)
	}
	// log updated status
	logrus.Info("Bot status updated")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// cleanly close down the Discord session.
	dg.Close()
}
