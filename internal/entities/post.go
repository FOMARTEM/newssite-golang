package entities

type Post struct {
	ID         int    `json:"id,omitempty"`
	Name       string `json:"title" validate:"required,min=3,max=255"`
	Text       string `json:"body"  validate:"required,min=50"`
	CreateDate string `json:"createdate,omitempty"`
	UpdateDate string `json:"updatedate,omitempty"`
	UserId     int    `json:"userid,omitempty"`
	UserName   string `json:"username,omitempty"`
}
