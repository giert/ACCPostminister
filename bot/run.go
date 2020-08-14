package bot

import (
	"ACCPostminister/language"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var lang = language.Language{}

func Run(s *discordgo.Session) error {
	log.Printf("Bot is now running. Hello!\nPress CTRL-C to exit . . .\n")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	return nil
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself, or outside the set botchannel
	if m.Author.ID == s.State.User.ID || (validChannelID(s, persistent.Botchannel) && persistent.Botchannel != m.ChannelID) {
		return
	}

	if command := commands[m.Content]; command != nil {
		err := s.MessageReactionAdd(m.ChannelID, m.ID, lang.Reaction)
		if err != nil {
			log.Printf("while reacting to command %s on channel %s: %v", m.Content, m.ChannelID, err)
		}

		resp, conf, adm := command(s, m)
		msg, err := messageSend(s, m, resp, conf, adm)
		if err != nil {
			log.Printf("while sending response to command %s on channel %s: %v", m.Content, m.ChannelID, err)
		}

		if effect := effects[m.Content]; effect != nil {
			err = effect(s, msg)
			if err != nil {
				log.Printf("while processing effects of command %s on channel %s: %v", m.Content, m.ChannelID, err)
			}
		}

		if validMessageID(s, m.ChannelID, persistent.User) {
			err = s.ChannelMessageDelete(m.ChannelID, persistent.User)
			if err != nil {
				log.Printf("while deleting message %s: %v", persistent.User, err)
			}
		}

		persistent.User = m.ID
	}
}

func messageSend(s *discordgo.Session, m *discordgo.MessageCreate, message string, isConfirmation bool, requiresAdmin bool) (*discordgo.Message, error) {
	if requiresAdmin && !isAdmin(m.Author) {
		message = lang.AdminError
		isConfirmation = true
	}

	if isConfirmation {
		return confirm(s, m.ChannelID, message)
	}

	return s.ChannelMessageSend(m.ChannelID, message)
}

func messageReactionAdd(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	if r.UserID == s.State.User.ID {
		return
	}

	if r.MessageID == persistent.LangMSG {
		err := langchg(s, r)
		if err != nil {
			log.Printf("while processing reaction on channel %s: %v", r.ChannelID, err)
		}
	}

	if r.MessageID == persistent.Role {
		err := rolechange(s, r.MessageReaction, lang.Role.ConfirmAdd)
		if err != nil {
			log.Printf("while processing reaction on channel %s: %v", r.ChannelID, err)
		}
	}
}

func langchg(s *discordgo.Session, r *discordgo.MessageReactionAdd) (err error) {
	persistent.Language = r.MessageReaction.Emoji.Name

	lang, err = language.Load(persistent.Language)
	if err != nil {
		return err
	}

	_, err = confirm(s, r.ChannelID, fmt.Sprintf(lang.Language.ConfirmChange+persistent.Language))
	if err != nil {
		return err
	}

	err = s.ChannelMessageDelete(r.ChannelID, persistent.LangMSG)
	if err != nil {
		return err
	}

	return nil
}

func messageReactionRemove(s *discordgo.Session, r *discordgo.MessageReactionRemove) {
	if r.MessageID != persistent.Role || r.UserID == s.State.User.ID {
		return
	}

	err := rolechange(s, r.MessageReaction, lang.Role.ConfirmRemove)
	if err != nil {
		log.Printf("while processing reaction on channel %s: %v", r.ChannelID, err)
	}
}
