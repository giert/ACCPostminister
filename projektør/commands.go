package projektør

import (
	"ACCPostminister/datastore"

	"github.com/bwmarrin/discordgo"
)

func ListMovies(_ *discordgo.Session, _ *discordgo.MessageCreate) (response string, isConfirmation, requiresAdmin bool) {
	s := datastore.Storage{
		Data: map[string]datastore.Entry{
			"Unstoppable":           movie{},
			"The baron and the kid": movieHolder{},
		},
	}

	return s.ListMovies(), true, false
}
