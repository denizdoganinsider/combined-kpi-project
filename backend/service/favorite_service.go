package service

import (
	"errors"
	"myapp-backend/domain"
	"myapp-backend/repository"
)

var (
	ErrAlreadyFavorited = errors.New("already favorited")
	ErrFavoriteNotFound = errors.New("favorite not found")
)

type FavoriteService interface {
	AddFavorite(userID, movieID int64) (*domain.Favorite, error)
	RemoveFavorite(userID, movieID int64) error
	ListFavorites(userID int64, limit, offset int) ([]*domain.Favorite, error)
	IsFavorited(userID, movieID int64) (bool, error)
}

type favoriteService struct {
	favoriteService repository.FavoriteRepository
}

func NewFavoriteService(r repository.FavoriteRepository) FavoriteService {
	return &favoriteService{favoriteService: r}
}

func (s *favoriteService) AddFavorite(userID, movieID int64) (*domain.Favorite, error) {
	exists, err := s.favoriteService.Exists(userID, movieID)
	if err != nil {
		return nil, err
	}
	if exists {
		return &domain.Favorite{UserID: userID, MovieID: movieID}, ErrAlreadyFavorited
	}
	f := &domain.Favorite{
		UserID:  userID,
		MovieID: movieID,
	}
	return s.favoriteService.Add(f)
}

func (s *favoriteService) RemoveFavorite(userID, movieID int64) error {
	if err := s.favoriteService.Remove(userID, movieID); err != nil {
		if err == repository.ErrNotFound {
			return ErrFavoriteNotFound
		}
		return err
	}
	return nil
}

func (s *favoriteService) ListFavorites(userID int64, limit, offset int) ([]*domain.Favorite, error) {
	return s.favoriteService.ListByUser(userID, limit, offset)
}

func (s *favoriteService) IsFavorited(userID, movieID int64) (bool, error) {
	return s.favoriteService.Exists(userID, movieID)
}
