package domain

import "time"

type Favorite struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	MovieID   int64     `json:"movie_id"`
	CreatedAt time.Time `json:"created_at"`
}
