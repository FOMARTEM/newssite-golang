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
		`INSERT INTO posts (title, body, createdate, updatedate, "user_id") VALUES ($1, $2, $3, $4, $5)  RETURNING "post_id"`,
		post.Name, post.Text, post.CreateDate, post.CreateDate, post.UserId,
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
		"SELECT \"post_id\", title, body, TO_CHAR(createdate, 'YYYY/MM/DD') AS createdate,  TO_CHAR(updatedate, 'YYYY/MM/DD') AS updatedate,  \"user_id\"  FROM public.posts WHERE id = $1",
		id,
	).Scan(&post.ID, &post.Name, &post.Text, &post.CreateDate, &post.UpdateDate, &post.UserId)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

// получение всех постов
func (p *Provider) SelectAllPosts() ([]*entities.Post, error) {
	posts := []*entities.Post{}

	rows, err := p.conn.Query(
		"SELECT \"post_id\", title, body, TO_CHAR(createdate, 'YYYY/MM/DD') AS createdate,  TO_CHAR(updatedate, 'YYYY/MM/DD') AS updatedate, \"user_id\"  FROM public.posts ORDER BY \"post_id\" ASC",
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return posts, nil
		}

		return nil, err
	}

	for rows.Next() {
		var post entities.Post
		if err := rows.Scan(&post.ID, &post.Name, &post.Text, &post.CreateDate, &post.UpdateDate, &post.UserId); err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}

	return posts, nil
}

// редактировние поста
func (p *Provider) UpdatePostById(post entities.Post) (*entities.Post, error) {
	var updatedPost entities.Post
	err := p.conn.QueryRow(
		"UPDATE public.posts SET title=$1, body=$2, updatedate=$3 WHERE id = $4 RETURNING title, body, createdate, updatedate, userid",
		post.Name, post.Text, post.UpdateDate, post.ID,
	).Scan(&updatedPost.Name, updatedPost.Text, updatedPost.CreateDate, updatedPost.UpdateDate, updatedPost.UserId)
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
