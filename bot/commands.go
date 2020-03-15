package bot

import (
	"ACCPostminister/language"
	"ACCPostminister/projektør"

	"github.com/bwmarrin/discordgo"
)

var commands = map[string]func(*discordgo.MessageCreate) string{
	"ping":              ping,
	"pong":              pong,
	language.Help:       help,
	language.ListMovies: projektør.ListMovies,
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
