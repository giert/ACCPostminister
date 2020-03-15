package bot

import (
	"ACCPostminister/language"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var commands = map[string]func(*discordgo.MessageCreate) string{
	"ping":              ping,
	"pong":              pong,
	language.ListMovies: listMovies,
}

func run(s *discordgo.Session) error {
	log.Printf("Bot is now running. Hello!\nPress CTRL-C to exit . . .\n")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	return nil
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	s.ChannelMessageSend(m.ChannelID, commands[m.Content](m))
}

func ping(m *discordgo.MessageCreate) string {
	return "Pong!"
}

func pong(m *discordgo.MessageCreate) string {
	return "Ping!"
}
