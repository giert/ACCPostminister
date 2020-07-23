package bot

import (
	"io/ioutil"

	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

func Startup() (*discordgo.Session, error) {
	token, err := ioutil.ReadFile("token.txt")
	if err != nil {
		return nil, errors.Wrap(err, "while reading token from file")
	}

	s, err := discordgo.New("Bot " + string(token))
	if err != nil {
		return nil, errors.Wrap(err, "while creating Discord session")
	}

	s.AddHandler(messageCreate)
	s.AddHandler(messageReactionAdd)
	s.AddHandler(messageReactionRemove)

	err = s.Open()
	if err != nil {
		return nil, errors.Wrap(err, "while opening connection")
	}

	err = initRoles()
	if err != nil {
		return nil, errors.Wrap(err, "while initiating configured roles")
	}

	err = findMessageIDs()
	if err != nil {
		return nil, errors.Wrap(err, "while finding message IDs")
	}

	return s, nil
}
