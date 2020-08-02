package bot

import (
	"ACCPostminister/projektør"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var commands = map[string]func(s *discordgo.Session, m *discordgo.MessageCreate) (string, bool){
	"ping":                  ping,
	"pong":                  pong,
	lang.Help.Command:       help,
	lang.BotChannel.Command: botChannel,
	lang.Cleanse.Command:    cleanse,
	lang.Movies.Command:     projektør.ListMovies,
	lang.Role.Command:       role,
}

func ping(_ *discordgo.Session, _ *discordgo.MessageCreate) (string, bool) {
	return "Pong!", true
}

func pong(_ *discordgo.Session, _ *discordgo.MessageCreate) (string, bool) {
	return "Ping!", true
}

func help(_ *discordgo.Session, _ *discordgo.MessageCreate) (string, bool) {
	return lang.Help.Response + lang.GetHelpStrings(), true
}

func botChannel(_ *discordgo.Session, _ *discordgo.MessageCreate) (string, bool) {
	return lang.BotChannel.Response, true
}

func cleanse(s *discordgo.Session, _ *discordgo.MessageCreate) (string, bool) {
	if !validChannelID(s, globalIDs.Botchannel) {
		return lang.BotChannel.Error, true
	}
	return lang.Cleanse.Response, true
}

func role(s *discordgo.Session, m *discordgo.MessageCreate) (string, bool) {
	if validMessageID(s, m.ChannelID, globalIDs.Role) {
		return lang.Role.Error, true
	}

	response := lang.Role.Response

	for _, role := range roles {
		rlstr := fmt.Sprintf("\n%s - %s", role.Emoji, role.Name)
		response += rlstr
	}

	return response, false
}
