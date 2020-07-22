package bot

import (
	"ACCPostminister/language"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

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

	err := s.MessageReactionAdd(m.ChannelID, m.ID, language.RecievedEmoji)
	if err != nil {
		log.Printf("while reacting to command %s on channel %s: %v", m.Content, m.ChannelID, err)
	}

	// like after recieve, generalize th check if then
	if command := commands[m.Content]; command != nil {
		resp, conf := command(m)
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
	}

	if msgIDs.user != "" {
		err = s.ChannelMessageDelete(m.ChannelID, msgIDs.user)
		if err != nil {
			log.Printf("while deleting message %s: %v", msgIDs.user, err)
		}
	}

	msgIDs.user = m.ID
}

func messageSend(s *discordgo.Session, m *discordgo.MessageCreate, message string, isConfirmation bool) (*discordgo.Message, error) {
	if isConfirmation {
		return confirm(s, m.ChannelID, message)
	}

	return s.ChannelMessageSend(m.ChannelID, message)
}

func messageReactionAdd(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	if r.MessageID != msgIDs.role || r.UserID == s.State.User.ID {
		return
	}

	err := rolechange(s, r.MessageReaction, addRole)
	if err != nil {
		log.Printf("while processing reaction on channel %s: %v", r.ChannelID, err)
	}
}

func messageReactionRemove(s *discordgo.Session, r *discordgo.MessageReactionRemove) {
	if r.MessageID != msgIDs.role || r.UserID == s.State.User.ID {
		return
	}

	err := rolechange(s, r.MessageReaction, removeRole)
	if err != nil {
		log.Printf("while processing reaction on channel %s: %v", r.ChannelID, err)
	}
}
