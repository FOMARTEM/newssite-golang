package entities

type Post struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Text       string `json:"text"`
	CreateDate string `json:"createdate"`
	EditDate   string `json:"editdate"`
	Admin      int    `json:"admin"`
}
