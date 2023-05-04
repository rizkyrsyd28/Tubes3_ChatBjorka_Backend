package entity

type Qna struct {
	idQna    int    `json:"id_qna"`
	Question string `json:"quest"`
	Answer   string `json:"ans"`
}
