package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

func confirm(s *discordgo.Session, channelID string, message string) (*discordgo.Message, error) {
	msg, err := s.ChannelMessageSend(channelID, message)
	if err != nil {
		return nil, errors.Wrap(err, "while sending confirmation message")
	}

	if validMessageID(s, channelID, persistent.Confirmation) {
		err = s.ChannelMessageDelete(channelID, persistent.Confirmation)
		if err != nil {
			return nil, errors.Wrapf(err, "while deleting message %s", persistent.Confirmation)
		}
	}

	persistent.Confirmation = msg.ID
	return msg, nil
}
