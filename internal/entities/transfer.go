package entities

type Transfer struct {
	SenderID    int     `json:"sender_id" binding:"required"`
	RecipientID int     `json:"recipient_id" binding:"required"`
	Amount      float64 `json:"amount" binding:"required"`
}
