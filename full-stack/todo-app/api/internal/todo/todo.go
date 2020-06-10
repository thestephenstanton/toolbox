package todo

type Todo struct {
	UID string `json:"id"`
	Text string `json:"text"`
	IsFinsihed bool `json:"isFinished"`
	ListUID string `json:"listId"`
}