package language

var english = Language{
	Reaction: "👌",
	Help: Feature{
		Command:  "help",
		Response: "Instructions for use:",
	},
	BotChannel: Feature{
		Command:  "botchannel",
		Help:     " - assign current channel as the bot channel",
		Response: "Bot channel set to current channel",
	},
	Cleanse: Feature{
		Command:  "cleanse",
		Help:     " - delete all irrelevant messages from the bot channel",
		Response: "Cleansing messages from bot channel",
		Error:    "Bot channel not set or invalid",
	},
	Role: RoleFeature{
		Feature: Feature{
			Command:  "role",
			Help:     " - get role assignment message",
			Response: "React with corresponing emoji to toggle role assignment",
			Error:    "Role message exists",
		},
		ConfirmAdd:    "%s added to role %s",
		ConfirmRemove: "%s removed from role %s",
	},
	Movies: Feature{
		Command: "list movies",
	},
}
