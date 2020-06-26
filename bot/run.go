package bot

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var roleMessageID string

func Run(s *discordgo.Session) error {
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

	msg, err := s.ChannelMessageSend(m.ChannelID, commands[m.Content](m))
	if err != nil {
		log.Fatal(err)
	}

	err = effects[m.Content](s, msg)
	if err != nil {
		log.Fatal(err)
	}
}

func messageReactionAdd(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	// Ignore all reactions created by the bot itself
	if r.UserID == s.State.User.ID {
		return
	}

	// set role

	if r.MessageID == roleMessageID {
		confirm(s, r.ChannelID, r.UserID+" added "+r.Emoji.Name)
	}
}

func messageReactionRemove(s *discordgo.Session, r *discordgo.MessageReactionRemove) {
	// Ignore all reactions removed by the bot itself
	if r.UserID == s.State.User.ID {
		return
	}

	//remove role

	if r.MessageID == roleMessageID {
		confirm(s, r.ChannelID, r.UserID+" removed "+r.Emoji.Name)
	}
}
