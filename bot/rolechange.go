package bot

import (
	"ACCPostminister/language"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/pkg/errors"

	"github.com/bwmarrin/discordgo"
)

const (
	addRole    = language.RoleConfirmAdd
	removeRole = language.RoleConfirmRemove
)

var roles = []struct {
	Name  string
	Emoji string
}{}

func initRoles() error {
	bytes, err := ioutil.ReadFile("roles.json")
	if err != nil {
		return errors.Wrap(err, "while reading json from file")
	}

	err = json.Unmarshal(bytes, &roles)
	if err != nil {
		return errors.Wrap(err, "while unmarshaling json")
	}

	return nil
}

func rolechange(s *discordgo.Session, r *discordgo.MessageReaction, action string) (err error) {
	rl := ""
	if action == addRole {
		rl, err = doRolechange(s, r, s.GuildMemberRoleAdd)
	} else {
		rl, err = doRolechange(s, r, s.GuildMemberRoleRemove)
	}
	if err != nil {
		return
	}

	usr, err := s.GuildMember(r.GuildID, r.UserID)
	if err != nil {
		return
	}

	_, err = confirm(s, r.ChannelID, fmt.Sprintf(action, usr.User.Username, rl))
	if err != nil {
		return errors.Wrapf(err, "while changing roles for user %s", r.UserID)
	}

	return
}

func doRolechange(s *discordgo.Session, r *discordgo.MessageReaction, action func(guildID, userID, roleID string) error) (string, error) {
	for _, role := range roles {
		if role.Emoji == r.Emoji.Name {
			rl, err := getRole(s, r.GuildID, role.Name)
			if err != nil {
				return "", errors.Wrapf(err, "while finding role ID for reaction %s", r.Emoji.Name)
			}

			return role.Name, action(r.GuildID, r.UserID, rl.ID)
		}
	}
	return "", errors.New("role not found for reaction " + r.Emoji.Name)
}

func getRole(s *discordgo.Session, guildID, name string) (*discordgo.Role, error) {
	rls, err := s.GuildRoles(guildID)
	if err != nil {
		return nil, errors.Wrapf(err, "while getting roles for guild %s", guildID)
	}

	for _, rl := range rls {
		if name == rl.Name {
			return rl, nil
		}
	}

	return nil, errors.New("role " + name + " not found for guild " + guildID)
}
