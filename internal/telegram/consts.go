package telegram

const apiHost string = "api.telegram.org"
const botPrefix string = "bot"

const (
	getUpdatesMethod  string = "getUpdates"
	sendMessageMethod string = "sendMessage"
	getMeMethod string = "getMe"
)

const (
	offsetQuery string = "offset"
	limitQuery  string = "limit"
	chatIDQuery string = "chat_id"
	textQuery   string = "text"
)