package services

import (
	"context"
	"time"
)

type Post struct {
	ID     		string `json:"id"`
	UserID 		string `json:"user_id"`
	Content		string `json:"content"`
	CreatedAt 	time.Time `json:"created_at"`
	UpdatedAt 	time.Time `json:"updated_at"`
}

type PostPayload struct {
	UserID 		string `json:"user_id"`
	Content		string `json:"content"`
}

type PostsList struct {
	Posts []Post `json:"posts"`
}


func (p *Post) GetAllPosts() ([]*Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `SELECT id, user_id, content, created_at, updated_at FROM posts`
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	var posts []*Post
	for rows.Next() {
		var post Post
		err := rows.Scan(
			&post.ID,
			&post.UserID,
			&post.Content,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}
	return posts, nil
}

func (p *Post) GetPostByID(id string) (*Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `SELECT id, user_id, content, created_at, updated_at FROM posts WHERE id = $1`
	row := db.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&p.ID,
		&p.UserID,
		&p.Content,
		&p.CreatedAt,
		&p.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (p *Post) CreatePost(post Post) (*Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `INSERT INTO posts (user_id, content, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id`
	err := db.QueryRowContext(ctx, query, post.UserID, post.Content, time.Now(), time.Now()).Scan(&post.ID)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (p *Post) UpdatePost(id string, post Post) (*Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `UPDATE posts SET content = $1, updated_at = $2 WHERE id = $3 RETURNING id, user_id, content, created_at, updated_at`
	err := db.QueryRowContext(ctx, query, post.Content, time.Now(), id).Scan(&post.ID, &post.UserID, &post.Content, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (p *Post) DeletePost(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `DELETE FROM posts WHERE id = $1`
	_, err := db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (p *Post) GetPostsByUserID(id string) ([]*Post, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `SELECT id, user_id, content, created_at, updated_at FROM posts WHERE user_id = $1`
	rows, err := db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	var posts []*Post
	for rows.Next() {
		var post Post
		err := rows.Scan(
			&post.ID,
			&post.UserID,
			&post.Content,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}
	return posts, nil
}