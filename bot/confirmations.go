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

	if msgIDs.confirmation != "" {
		err = s.ChannelMessageDelete(channelID, msgIDs.confirmation)
		if err != nil {
			return nil, errors.Wrapf(err, "while deleting message %s", msgIDs.confirmation)
		}
	}

	msgIDs.confirmation = msg.ID
	return msg, nil
}
