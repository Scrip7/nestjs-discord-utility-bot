package handlers

import (
	"github.com/Scrip7/nestjs-discord-utility-bot/commands"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

func Ready(s *discordgo.Session, m *discordgo.Ready) {
	logrus.WithField("id", s.State.User.ID).Infof("Logged in as '%v#%v'", s.State.User.Username, s.State.User.Discriminator)

	globalRegisteredCommands, err := s.ApplicationCommands(s.State.User.ID, "")
	if err != nil {
		logrus.Fatalf("Could not fetch registered commands: %v", err)
	}

	// TODO: fix "Max number of daily application command creates has been reached (200)" problem
	for _, v := range globalRegisteredCommands {
		logrus.Warnf("Removing registered slash command: %v", v.Name)
		err := s.ApplicationCommandDelete(s.State.User.ID, "", v.ID)
		if err != nil {
			logrus.Fatalf("Cannot delete '%v' command: %v", v.Name, err)
		}
	}

	// TODO: register only non-registered commands
	for _, v := range commands.Commands {
		logrus.Warnf("Registering command: %s", v.Name)
		_, err := s.ApplicationCommandCreate(s.State.User.ID, "", v)
		if err != nil {
			logrus.Errorf("Cannot create '%v' command: %v", v.Name, err)
			return
		}
	}
	logrus.Info("Slash commands registered.")

	if err := s.UpdateStatusComplex(discordgo.UpdateStatusData{
		Activities: []*discordgo.Activity{
			{
				Name: "slash commands!",
				Type: discordgo.ActivityTypeListening,
			},
		},
	}); err != nil {
		logrus.Fatalf("error updating bot self status: %v", err)
	}
	logrus.Info("Bot status updated")

	logrus.Info("Bot is ready")
}
