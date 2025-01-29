package provider

import (
	"database/sql"
	"errors"

	"github.com/FOMARTEM/newssite-golang/internal/entities"
)

// Функции с таблицей post
// создание поста
func (p *Provider) InsertPost(post entities.Post) (*entities.Post, error) {
	var id int
	err := p.conn.QueryRow(
		`CALL create_post($1, $2, $3, $4, n_id := NULL)`,
		post.Name, post.Text, post.CreateDate, post.UserId,
	).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &entities.Post{
		ID:         id,
		Name:       post.Name,
		Text:       post.Text,
		CreateDate: post.CreateDate,
		UpdateDate: post.CreateDate,
		UserId:     post.UserId,
	}, nil
}

// поиск поста по id
func (p *Provider) SelectPostById(id int) (*entities.Post, error) {
	var post entities.Post
	err := p.conn.QueryRow(
		"SELECT * FROM get_post($1)",
		id,
	).Scan(&post.ID, &post.Name, &post.Text, &post.CreateDate, &post.UpdateDate, &post.UserId, &post.UserName)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

// получение всех постов
func (p *Provider) SelectAllPosts() ([]*entities.Post, error) {
	posts := []*entities.Post{}

	rows, err := p.conn.Query(
		"SELECT * FROM get_posts()",
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return posts, nil
		}

		return nil, err
	}

	for rows.Next() {
		var post entities.Post
		if err := rows.Scan(&post.ID, &post.Name, &post.Text, &post.CreateDate, &post.UpdateDate, &post.UserId, &post.UserName); err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}

	return posts, nil
}

func (p *Provider) SelectUserPosts(userId int) ([]*entities.Post, error) {
	posts := []*entities.Post{}

	rows, err := p.conn.Query(
		"SELECT * FROM get_user_posts($1)",
		userId,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return posts, nil
		}

		return nil, err
	}

	for rows.Next() {
		var post entities.Post
		if err := rows.Scan(&post.ID, &post.Name, &post.Text, &post.CreateDate, &post.UpdateDate, &post.UserId, &post.UserName); err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}

	return posts, nil
}

// редактировние поста
func (p *Provider) UpdatePostById(post entities.Post) (*entities.Post, error) {
	var updatedPost entities.Post
	_, err := p.conn.Query(
		"CALL update_post($1, $2, $3, $4)",
		post.ID, post.Name, post.Text, post.UpdateDate,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, entities.ErrPostNotFound
		}

		return nil, err
	}

	return &updatedPost, nil
}

// удаление поста
func (p *Provider) DeletePostById(id int) error {
	_, err := p.conn.Exec("DELETE FROM posts WHERE id = $1", id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entities.ErrPostNotFound
		}

		return err
	}

	return nil
}
