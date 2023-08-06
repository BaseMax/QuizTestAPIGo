package entity

type QuestionRequest struct {
	Text string `json:"text" binding:"required"`
}
