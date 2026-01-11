package service

import (
	"myapp-backend/domain"
	"myapp-backend/repository"
)

type ITransactionQueryService interface {
	GetHistory(filter domain.TransactionFilter) ([]domain.Transaction, int, error)
}

type TransactionQueryService struct {
	repo repository.ITransactionQueryRepository
}

func NewTransactionQueryService(repo repository.ITransactionQueryRepository) ITransactionQueryService {
	return &TransactionQueryService{repo: repo}
}

func (s *TransactionQueryService) GetHistory(filter domain.TransactionFilter) ([]domain.Transaction, int, error) {
	return s.repo.GetTransactionHistory(filter)
}
