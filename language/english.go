package language

const RecievedEmoji = "👌"

const Help = "help"

const BotChannelCommand = "botchannel"
const BotChannelHelp = BotChannelCommand + " - assign current channel as the bot channel"
const BotChannelResponse = "Bot channel set to current channel"
const BotChannelUnsetError = "Bot channel not set"

const CleanseCommand = "cleanse"
const CleanseHelp = CleanseCommand + " - delete all irrelevant messages from the bot channel"
const CleanseResponse = "Cleansing messages from bot channel"

const ListMovies = "list movies"

const RoleCommand = "role"
const RoleHelp = RoleCommand + " - get role assignment message"
const RoleResponse = "React with corresponing emoji to toggle role assignment"
const RoleExistsError = "Role message exists"
const RoleConfirmAdd = "%s added to role %s"
const RoleConfirmRemove = "%s removed from role %s"
