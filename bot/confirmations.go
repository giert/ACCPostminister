package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

var confirmationMessageID string

func confirm(s *discordgo.Session, channelID string, c string) (*discordgo.Message, error) {
	msg, err := s.ChannelMessageSend(channelID, c)
	if err != nil {
		return nil, errors.Wrap(err, "while sending confirmation message")
	}

	if confirmationMessageID != "" {
		err = s.ChannelMessageDelete(channelID, confirmationMessageID)
		if err != nil {
			return nil, errors.Wrapf(err, "while deleting message %s", confirmationMessageID)
		}
	}

	confirmationMessageID = msg.ID
	return msg, nil
}
