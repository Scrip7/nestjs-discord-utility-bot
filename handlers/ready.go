package handlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

func Ready(s *discordgo.Session, m *discordgo.Ready) {
	logrus.WithField("id", s.State.User.ID).Infof("Logged in as '%v#%v'", s.State.User.Username, s.State.User.Discriminator)

	if err := s.UpdateStatusComplex(discordgo.UpdateStatusData{
		Activities: []*discordgo.Activity{
			{
				Name: COMMAND_PREFIX,
				Type: discordgo.ActivityTypeListening,
			},
		},
	}); err != nil {
		logrus.Fatalf("error updating bot self status: %v", err)
	}
	// log updated status
	logrus.Info("Bot status updated")
}
