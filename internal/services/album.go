package services

import (
	"context"
	"time"
)

type Album struct {
	ID     		string `json:"id"`
	Title  		string `json:"title"`
	Description string `json:"description"`
	CreatedAt 	time.Time `json:"created_at"`
	UpdatedAt 	time.Time `json:"updated_at"`
}

type AlbumPayload struct {
	Title  		string `json:"title"`
	Description string `json:"description"`
}

type AlbumsList struct {
	Albums []Album `json:"albums"`
}
type UserAlbum struct {
	UserID string `json:"user_id"`
	AlbumID string `json:"album_id"`
	AddedAt time.Time `json:"added_at"`
}

type UserAlbumPayload struct {
	UserID string `json:"user_id"`
	AlbumID string `json:"album_id"`
}

type UserAlbumsList struct {
	UserAlbums []UserAlbum `json:"user_albums"`
}

func (a *Album) GetAllAlbums() ([]*Album, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `SELECT id, title, description, created_at, updated_at FROM albums`
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	var albums []*Album
	for rows.Next() {
		var album Album
		err := rows.Scan(
			&album.ID,
			&album.Title,
			&album.Description,
			&album.CreatedAt,
			&album.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		albums = append(albums, &album)
	}
	return albums, nil
}

func (a *Album) GetAlbumByID(id string) (*Album, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `SELECT id, title, description, created_at, updated_at FROM albums WHERE id = $1`
	row := db.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&a.ID,
		&a.Title,
		&a.Description,
		&a.CreatedAt,
		&a.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (a *Album) CreateAlbum(album Album) (*Album, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `INSERT INTO albums (title, description, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id, title, description`
	err := db.QueryRowContext(ctx, query, album.Title, album.Description, time.Now(), time.Now()).Scan(&album.ID, &album.Title, &album.Description)
	if err != nil {
		return nil, err
	}
	return &album, nil
}

func (a *Album) UpdateAlbum(id string, album Album) (*Album, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `UPDATE albums SET title = $1, description = $2, updated_at = $3 WHERE id = $4 RETURNING id, title, description, updated_at`
	row := db.QueryRowContext(ctx, query, album.Title, album.Description, time.Now(), id)
	err := row.Scan(&album.ID, &album.Title, &album.Description, &album.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &album, nil
}

func (a *Album) DeleteAlbum(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `DELETE FROM albums WHERE id = $1`
	_, err := db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (a *Album) GetUserAlbums(userID string) ([]*Album, error) {
    ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
    defer cancel()

    query := `
        SELECT a.id, a.title, a.description, a.created_at, a.updated_at 
        FROM albums a
        JOIN user_albums ua ON a.id = ua.album_id
        WHERE ua.user_id = $1
    `

    rows, err := db.QueryContext(ctx, query, userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var albums []*Album
    for rows.Next() {
        var album Album
        err := rows.Scan(
            &album.ID,
            &album.Title,
            &album.Description,
            &album.CreatedAt,
            &album.UpdatedAt,
        )
        if err != nil {
            return nil, err
        }
        albums = append(albums, &album)
    }

    if err = rows.Err(); err != nil {
        return nil, err
    }

    return albums, nil
}

func (a *Album) AddAlbumToUser(userAlbum UserAlbum) (*UserAlbum, error) {
    ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
    defer cancel()

    query := `INSERT INTO user_albums (user_id, album_id, added_at) VALUES ($1, $2, $3)`

    _, err := db.ExecContext(ctx, query, userAlbum.UserID, userAlbum.AlbumID, time.Now())
    if err != nil {
        return nil,err
    }

    return &userAlbum, nil
}

func (a *Album) RemoveAlbumFromUser(userID string, albumID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `DELETE FROM user_albums WHERE user_id = $1 AND album_id = $2`

	_, err := db.ExecContext(ctx, query, userID, albumID)
	if err != nil {
		return err
	}

	return nil
}