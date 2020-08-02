package language

var Languages = map[string]Language{
	"English": english,
}

type Language struct {
	Reaction   string
	Help       Feature
	BotChannel Feature
	Cleanse    Feature
	Role       RoleFeature
	Movies     Feature
}

type Feature struct {
	Command  string
	Help     string
	Response string
	Error    string
}

type RoleFeature struct {
	Feature
	ConfirmAdd    string
	ConfirmRemove string
}

func (l Language) GetHelpStrings() (result string) {
	result += getHelpString(l.Help)
	result += getHelpString(l.BotChannel)
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
