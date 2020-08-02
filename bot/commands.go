package bot

import (
	"ACCPostminister/projektør"
	"fmt"
)

var commands = map[string]func() (string, bool){
	"ping":                  ping,
	"pong":                  pong,
	lang.Help.Command:       help,
	lang.BotChannel.Command: botChannel,
	lang.Cleanse.Command:    cleanse,
	lang.Movies.Command:     projektør.ListMovies,
	lang.Role.Command:       role,
}

func ping() (string, bool) {
	return "Pong!", true
}

func pong() (string, bool) {
	return "Ping!", true
}

func help() (string, bool) {
	return lang.Help.Response + lang.GetHelpStrings(), true
}

func botChannel() (string, bool) {
	return lang.BotChannel.Response, true
}

func cleanse() (string, bool) {
	if !validChannelID(globalIDs.Botchannel) {
		return lang.BotChannel.Error, true
	}
	return lang.Cleanse.Response, true
}

func role() (string, bool) {
	if validMessageID(globalIDs.Role) {
		return lang.Role.Error, true
	}

	response := lang.Role.Response

	for _, role := range roles {
		rlstr := fmt.Sprintf("\n%s - %s", role.Emoji, role.Name)
		response += rlstr
	}

	return response, false
}
