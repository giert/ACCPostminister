package bot

import (
	"ACCPostminister/language"
	"ACCPostminister/projektør"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var commands = map[string]func(*discordgo.MessageCreate) string{
	"ping":               ping,
	"pong":               pong,
	language.Help:        help,
	language.ListMovies:  projektør.ListMovies,
	language.RoleCommand: role,
}

func ping(m *discordgo.MessageCreate) string {
	return "Pong!"
}

func pong(m *discordgo.MessageCreate) string {
	return "Ping!"
}

func help(m *discordgo.MessageCreate) string {
	return "no"
}

func role(m *discordgo.MessageCreate) string {
	response := language.RoleResponse

	for _, role := range roles {
		rlstr := fmt.Sprintf("\n%s - %s", role.emoji, role.name)
		response += rlstr
	}

	return response
}
