package bot

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

type persistentData struct {
	Language     string
	Botchannel   string
	AdminUsers   []string
	Confirmation string
	User         string
	Help         string
	LangMSG      string
	Role         string
}

var persistent persistentData

func (p persistentData) contains(messageID string) bool {
	switch messageID {
	case p.Botchannel:
		return true
	case p.Confirmation:
		return true
	case p.User:
		return true
	case p.Help:
		return true
	case p.LangMSG:
		return true
	case p.Role:
		return true
	default:
		return false
	}
}

func (p *persistentData) init(s *discordgo.Session) error {
	err := ReadFromFile(&p, messagefile)
	if err != nil {
		return errors.Wrap(err, "reading messages from file")
	}

	err = p.ensureMessageID(s, lang.Help.Response, &p.Help)
	err = p.ensureMessageID(s, lang.Role.Response, &p.Role)

	return nil
}

func (p *persistentData) ensureMessageID(s *discordgo.Session, message string, messageID *string) (err error) {
	if validChannelID(s, p.Botchannel) && !validMessageID(s, p.Botchannel, *messageID) {
		*messageID, err = findMessageID(s, p.Botchannel, message)
	}
	return
}

func findMessageID(s *discordgo.Session, channelID, partialMessage string) (string, error) {
	messages, err := s.ChannelMessages(channelID, 0, "", "", "")
	if err != nil {
		return "", errors.Wrap(err, "while searching for message")
	}

	for _, message := range messages {
		if strings.Contains(message.Content, partialMessage) {
			return message.ID, nil
		}
	}

	return "", nil
}

// is not empty and actually exists
func validMessageID(s *discordgo.Session, channelID, messageID string) bool {
	msg, err := s.ChannelMessage(channelID, messageID)
	return messageID != "" || msg != nil || err == nil
}

func validChannelID(s *discordgo.Session, channelID string) bool {
	ch, err := s.Channel(channelID)
	return channelID != "" || ch != nil || err == nil
}

func isAdmin(u *discordgo.User) bool {
	if len(persistent.AdminUsers) == 0 {
		return true
	}

	for _, a := range persistent.AdminUsers {
		if a == u.ID {
			return true
		}
	}
	
	return false
}

func botchannelHelpMessage(s *discordgo.Session) error {
	if !validChannelID(s, persistent.Botchannel) || validMessageID(s, persistent.Botchannel, persistent.Help) {
		return nil
	}

	message, err := s.ChannelMessageSend(persistent.Botchannel, lang.Help.Response+lang.GetHelpStrings()) // two occurences of lang+strings
	if err != nil {
		return err
	}

	err = s.ChannelMessagePin(persistent.Botchannel, message.ID)
	if err != nil {
		return err
	}

	persistent.Help = message.ID

	return nil
}

func botchannelHelpMessageDelete(s *discordgo.Session) error {
	if !validChannelID(s, persistent.Botchannel) || !validMessageID(s, persistent.Botchannel, persistent.Help) {
		return nil
	}

	err := s.ChannelMessageDelete(persistent.Botchannel, persistent.Help)
	if err != nil {
		return err
	}

	persistent.Help = ""

	return nil
}
