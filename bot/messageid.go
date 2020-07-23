package bot

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/pkg/errors"
)

type messageIDs struct {
	Role         string
	Confirmation string
	User         string
}

var msgIDs messageIDs

func findMessageIDs() error {
	bytes, err := ioutil.ReadFile(messagefile)
	if err != nil {
		log.Printf("Error reading from %s\nContinuing...\n", messagefile)
		return nil
	}

	err = json.Unmarshal(bytes, &msgIDs)
	if err != nil {
		return errors.Wrap(err, "while unmarshaling json")
	}

	return nil
}
