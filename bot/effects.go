package bot

import (
	"ACCPostminister/language"

	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

var effects = map[string]func(*discordgo.Session, *discordgo.Message) error{
	language.BotChannelCommand: setBotChannel,
	language.CleanseCommand:    cleansing,
	language.RoleCommand:       roleReactions,
}

func setBotChannel(s *discordgo.Session, m *discordgo.Message) error {
	globalIDs.Botchannel = m.ChannelID
	return nil
}

func cleansing(s *discordgo.Session, m *discordgo.Message) error {
	if globalIDs.Botchannel == "" {
		return nil
	}

	messages, err := s.ChannelMessages(globalIDs.Botchannel, 0, "", "", "")
	if err != nil {
		return errors.Wrap(err, "while getting messages from bot channel")
	}

	for _, message := range messages {
		if !globalIDs.contains(message.ID) {
			err = s.ChannelMessageDelete(globalIDs.Botchannel, message.ID)
			if err != nil {
				return errors.Wrap(err, "while deleting messages from bot channel")
			}
		}
	}

	return nil
}

func roleReactions(s *discordgo.Session, m *discordgo.Message) error {
	if globalIDs.Role == "" {
		globalIDs.Role = m.ID

		for _, role := range roles {
			err := s.MessageReactionAdd(m.ChannelID, globalIDs.Role, role.Emoji)
			if err != nil {
				return errors.Wrapf(err, "while adding %s to role message", role.Emoji)
			}
		}
	}
	return nil
}
