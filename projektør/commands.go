package projektør

import (
	"ACCPostminister/datastore"

	"github.com/bwmarrin/discordgo"
)

func ListMovies(m *discordgo.MessageCreate) (string, bool) {
	s := datastore.Storage{
		Data: map[string]datastore.Entry{
			"Unstoppable":           movie{},
			"The baron and the kid": movieHolder{},
		},
	}

	return s.ListMovies(), true
}
