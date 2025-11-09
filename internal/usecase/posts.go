package usecase

import (
	"time"

	"github.com/FOMARTEM/newssite-golang/internal/entities"
)

// CreatePost
func (u *Usecase) CreatePost(post entities.Post) (*entities.Post, error) {
	adminRules, err := u.p.SelectUserRulesById(post.UserId)

	if err != nil {
		return nil, err
	}

	if *adminRules < 6 {
		return nil, entities.ErrAccessDenied
	}

	newPost, err := u.p.InsertPost(post)
	if err != nil {
		return nil, err
	}

	return newPost, err
}

// SelectPost
func (u *Usecase) SelectPost(id int) (*entities.Post, error) {
	post, err := u.p.SelectPostById(id)
	if err != nil {
		return nil, err
	} else if post == nil {
		return nil, entities.ErrPostNotFound
	}

	return post, nil
}

// ListPosts
func (u *Usecase) ListPosts() ([]*entities.Post, error) {
	posts, err := u.p.SelectAllPosts()
	if err != nil {
		return nil, err
	}

	return posts, err
}

func (u *Usecase) ListUserPosts(userId int) ([]*entities.Post, error) {
	posts, err := u.p.SelectUserPosts(userId)
	if err != nil {
		return nil, err
	}

	return posts, err
}

func (u *Usecase) UpdatePost(post entities.Post, userId int) (*entities.Post, error) {

	currPost, err := u.p.SelectPostById(post.ID)
	if err != nil {
		return nil, err
	}

	if userId != currPost.UserId {
		return nil, entities.ErrPostNotFound
	}

	adminRules, err := u.p.SelectUserRulesById(userId)

	if err != nil {
		return nil, err
	}

	if *adminRules < 6 {
		return nil, entities.ErrAccessDenied
	}

	updatedPost, err := u.p.UpdatePostById(post)
	if err != nil {
		return nil, err
	}

	return updatedPost, nil
}

func (u *Usecase) HidePost(postId int, userId int) error {
	var hide int

	adminRules, err := u.p.SelectUserRulesById(userId)

	if err != nil {
		return err
	}

	if *adminRules != 7 {
		return entities.ErrAccessDenied
	}

	post, err := u.p.SelectPostById(postId)

	if err != nil {
		return err
	}

	if post.Hide == 0 {
		hide = 1
	} else {
		hide = 0
	}

	UpdateDate := time.Now().Format("2006/01/02")

	err = u.p.EditVisibilityById(postId, hide, UpdateDate)

	if err != nil {
		return err
	}

	return nil
}

func (u *Usecase) DeletePost(id int, userId int) error {
	currPost, err := u.p.SelectPostById(id)
	if err != nil {
		return err
	}

	if userId != currPost.UserId {
		return entities.ErrPostNotFound
	}

	adminRules, err := u.p.SelectUserRulesById(userId)

	if err != nil {
		return err
	}

	if *adminRules < 6 {
		return entities.ErrAccessDenied
	}

	if err := u.p.DeleteCommentsInPost(id); err != nil {
		return err
	}

	if err := u.p.DeletePostById(id); err != nil {
		return err
	}

	return nil
}
