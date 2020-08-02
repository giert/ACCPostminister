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

	if validMessageID(s, channelID, globalIDs.Confirmation) {
		err = s.ChannelMessageDelete(channelID, globalIDs.Confirmation)
		if err != nil {
			return nil, errors.Wrapf(err, "while deleting message %s", globalIDs.Confirmation)
		}
	}

	globalIDs.Confirmation = msg.ID
	return msg, nil
}
