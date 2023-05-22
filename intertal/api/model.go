package api

type CreateRequest struct {
	Name    string `json:"name" binding:"gte=0"`
	Contact string `json:"contact" binding:"gte=0"`
	//Contact     Contact `json:"contact"`
	Description string `json:"description" binding:"gte=0"`
}

//type Contact struct {
//	PhoneNumber string `json:"phone"`
//	Email       string `json:"e-mail" binding:"gte=0"`
//	Telegram    string `json:"telegram"`
//}
