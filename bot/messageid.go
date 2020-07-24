package bot

import (
	"ACCPostminister/language"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

type importantIDs struct {
	Botchannel   string
	Confirmation string
	User         string
	Role         string
}

var globalIDs importantIDs

func (ids importantIDs) contains(messageID string) bool {
	switch messageID {
	case ids.Botchannel:
		return true
	case ids.Confirmation:
		return true
	case ids.User:
		return true
	case ids.Role:
		return true
	default:
		return false
	}
}

func findMessageIDs(s *discordgo.Session) error {
	err := readFromFile(&globalIDs, messagefile)
	if err != nil {
		return errors.Wrap(err, "reading messages from file")
	}

	if globalIDs.Role == "" && globalIDs.Botchannel != "" {
		globalIDs.Role, err = findMessageID(s, globalIDs.Botchannel, language.RoleResponse)
	}

	return nil
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
