package handlers

import (
	"fmt"

	"github.com/Scrip7/nestjs-discord-utility-bot/commands"
	"github.com/bwmarrin/discordgo"
	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
)

func Ready(s *discordgo.Session, m *discordgo.Ready) {
	logrus.WithFields(logrus.Fields{
		"id":       s.State.User.ID,
		"username": fmt.Sprintf("%s#%s", s.State.User.Username, s.State.User.Discriminator),
	}).Infof("Logged in as")

	registeredCommands, err := s.ApplicationCommands(s.State.User.ID, "")
	if err != nil {
		logrus.Fatalf("Failed to fetch registered app commands: %v", err)
	}

	// anonymous function to avoid code duplication
	// in for loops below
	isCMDRegistered := func(commandName string) bool {
		if len(registeredCommands) == 0 {
			return false
		}

		predicate := func(x *discordgo.ApplicationCommand) bool {
			return commandName == x.Name
		}

		return lo.ContainsBy(registeredCommands, predicate)
	}

	for _, v := range commands.Commands {
		// ignore all commands that are not registered by the bot
		if !isCMDRegistered(v.Name) {
			continue
		}

		logrus.WithField("cmd", v.Name).Warn("Deleting slash command")
		err := s.ApplicationCommandDelete(s.State.User.ID, "", v.ID)
		if err != nil {
			logrus.Fatalf("Failed to delete the '%v' slash command: %v", v.Name, err)
		}
	}

	for _, v := range commands.Commands {
		// ignore all commands that are registered by the bot
		if isCMDRegistered(v.Name) {
			continue
		}

		logrus.WithField("cmd", v.Name).Warn("Creating slash command")
		_, err := s.ApplicationCommandCreate(s.State.User.ID, "", v)
		if err != nil {
			logrus.Errorf("Failed to create the '%v' slash command: %v", v.Name, err)
			return
		}
	}
	logrus.Info("Done processing slash commands")

	err = s.UpdateStatusComplex(discordgo.UpdateStatusData{
		Activities: []*discordgo.Activity{
			{
				Name: "slash commands.",
				Type: discordgo.ActivityTypeListening,
			},
		},
	})
	if err != nil {
		logrus.Fatalf("error updating bot self status: %v", err)
	}
	logrus.Info("Bot status updated")

	logrus.Info("Bot is ready")
}
