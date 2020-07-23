package bot

import (
	"github.com/pkg/errors"
)

type messageIDs struct {
	Role         string
	Confirmation string
	User         string
}

var msgIDs messageIDs

func findMessageIDs() error {
	err := readFromFile(&msgIDs, messagefile)
	if err != nil {
		return errors.Wrap(err, "reading messages from file")
	}

	return nil
}
