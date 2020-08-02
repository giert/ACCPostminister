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

	if !validMessageID(ids.Role) && validChannelID(ids.Botchannel) { // MOVE TO FUNC
		ids.Role, err = findMessageID(s, ids.Botchannel, lang.Role.Response)
	}

	if !validMessageID(ids.Role) && validChannelID(ids.Botchannel) { // HELP
		ids.Role, err = findMessageID(s, ids.Botchannel, lang.Role.Response)
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

// is not empty and actually exists
func validMessageID(messageID string) bool {
	return messageID != ""
}

func validChannelID(channelID string) bool {
	return channelID != ""
}

func botchannelHelpMessage() {
	if !validChannelID(globalIDs.Botchannel) || validMessageID(globalIDs.Help) {
		return
	}

	//send
	//pin
	//save
	//add startup search
}
