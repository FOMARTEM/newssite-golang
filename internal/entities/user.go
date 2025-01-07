package entities

type User struct {
	ID       int    `json:"id,omitempty"`
	Name     string `json:"name" validate:"required,min=3,max=255"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=255"`
	Admin    bool   `json:"admin,omitempty"`
	Token    string `json:"token,omitempty"`
}
