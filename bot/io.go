package bot

import (
	"encoding/json"
	"io/ioutil"

	"github.com/pkg/errors"
)

const (
	rolefile    = "roles.json"
	messagefile = "messages.json"
)

func saveToFile(target interface{}, filename string) error {
	bytes, err := json.MarshalIndent(target, "", "    ")
	if err != nil {
		return errors.Wrap(err, "while marshaling")
	}

	err = ioutil.WriteFile(filename, bytes, 0644)
	if err != nil {
		return errors.Wrapf(err, "while writing to file %s", filename)
	}

	return nil
}
