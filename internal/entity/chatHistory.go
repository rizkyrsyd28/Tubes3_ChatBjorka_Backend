package entity

type ChatContent struct {
	IDChat   int    `json:"id_chat"`
	UserChat string `json:"user_chat"`
	BotChat  string `json:"bot_chat"`
}

type ChatHistory struct {
	IDTitle string        `json:"id_title"`
	Chat    []ChatContent `json:"chat_content"`
}
