package repository

import (
	"database/sql"
	"errors"
	"myapp-backend/domain"
)

var ErrNotFound = errors.New("not found")

type FavoriteRepository interface {
	Add(f *domain.Favorite) (*domain.Favorite, error)
	Remove(userID, movieID int64) error
	ListByUser(userID int64, limit, offset int) ([]*domain.Favorite, error)
	Exists(userID, movieID int64) (bool, error)
}

type favoriteRepo struct {
	db *sql.DB
}

func NewFavoriteRepository(db *sql.DB) FavoriteRepository {
	return &favoriteRepo{db: db}
}

func (r *favoriteRepo) Add(f *domain.Favorite) (*domain.Favorite, error) {
	res, err := r.db.Exec(`
        INSERT INTO favorites (user_id, movie_id)
        VALUES (?, ?)
        ON DUPLICATE KEY UPDATE id = id
    `, f.UserID, f.MovieID)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		row := r.db.QueryRow("SELECT id, created_at FROM favorites WHERE user_id = ? AND movie_id = ?", f.UserID, f.MovieID)
		var existingID int64
		var createdAt sql.NullTime
		if err2 := row.Scan(&existingID, &createdAt); err2 != nil {
			return nil, err2
		}
		f.ID = existingID
		if createdAt.Valid {
			f.CreatedAt = createdAt.Time
		}
		return f, nil
	}
	f.ID = id
	row := r.db.QueryRow("SELECT created_at FROM favorites WHERE id = ?", f.ID)
	var createdAt sql.NullTime
	if err := row.Scan(&createdAt); err == nil && createdAt.Valid {
		f.CreatedAt = createdAt.Time
	}
	return f, nil
}

func (r *favoriteRepo) Remove(userID, movieID int64) error {
	res, err := r.db.Exec("DELETE FROM favorites WHERE user_id = ? AND movie_id = ?", userID, movieID)
	if err != nil {
		return err
	}
	n, _ := res.RowsAffected()
	if n == 0 {
		return ErrNotFound
	}
	return nil
}

func (r *favoriteRepo) ListByUser(userID int64, limit, offset int) ([]*domain.Favorite, error) {
	rows, err := r.db.Query("SELECT id, user_id, movie_id, created_at FROM favorites WHERE user_id = ? ORDER BY created_at DESC LIMIT ? OFFSET ?", userID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []*domain.Favorite
	for rows.Next() {
		var f domain.Favorite
		if err := rows.Scan(&f.ID, &f.UserID, &f.MovieID, &f.CreatedAt); err != nil {
			return nil, err
		}
		out = append(out, &f)
	}
	return out, nil
}

func (r *favoriteRepo) Exists(userID, movieID int64) (bool, error) {
	var cnt int
	err := r.db.QueryRow("SELECT 1 FROM favorites WHERE user_id = ? AND movie_id = ? LIMIT 1", userID, movieID).Scan(&cnt)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
