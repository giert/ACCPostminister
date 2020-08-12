package bot

import (
	"ACCPostminister/language"

	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

var effects = map[string]func(*discordgo.Session, *discordgo.Message) error{}

func initEffects() {
	effects = map[string]func(*discordgo.Session, *discordgo.Message) error{
		lang.BotChannel.Command:      setBotchannel,
		lang.UnsetBotChannel.Command: unsetBotchannel,
		lang.Cleanse.Command:         cleansing,
		lang.Language.Command:        languageReactions,
		lang.Role.Command:            roleReactions,
	}
}

func setBotchannel(s *discordgo.Session, m *discordgo.Message) error {
	persistent.Botchannel = m.ChannelID
	err := botchannelHelpMessage(s)
	return err
}

func unsetBotchannel(s *discordgo.Session, m *discordgo.Message) error {
	err := botchannelHelpMessageDelete(s)
	persistent.Botchannel = ""
	return err
}

func cleansing(s *discordgo.Session, m *discordgo.Message) error {
	if !validChannelID(s, persistent.Botchannel) {
		return nil
	}

	messages, err := s.ChannelMessages(persistent.Botchannel, 0, "", "", "")
	if err != nil {
		return errors.Wrap(err, "while getting messages from bot channel")
	}

	for _, message := range messages {
		if !persistent.contains(message.ID) {
			err = s.ChannelMessageDelete(persistent.Botchannel, message.ID)
			if err != nil {
				return errors.Wrap(err, "while deleting messages from bot channel")
			}
		}
	}

	return nil
}

func languageReactions(s *discordgo.Session, m *discordgo.Message) error {
	persistent.LangMSG = m.ID

	for l := range language.Languages {
		err := s.MessageReactionAdd(m.ChannelID, persistent.LangMSG, l)
		if err != nil {
			return errors.Wrapf(err, "while adding %s to language message", l)
		}
	}
	return nil
}

func roleReactions(s *discordgo.Session, m *discordgo.Message) error {
	if !validMessageID(s, m.ChannelID, persistent.Role) {
		persistent.Role = m.ID

		for _, role := range roles {
			err := s.MessageReactionAdd(m.ChannelID, persistent.Role, role.Emoji)
			if err != nil {
				return errors.Wrapf(err, "while adding %s to role message", role.Emoji)
			}
		}
	}
	return nil
}
