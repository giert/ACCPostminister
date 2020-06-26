package bot

import (
	"github.com/bwmarrin/discordgo"
)

var confirmationMessageID string

func confirm(s *discordgo.Session, channelID string, c string) {
	msg, _ := s.ChannelMessageSend(channelID, c)
	s.ChannelMessageDelete(channelID, confirmationMessageID)
	confirmationMessageID = msg.ID
}
