package entities

type Comment struct {
	ID          int    `json:"id,omitempty"`
	CommentText string `json:"comment"`
	PostId      int    `json:"postid"`
	UserId      int    `json:"userid,omitempty"`
	UserName    string `json:"username,omitempty"`
}
