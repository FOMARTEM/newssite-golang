package provider

import (
  "database/sql"
  "errors"

  "github.com/FOMARTEM/newssite-golang/internal/entities"
)

func (p *Provider) InsertComment(comment entities.Comment) (*entities.Comment, error) {
  var id int

  err := p.conn.QueryRow(
    `CALL create_comment($1, $2, $3, n_id := NULL)`,
    comment.CommentText, comment.PostId, comment.UserId,
  ).Scan(&id)

  if err != nil {
    return nil, err
  }

  comment.ID = id
  return &comment, nil
}

unc (p *Provider) GetCommentsForPost(post_id int, limit int, offset int) ([]*entities.Comment, error) {
  comments := []*entities.Comment{}

  rows, err := p.conn.Query(
    "SELECT * FROM get_comments_by_post_id($1, $2, $3)",
    post_id,
    limit,
    offset,
  )

  if err != nil {
    if errors.Is(err, sql.ErrNoRows) {
      return comments, nil
    }

    return nil, err
  }

  for rows.Next() {
    var comment entities.Comment

    err := rows.Scan(&comment.ID, &comment.CommentText, &comment.PostId, &comment.UserId, &comment.UserName)

    if err != nil {
      return nil, err
    }
    comments = append(comments, &comment)
  }

  return comments, nil
}

func (p *Provider) UpdateComment(comment entities.Comment) (*entities.Comment, error) {
  _, err := p.conn.Query(
    "CALL update_comment($1, $2)",
    comment.ID, comment.CommentText,
  )

  if err != nil {
    if errors.Is(err, sql.ErrNoRows) {
      return nil, entities.ErrCommentNotFound
    }

    return nil, err
  }

  return &comment, nil
}

func (p *Provider) DeleteComment(id int) error {
  _, err := p.conn.Query(
    "CALL  delete_comment($1)",
    id,
  )

  if err != nil {
    return err
  }

  return nil
}

func (p *Provider) DeleteCommentsInPost(postId int) error {
  _, err := p.conn.Query(
    "CALL  delete_comments($1)",
    postId,
  )

  if err != nil {
    return err
  }

  return nil
}
