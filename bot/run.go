package bot

import (
	"ACCPostminister/language"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var lang = language.Languages["English"]

func Run(s *discordgo.Session) error {
	log.Printf("Bot is now running. Hello!\nPress CTRL-C to exit . . .\n")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	return nil
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself, or outside the set botchannel
	if m.Author.ID == s.State.User.ID || (validChannelID(s, globalIDs.Botchannel) && globalIDs.Botchannel != m.ChannelID) {
		return
	}

	if command := commands[m.Content]; command != nil {
		err := s.MessageReactionAdd(m.ChannelID, m.ID, lang.Reaction)
		if err != nil {
			log.Printf("while reacting to command %s on channel %s: %v", m.Content, m.ChannelID, err)
		}

		resp, conf := command(s, m)
		msg, err := messageSend(s, m, resp, conf)
		if err != nil {
			log.Printf("while sending response to command %s on channel %s: %v", m.Content, m.ChannelID, err)
		}

		if effect := effects[m.Content]; effect != nil {
			err = effect(s, msg)
			if err != nil {
				log.Printf("while processing effects of command %s on channel %s: %v", m.Content, m.ChannelID, err)
			}
		}

		if validMessageID(s, m.ChannelID, globalIDs.User) {
			err = s.ChannelMessageDelete(m.ChannelID, globalIDs.User)
			if err != nil {
				log.Printf("while deleting message %s: %v", globalIDs.User, err)
			}
		}

		globalIDs.User = m.ID
	}
}

func messageSend(s *discordgo.Session, m *discordgo.MessageCreate, message string, isConfirmation bool) (*discordgo.Message, error) {
	if isConfirmation {
		return confirm(s, m.ChannelID, message)
	}

	return s.ChannelMessageSend(m.ChannelID, message)
}

func messageReactionAdd(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	if r.MessageID != globalIDs.Role || r.UserID == s.State.User.ID {
		return
	}

	err := rolechange(s, r.MessageReaction, lang.Role.ConfirmAdd)
	if err != nil {
		log.Printf("while processing reaction on channel %s: %v", r.ChannelID, err)
	}
}

func messageReactionRemove(s *discordgo.Session, r *discordgo.MessageReactionRemove) {
	if r.MessageID != globalIDs.Role || r.UserID == s.State.User.ID {
		return
	}

	err := rolechange(s, r.MessageReaction, lang.Role.ConfirmRemove)
	if err != nil {
		log.Printf("while processing reaction on channel %s: %v", r.ChannelID, err)
	}
}
