package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Scrip7/nestjs-discord-utility-bot/cache"
	"github.com/Scrip7/nestjs-discord-utility-bot/core"
	"github.com/Scrip7/nestjs-discord-utility-bot/handlers"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	core.InitLogger()
	_ = godotenv.Load()

	cache.Init()
	cache.BootstrapContent()

	// Create a new Discord session
	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
	if err != nil {
		logrus.Fatalf("error creating Discord session: %v", err)
	}

	// Discord event handlers
	dg.AddHandler(handlers.Ready)
	dg.AddHandler(handlers.MessageCreate)
	dg.AddHandler(handlers.InteractionCreate)

	// we only care about receiving message events
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		logrus.Fatalf("error opening connection: %v", err)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGTERM)
	<-sc

	// cleanly close down the Discord session.
	dg.Close()
}
