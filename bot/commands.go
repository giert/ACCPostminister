package bot

import (
	"ACCPostminister/language"
	"ACCPostminister/projektør"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var commands = map[string]func(*discordgo.MessageCreate) (string, bool){
	"ping":               ping,
	"pong":               pong,
	language.Help:        help,
	language.ListMovies:  projektør.ListMovies,
	language.RoleCommand: role,
}

func ping(m *discordgo.MessageCreate) (string, bool) {
	return "Pong!", true
}

func pong(m *discordgo.MessageCreate) (string, bool) {
	return "Ping!", true
}

func help(m *discordgo.MessageCreate) (string, bool) {
	return "no", true
}

func role(m *discordgo.MessageCreate) (string, bool) {
	if msgIDs.role != "" {
		return language.RoleExistsError, true
	}

	response := language.RoleResponse

	for _, role := range roles {
		rlstr := fmt.Sprintf("\n%s - %s", role.Emoji, role.Name)
		response += rlstr
	}

	return response, false
}
