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
	if msgIDs.Role == "" {
		msgIDs.Role = msg.ID

		for _, role := range roles {
			err := s.MessageReactionAdd(msg.ChannelID, msgIDs.Role, role.Emoji)
			if err != nil {
				return errors.Wrapf(err, "while adding %s to role message", role.Emoji)
			}
		}
	}
	return nil
}
