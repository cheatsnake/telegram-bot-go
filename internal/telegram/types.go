package telegram

type User struct {
	ID int `json:"id"`
	Username string `json:"username"`
}

type Chat struct {
	ID int `json:"id"`
}

type Message struct {
	Text string `json:"text"`
	From User `json:"from"`
	Chat Chat `json:"chat"`
	Date int `json:"date"`
}

type Update struct {
	ID      int    `json:"update_id"`
	Message *Message `json:"message"`
}

type Response[T any] struct {
	Ok     bool     `json:"ok"`
	Result T `json:"result"`
}
