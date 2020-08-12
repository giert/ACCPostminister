package language

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/pkg/errors"
)

var def = ""

var Languages = map[string]string{
	"🇬🇧": "english.json",
}

type Language struct {
	Reaction        string
	Help            Feature
	BotChannel      Feature
	UnsetBotChannel Feature
	Cleanse         Feature
	Language        LanguageFeature
	Role            RoleFeature
	Movies          Feature
}

type Feature struct {
	Command  string
	Help     string
	Response string
	Error    string
}

type LanguageFeature struct {
	Feature
	ConfirmChange string
}

type RoleFeature struct {
	Feature
	ConfirmAdd    string
	ConfirmRemove string
}

func Init() {
	for l := range Languages {
		def = l
		return
	}
}

func Load(language string) (l Language, err error) {
	if language == "" {
		language = def
	}
	err = ReadFromFile(&l, "language/"+Languages[language])
	return
}

func (l Language) GetHelpStrings() (result string) {
	result += getHelpString(l.Help)
	result += getHelpString(l.BotChannel)
	result += getHelpString(l.UnsetBotChannel)
	result += getHelpString(l.Cleanse)
	result += getHelpString(l.Role.Feature)
	result += getHelpString(l.Movies)
	return
}

func getHelpString(feature Feature) string {
	if feature.Help == "" {
		return ""
	}
	return "\n" + feature.Command + feature.Help
}

// MOVE
func ReadFromFile(target interface{}, filename string) error {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Printf("Error reading from %s\nContinuing...\n", filename)
		return nil
	}

	err = json.Unmarshal(bytes, target)
	if err != nil {
		return errors.Wrap(err, "while unmarshaling json")
	}

	return nil
}
