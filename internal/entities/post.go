package entities

type Post struct {
	ID         int    `json:"id,omitempty"`
	Name       string `json:"name" validate:"required,min=10,max=255"`
	Text       string `json:"text"`
	CreateDate string `json:"createdate"`
	UpdateDate string `json:"updatedate"`
	UserId     int    `json:"userid"`
}
