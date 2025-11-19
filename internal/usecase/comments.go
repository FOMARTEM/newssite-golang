package usecase

import (
	"github.com/FOMARTEM/newssite-golang/internal/entities"
)

func (u *Usecase) CreateComment(comment entities.Comment) (*entities.Comment, error) {
	adminRules, err := u.p.SelectUserRulesById(comment.UserId)

	if err != nil {
		return nil, err
	}

	if *adminRules < 1 {
		return nil, entities.ErrCommentNotFound
	}

	createdComment, err := u.p.InsertComment(comment)

	if err != nil {
		return nil, err
	}

	return createdComment, nil
}

func (u *Usecase) PostComments(post_id int) ([]*entities.Comment, error) {
	comments, err := u.p.GetCommentsForPost(post_id)

	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (u *Usecase) UpdateComment(comment entities.Comment) (*entities.Comment, error) {
	newComment, err := u.p.UpdateComment(comment)

	if err != nil {
		return nil, err
	} else if newComment == nil {
		return newComment, entities.ErrCommentNotFound
	}

	return newComment, nil
}

func (u *Usecase) DeleteComment(id int) error {
	err := u.p.DeleteComment(id)

	if err != nil {
		return err
	}

	return nil
}

func (u *Usecase) DeleteCommentsInPost(post_id int) error {
	err := u.p.DeleteCommentsInPost(post_id)

	if err != nil {
		return err
	}

	return nil
}
