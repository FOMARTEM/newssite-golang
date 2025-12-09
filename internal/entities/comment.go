package entities

type Comment struct {
	ID          int    `json:"id,omitempty"`
	CommentText string `json:"comment" form:"comment"`
	PostId      int    `json:"postid" form:"postid"`
	UserId      int    `json:"userid,omitempty"`
	UserName    string `json:"username,omitempty"`
}
