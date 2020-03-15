package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

const (
	token = "NjgxODYwMTg4MTA5NjY4MzU2.XlUmEg.rsw0ahyJsWCHz5Hh3j50UwDCGjw"
)

func startup() (*discordgo.Session, error) {
	s, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, errors.Wrap(err, "while creating Discord session")
	}

	s.AddHandler(messageCreate)

	err = s.Open()
	if err != nil {
		return nil, errors.Wrap(err, "while opening connection")
	}

	return s, nil
}
