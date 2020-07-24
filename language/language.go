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
