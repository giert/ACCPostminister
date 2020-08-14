package bot

import (
	"ACCPostminister/projektør"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var commands = map[string]func(s *discordgo.Session, m *discordgo.MessageCreate) (response string, isConfirmation, requiresAdmin bool){}

func initCommands() {
	commands = map[string]func(s *discordgo.Session, m *discordgo.MessageCreate) (response string, isConfirmation, requiresAdmin bool){
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

func ping(_ *discordgo.Session, _ *discordgo.MessageCreate) (response string, isConfirmation, requiresAdmin bool) {
	return "Pong!", true, false
}

func pong(_ *discordgo.Session, _ *discordgo.MessageCreate) (response string, isConfirmation, requiresAdmin bool) {
	return "Ping!", true, false
}

func help(_ *discordgo.Session, _ *discordgo.MessageCreate) (response string, isConfirmation, requiresAdmin bool) {
	return lang.Help.Response + lang.GetHelpStrings(), true, false
}

func botChannel(_ *discordgo.Session, _ *discordgo.MessageCreate) (response string, isConfirmation, requiresAdmin bool) {
	return lang.BotChannel.Response, true, true
}

func unsetBotChannel(_ *discordgo.Session, _ *discordgo.MessageCreate) (response string, isConfirmation, requiresAdmin bool) {
	return lang.UnsetBotChannel.Response, true, true
}

func cleanse(s *discordgo.Session, _ *discordgo.MessageCreate) (response string, isConfirmation, requiresAdmin bool) {
	if !validChannelID(s, persistent.Botchannel) {
		return lang.BotChannel.Error, true, true
	}
	return lang.Cleanse.Response, true, true
}

func languageChange(_ *discordgo.Session, _ *discordgo.MessageCreate) (response string, isConfirmation, requiresAdmin bool) {
	return lang.Language.Response, false, true
}

func role(s *discordgo.Session, m *discordgo.MessageCreate) (response string, isConfirmation, requiresAdmin bool) {
	if validMessageID(s, m.ChannelID, persistent.Role) {
		return lang.Role.Error, true, true
	}

	response = lang.Role.Response
	for _, role := range roles {
		rlstr := fmt.Sprintf("\n%s - %s", role.Emoji, role.Name)
		response += rlstr
	}

	return response, false, true
}
