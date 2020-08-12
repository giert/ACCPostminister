package bot

import (
	"ACCPostminister/projektør"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var commands = map[string]func(s *discordgo.Session, m *discordgo.MessageCreate) (string, bool){}

func initCommands() {
	commands = map[string]func(s *discordgo.Session, m *discordgo.MessageCreate) (string, bool){
		"ping":                       ping,
		"pong":                       pong,
		lang.Help.Command:            help,
		lang.BotChannel.Command:      botChannel,
		lang.UnsetBotChannel.Command: unsetBotChannel,
		lang.Cleanse.Command:         cleanse,
		lang.Language.Command:        languageChange,
		lang.Role.Command:            role,
		lang.Movies.Command:          projektør.ListMovies,
	}
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

func unsetBotChannel(_ *discordgo.Session, _ *discordgo.MessageCreate) (string, bool) {
	return lang.UnsetBotChannel.Response, true
}

func cleanse(s *discordgo.Session, _ *discordgo.MessageCreate) (string, bool) {
	if !validChannelID(s, persistent.Botchannel) {
		return lang.BotChannel.Error, true
	}
	return lang.Cleanse.Response, true
}

func languageChange(_ *discordgo.Session, _ *discordgo.MessageCreate) (string, bool) {
	return lang.Language.Response, false
}

func role(s *discordgo.Session, m *discordgo.MessageCreate) (string, bool) {
	if validMessageID(s, m.ChannelID, persistent.Role) {
		return lang.Role.Error, true
	}

	response := lang.Role.Response

	for _, role := range roles {
		rlstr := fmt.Sprintf("\n%s - %s", role.Emoji, role.Name)
		response += rlstr
	}

	return response, false
}
