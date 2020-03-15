package projektør

import (
	"ACCPostminister/datastore"

	"github.com/bwmarrin/discordgo"
)

func listMovies(m *discordgo.MessageCreate) string {
	s := datastore.Storage{
		Data: map[string]datastore.Entry{
			"Unstoppable":           movie{},
			"The baron and the kid": movieHolder{},
		},
	}

	return s.listMovies()
}
