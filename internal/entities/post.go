package entities

type Post struct {
	ID         int    `json:"id,omitempty"`
	Name       string `json:"title"`
	Text       string `json:"body"`
	CreateDate string `json:"createdate"`
	UpdateDate string `json:"updatedate"`
	UserId     int    `json:"userid"`
}
