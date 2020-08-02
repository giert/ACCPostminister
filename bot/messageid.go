package bot

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

type IDs struct {
	Botchannel   string
	Confirmation string
	User         string
	Help         string
	Role         string
}

var globalIDs IDs

func (ids IDs) contains(messageID string) bool {
	switch messageID {
	case ids.Botchannel:
		return true
	case ids.Confirmation:
		return true
	case ids.User:
		return true
	case ids.Help:
		return true
	case ids.Role:
		return true
	default:
		return false
	}
}

func (ids IDs) init(s *discordgo.Session) error {
	err := readFromFile(&ids, messagefile)
	if err != nil {
		return errors.Wrap(err, "reading messages from file")
	}

	err = ids.ensureMessageID(s, lang.Help.Response, &ids.Help)
	err = ids.ensureMessageID(s, lang.Role.Response, &ids.Role)

	return nil
}

func (ids IDs) ensureMessageID(s *discordgo.Session, message string, messageID *string) (err error) {
	if validChannelID(s, ids.Botchannel) && !validMessageID(s, ids.Botchannel, *messageID) {
		*messageID, err = findMessageID(s, ids.Botchannel, message)
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

func botchannelHelpMessage() {
	//if !validChannelID(globalIDs.Botchannel) || validMessageID(globalIDs.Help) {
	//	return
	//}

	//send
	//pin
	//save
	//add startup search
}
