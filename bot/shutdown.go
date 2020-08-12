package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

func Shutdown(s *discordgo.Session) error {
	err := saveToFile(roles, rolefile)
	if err != nil {
		return errors.Wrap(err, "while saving roles")
	}

	err = saveToFile(persistent, messagefile)
	if err != nil {
		return errors.Wrap(err, "while saving message IDs")
	}
	saveToFile(lang, "english.json")
	s.Close()

	log.Println("Bot shutdown. Good bye!")
	return nil
}
