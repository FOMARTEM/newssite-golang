package entities

type Admin struct {
	ID      int    `json:"id"`
	User_id int    `json:"user_id"`
	Email   string `json:"email"`
}
