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

	if command := commands[m.Content]; command != nil {
		msg, err := s.ChannelMessageSend(m.ChannelID, command(m))
		if err != nil {
			log.Printf("while sending response to command %s on channel %s: %v", m.Content, m.ChannelID, err)
		}

		if effect := effects[m.Content]; effect != nil {
			err = effect(s, msg)
			if err != nil {
				log.Printf("while processing effects of command %s on channel %s: %v", m.Content, m.ChannelID, err)
			}
		}
	}
}

func messageReactionAdd(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	if r.MessageID != roleMessageID || r.UserID == s.State.User.ID {
		return
	}

	err := rolechange(s, r.MessageReaction, addRole)
	if err != nil {
		log.Printf("while processing reaction on channel %s: %v", r.ChannelID, err)
	}
}

func messageReactionRemove(s *discordgo.Session, r *discordgo.MessageReactionRemove) {
	if r.MessageID != roleMessageID || r.UserID == s.State.User.ID {
		return
	}

	err := rolechange(s, r.MessageReaction, removeRole)
	if err != nil {
		log.Printf("while processing reaction on channel %s: %v", r.ChannelID, err)
	}
}
