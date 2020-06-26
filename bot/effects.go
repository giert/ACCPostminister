package bot

import (
	"ACCPostminister/language"

	"github.com/bwmarrin/discordgo"
)

type tst interface {
}

var effects = map[string]func(*discordgo.Session, *discordgo.Message) error{
	language.RoleCommand: setRole,
}

func setRole(s *discordgo.Session, msg *discordgo.Message) error {
	roleMessageID = msg.ID

	err := s.MessageReactionAdd(msg.ChannelID, roleMessageID, "😃")
	if err != nil {
		return err
	}

	err = s.MessageReactionAdd(msg.ChannelID, roleMessageID, "🙂")
	return err
}
