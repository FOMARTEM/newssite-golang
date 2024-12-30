package entities

type Post struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Text       string `json:"text"`
	CreateDate string `json:"createdate"`
	UpdateDate string `json:"updatedate"`
	EditDate   string `json:"editdate"`
	UserId     int    `json:"userid"`
}
