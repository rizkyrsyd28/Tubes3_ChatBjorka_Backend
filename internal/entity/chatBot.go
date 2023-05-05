package entity

type Qna struct {
	IDQna    int    `json:"id_qna"`
	Question string `json:"quest"`
	Answer   string `json:"ans"`
}
