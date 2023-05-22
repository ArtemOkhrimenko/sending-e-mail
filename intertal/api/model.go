package api

type CreateRequest struct {
	Name        string `json:"name" binding:"gte=0"`
	Contact     string `json:"contact" binding:"gte=0"`
	Description string `json:"description" binding:"gte=0"`
}
