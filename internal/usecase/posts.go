package usecase

import (
	"github.com/FOMARTEM/newssite-golang/internal/entities"
)

// CreatePost
func (u *Usecase) CreatePost(post entities.Post) (*entities.Post, error) {
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

func (u *Usecase) UpdatePost(post entities.Post) (*entities.Post, error) {
	_, err := u.p.SelectPostById(post.ID)
	if err != nil {
		return nil, err
	}

	updatedUser, err := u.p.UpdatePostById(post)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (u *Usecase) DeletePost(id int) error {
	if err := u.p.DeletePostById(id); err != nil {
		return err
	}

	return nil
}
