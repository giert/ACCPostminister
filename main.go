package main

import (
	"ACCPanorama/language"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

const (
	token = "NjgxODYwMTg4MTA5NjY4MzU2.XlUmEg.rsw0ahyJsWCHz5Hh3j50UwDCGjw"
)

var commands = map[string]func(*discordgo.MessageCreate) string{
	"ping":              ping,
	"pong":              pong,
	language.ListMovies: listMovies,
}

func main() {
	session, err := startup()
	if err != nil {
		log.Fatal("While setting up bot:", err)
	}

	err = run(session)
	if err != nil {
		log.Fatal("While running bot:", err)
	}

	err = shutdown(session)
	if err != nil {
		log.Fatal("While shutting down bot:", err)
	}
}

func startup() (*discordgo.Session, error) {
	s, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, errors.Wrap(err, "while creating Discord session")
	}

	s.AddHandler(messageCreate)

	err = s.Open()
	if err != nil {
		return nil, errors.Wrap(err, "while opening connection")
	}

	return s, nil
}

func run(s *discordgo.Session) error {
	log.Printf("Bot is now running. Hello!\nPress CTRL-C to exit . . .\n")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	return nil
}

func shutdown(s *discordgo.Session) error {
	s.Close()
	log.Printf("Bot shutdown. Good bye!")
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

func listMovies(m *discordgo.MessageCreate) string {
	s := storage{
		Data: map[string]entry{
			"Unstoppable":           movie{},
			"The baron and the kid": movieHolder{},
		},
	}

	return s.listMovies()
}
