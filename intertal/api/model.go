package api

type CreateRequest struct {
	Contact string `json:"number" binding:"gte=0"`
}
