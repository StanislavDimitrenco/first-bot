package main

//structure for send message
type SendMessage struct {
	Ok     bool    `json:"ok"`
	Result Message `json:"result"`
}

type Message struct {
	MessageId int                    `json:"message_id"`
	From      GetUpdateResultMessage `json:"from"`
	Chat      GetUpdateResultChat    `json:"chat"`
	Date      int                    `json:"date"`
	Text      string                 `json:"text"`
}

type GetUpdateResultMessage struct {
	Id        int    `json:"id"`
	IsBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
}

type GetUpdateResultChat struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
	Type      string `json:"type"`
}
