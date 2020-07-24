package projektør

import (
	"ACCPostminister/datastore"
)

func ListMovies() (string, bool) {
	s := datastore.Storage{
		Data: map[string]datastore.Entry{
			"Unstoppable":           movie{},
			"The baron and the kid": movieHolder{},
		},
	}

	return s.ListMovies(), true
}
