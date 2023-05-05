package entity

type HistoryOutput struct {
	Title   string `json:"title"`
	IDTitle string `json:"id_title"`
}

type ChatHistoryOutput struct {
	User string `json:"user"`
	Bot  string `json:"bot"`
}

type UserInput struct {
	Message string `json:"message"`
	Algo    string `json:"algo"`
}

type BotOutput struct {
	Message string `json:"respond"`
	IDTitle string `json:"id_title"`
}
