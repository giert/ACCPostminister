package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func Shutdown(s *discordgo.Session) error {
	s.Close()
	log.Printf("Bot shutdown. Good bye!")
	return nil
}
