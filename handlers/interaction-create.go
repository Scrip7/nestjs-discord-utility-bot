package handlers

import (
	"github.com/Scrip7/nestjs-discord-utility-bot/cache"
	"github.com/bwmarrin/discordgo"
)

var (
	Commands = []*discordgo.ApplicationCommand{}
)

func InteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	name := i.ApplicationCommandData().Name
	// check cache if the tag exists
	content, ok := cache.Driver.Get(name)
	if ok {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: content.(string),
			},
		})
	} else {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Content not found.",
			},
		})
		s.ApplicationCommandDelete(s.State.User.ID, "", i.ID)
		s.ApplicationCommandDelete(s.State.User.ID, i.GuildID, i.ID)
	}
}
