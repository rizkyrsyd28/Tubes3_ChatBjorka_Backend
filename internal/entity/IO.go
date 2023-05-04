package entity

type Input struct {
	Prompt string `json:"prompt"`
}

type Output struct {
	Respond string `json:"respond"`
}

type HistoryOutput struct {
	Title   string `json:"title"`
	IDTitle string `json:"id_title"`
}

type ChatHistoryOutput struct {
	User string `json:"user"`
	Bot  string `json:"bot"`
}
