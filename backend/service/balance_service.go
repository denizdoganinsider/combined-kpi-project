package service

import (
	"errors"
	"log"
	"sync"

	"myapp-backend/domain"
	"myapp-backend/repository"
)

type IBalanceService interface {
	GetBalanceByUserID(userID int64) (*domain.Balance, error)
	UpdateBalance(userID int64, amount float64) error
	CreateBalance(userID int64, amount float64) error
}

type BalanceService struct {
	balanceRepository repository.IBalanceRepository
	mu                sync.Mutex
}

func NewBalanceService(balanceRepository repository.IBalanceRepository) IBalanceService {
	return &BalanceService{
		balanceRepository: balanceRepository,
	}
}

func (balanceService *BalanceService) GetBalanceByUserID(userID int64) (*domain.Balance, error) {
	// Read-only → lock gerekmez (istersen RWMutex ile optimize edilebilir)
	balance, err := balanceService.balanceRepository.GetBalanceByUserID(userID)
	if err != nil {
		return nil, err
	}
	return balance, nil
}

func (balanceService *BalanceService) UpdateBalance(userID int64, amount float64) error {
	balanceService.mu.Lock()
	defer balanceService.mu.Unlock()

	balance, err := balanceService.balanceRepository.GetBalanceByUserID(userID)
	if err != nil && err.Error() != "user doesn't have balance" {
		return err
	}

	// Kullanıcının bakiyesi yoksa
	if balance == nil {
		if amount < 0 {
			return errors.New("insufficient balance")
		}

		err = balanceService.balanceRepository.CreateBalance(userID, amount)
		if err != nil {
			return err
		}

		log.Printf("New balance created for user %d with amount %f", userID, amount)
		return nil
	}

	newAmount := balance.Amount + amount
	if newAmount < 0 {
		return errors.New("insufficient balance")
	}

	err = balanceService.balanceRepository.UpdateBalance(userID, newAmount)
	if err != nil {
		return err
	}

	log.Printf("Balance for user %d updated to %f", userID, newAmount)
	return nil
}

func (balanceService *BalanceService) CreateBalance(userID int64, amount float64) error {
	balanceService.mu.Lock()
	defer balanceService.mu.Unlock()

	err := balanceService.balanceRepository.CreateBalance(userID, amount)
	if err != nil {
		return err
	}

	log.Printf("New balance created for user %d with amount %f", userID, amount)
	return nil
}
