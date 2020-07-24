package bot

import (
	"ACCPostminister/language"
	"ACCPostminister/projektør"
	"fmt"
)

var commands = map[string]func() (string, bool){
	"ping":                     ping,
	"pong":                     pong,
	language.Help:              help,
	language.BotChannelCommand: botChannel,
	language.CleanseCommand:    cleanse,
	language.ListMovies:        projektør.ListMovies,
	language.RoleCommand:       role,
}

func ping() (string, bool) {
	return "Pong!", true
}

func pong() (string, bool) {
	return "Ping!", true
}

func help() (string, bool) {
	return "no", true
}

func botChannel() (string, bool) {
	return language.BotChannelResponse, true
}

func cleanse() (string, bool) {
	if globalIDs.Botchannel == "" {
		return language.BotChannelUnsetError, true
	}
	return language.CleanseResponse, true
}

func role() (string, bool) {
	if globalIDs.Role != "" {
		return language.RoleExistsError, true
	}

	response := language.RoleResponse

	for _, role := range roles {
		rlstr := fmt.Sprintf("\n%s - %s", role.Emoji, role.Name)
		response += rlstr
	}

	return response, false
}
