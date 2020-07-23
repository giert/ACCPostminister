package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

func confirm(s *discordgo.Session, channelID string, c string) (*discordgo.Message, error) {
	msg, err := s.ChannelMessageSend(channelID, c)
	if err != nil {
		return nil, errors.Wrap(err, "while sending confirmation message")
	}

	if msgIDs.Confirmation != "" {
		err = s.ChannelMessageDelete(channelID, msgIDs.Confirmation)
		if err != nil {
			return nil, errors.Wrapf(err, "while deleting message %s", msgIDs.Confirmation)
		}
	}

	msgIDs.Confirmation = msg.ID
	return msg, nil
}
