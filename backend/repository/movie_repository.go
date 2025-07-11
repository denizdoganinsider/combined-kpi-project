package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"myapp-backend/domain"
)

type MovieRepository struct {
	db *sql.DB
}

func NewMovieRepository(db *sql.DB) *MovieRepository {
	return &MovieRepository{db: db}
}

func (mr *MovieRepository) CreateMovie(movie domain.Movie) error {
	_, err := mr.db.Exec(
		"INSERT INTO movies(title, director, year, genre) VALUES (?, ?, ?, ?)",
		movie.Title, movie.Director, movie.Year, movie.Genre,
	)

	if err != nil {
		return fmt.Errorf("movie not added to db: %w", err)
	}
	return nil
}

func (mr *MovieRepository) FindByTitle(title string) (*domain.Movie, error) {
	query := "SELECT title, director, year, genre FROM movies WHERE title = ?"

	row := mr.db.QueryRow(query, title)

	var movie domain.Movie
	err := row.Scan(&movie.Title, &movie.Director, &movie.Year, &movie.Genre)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.New("error occurred while getting movie data")
	}

	return &movie, nil
}
