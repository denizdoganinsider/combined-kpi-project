package service

import (
	"errors"
	"myapp-backend/controller/request"
	"myapp-backend/domain"
	"myapp-backend/repository"
)

type MovieService struct {
	movieRepository *repository.MovieRepository
}

func NewMovieService(movieRepository *repository.MovieRepository) *MovieService {
	return &MovieService{
		movieRepository: movieRepository,
	}
}

func (ms *MovieService) CreateMovie(movieRequest request.Movie) error {
	moviePointer, err := ms.movieRepository.FindByTitle(movieRequest.Title)
	if err != nil {
		return err
	}

	if moviePointer != nil {
		return errors.New("movie with this title already exists")
	}

	return ms.movieRepository.CreateMovie(domain.Movie{
		Title:    movieRequest.Title,
		Director: movieRequest.Director,
		Year:     movieRequest.Year,
		Genre:    movieRequest.Genre,
	})
}
