package entities

type Admin struct {
	ID     int    `json:"id"`
	UserId int    `json:"user_id"`
	Email  string `json:"email"`
}
