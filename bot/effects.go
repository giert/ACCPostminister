package bot

import (
	"ACCPostminister/language"

	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

var effects = map[string]func(*discordgo.Session, *discordgo.Message) error{
	language.RoleCommand: roleReactions,
}

func roleReactions(s *discordgo.Session, msg *discordgo.Message) error {
	roleMessageID = msg.ID

	for _, role := range roles {
		err := s.MessageReactionAdd(msg.ChannelID, roleMessageID, role.emoji)
		if err != nil {
			return errors.Wrapf(err, "while adding %s to role message", role.emoji)
		}
	}

	return nil
}
