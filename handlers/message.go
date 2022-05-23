package handlers

import (
	"strings"

	"github.com/Scrip7/nestjs-discord-utility-bot/cache"
	"github.com/bwmarrin/discordgo"
)

const (
	COMMAND_PREFIX = "?tag "
)

// messageCreate will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	// TODO: only allow admins to use this command
	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
		return
	}

	// TODO: refactor this

	// get length of prefix
	prefixLen := len(COMMAND_PREFIX)
	if len(m.Content) < prefixLen {
		return
	}
	// extract prefix length from message content
	prefix := m.Content[:prefixLen]
	// convert to lowercase
	prefix = strings.ToLower(prefix)

	// if message starts with the prefix
	if prefix == COMMAND_PREFIX {
		message := strings.ToLower(m.Content)
		// remove the prefix from the message content
		message = strings.TrimPrefix(message, COMMAND_PREFIX)
		// get the first line of the message
		message = strings.Split(message, "\n")[0]
		// replace spaces with dash
		message = strings.Replace(message, " ", "-", -1)

		// check cache if the tag exists
		content, ok := cache.Driver.Get(message)
		if ok {
			s.ChannelMessageSend(m.ChannelID, content.(string))
		} else {
			s.ChannelMessageSend(m.ChannelID, "Tag not found")
		}
	}
}
