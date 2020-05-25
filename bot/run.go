package bot

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var roleMessageID string
var confirmationMessageID string

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
	roleMessageID = msg.ID
	err = s.MessageReactionAdd(m.ChannelID, msg.ID, "😃")
	err = s.MessageReactionAdd(m.ChannelID, msg.ID, "🙂")
	if err != nil {
		return
	}
}

func messageReactionAdd(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	// Ignore all reactions created by the bot itself
	if r.UserID == s.State.User.ID {
		return
	}

	if r.MessageID == roleMessageID {
		msg, _ := s.ChannelMessageSend(r.ChannelID, r.UserID+" added "+r.Emoji.Name)
		s.ChannelMessageDelete(r.ChannelID, confirmationMessageID)
		confirmationMessageID = msg.ID
	}
}

func messageReactionRemove(s *discordgo.Session, r *discordgo.MessageReactionRemove) {
	// Ignore all reactions removed by the bot itself
	if r.UserID == s.State.User.ID {
		return
	}

	msg, _ := s.ChannelMessageSend(r.ChannelID, r.UserID+" removed "+r.Emoji.Name)
	s.ChannelMessageDelete(r.ChannelID, confirmationMessageID)
	confirmationMessageID = msg.ID
}
