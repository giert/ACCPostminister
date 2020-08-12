package bot

import (
	"ACCPostminister/language"
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

	language.Init()

	err = initRoles(rolefile)
	if err != nil {
		return nil, errors.Wrap(err, "while initiating configured roles")
	}

	err = persistent.init(s)
	if err != nil {
		return nil, errors.Wrap(err, "while initiating global IDs")
	}

	lang, err = language.Load(persistent.Language)
	if err != nil {
		return nil, err
	}

	initCommands()
	initEffects()

	return s, nil
}
