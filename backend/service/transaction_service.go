package service

import (
	"errors"
	"time"

	"myapp-backend/domain"
	"myapp-backend/repository"
)

type ITransactionService interface {
	Credit(userID int64, amount float64) (*domain.Transaction, error)
	Debit(userID int64, amount float64) (*domain.Transaction, error)
	Transfer(fromUserID int64, toUserID int64, amount float64) (*domain.Transaction, error)
	GetTransactionHistory(userID int64) ([]domain.Transaction, error)
	GetTransactionByID(transactionID int64) (*domain.Transaction, error)
}

type TransactionService struct {
	transactionRepository repository.ITransactionRepository
	balanceRepo           repository.IBalanceRepository
}

func NewTransactionService(
	transactionRepository repository.ITransactionRepository,
	balanceRepo repository.IBalanceRepository,
) ITransactionService {
	return &TransactionService{
		transactionRepository: transactionRepository,
		balanceRepo:           balanceRepo,
	}
}

/* ---------------- CREDIT ---------------- */

func (s *TransactionService) Credit(userID int64, amount float64) (*domain.Transaction, error) {
	if amount <= 0 {
		return nil, errors.New("amount must be greater than zero")
	}

	dbTx, err := s.transactionRepository.BeginTx()
	if err != nil {
		return nil, err
	}

	defer dbTx.Rollback()

	tx := &domain.Transaction{
		FromUser:  userID,
		Amount:    amount,
		Type:      domain.CreditTransaction,
		Status:    domain.Pending,
		CreatedAt: time.Now(),
	}

	if err := s.transactionRepository.CreateTransaction(tx, dbTx); err != nil {
		return nil, err
	}

	if err := s.balanceRepo.UpdateBalance(userID, amount); err != nil {
		s.transactionRepository.UpdateTransactionStatus(tx.ID, domain.Failed, dbTx)
		return nil, err
	}

	if err := s.transactionRepository.UpdateTransactionStatus(tx.ID, domain.Completed, dbTx); err != nil {
		return nil, err
	}

	if err := dbTx.Commit(); err != nil {
		return nil, err
	}

	return tx, nil
}

/* ---------------- DEBIT ---------------- */

func (s *TransactionService) Debit(userID int64, amount float64) (*domain.Transaction, error) {
	if amount <= 0 {
		return nil, errors.New("amount must be greater than zero")
	}

	dbTx, err := s.transactionRepository.BeginTx()
	if err != nil {
		return nil, err
	}
	defer dbTx.Rollback()

	tx := &domain.Transaction{
		FromUser:  userID,
		Amount:    -amount,
		Type:      domain.DebitTransaction,
		Status:    domain.Pending,
		CreatedAt: time.Now(),
	}

	if err := s.transactionRepository.CreateTransaction(tx, dbTx); err != nil {
		return nil, err
	}

	if err := s.balanceRepo.UpdateBalance(userID, -amount); err != nil {
		s.transactionRepository.UpdateTransactionStatus(tx.ID, domain.Failed, dbTx)
		return nil, err
	}

	if err := s.transactionRepository.UpdateTransactionStatus(tx.ID, domain.Completed, dbTx); err != nil {
		return nil, err
	}

	if err := dbTx.Commit(); err != nil {
		return nil, err
	}

	return tx, nil
}

/* ---------------- TRANSFER ---------------- */

func (s *TransactionService) Transfer(fromUserID, toUserID int64, amount float64) (*domain.Transaction, error) {
	if amount <= 0 {
		return nil, errors.New("amount must be greater than zero")
	}
	if fromUserID == toUserID {
		return nil, errors.New("cannot transfer to same user")
	}

	dbTx, err := s.transactionRepository.BeginTx()
	if err != nil {
		return nil, err
	}
	defer dbTx.Rollback()

	tx := &domain.Transaction{
		FromUser:  fromUserID,
		ToUser:    &toUserID,
		Amount:    amount,
		Type:      domain.TransferTransaction,
		Status:    domain.Pending,
		CreatedAt: time.Now(),
	}

	if err := s.transactionRepository.CreateTransaction(tx, dbTx); err != nil {
		return nil, err
	}

	if err := s.balanceRepo.UpdateBalance(fromUserID, -amount); err != nil {
		s.transactionRepository.UpdateTransactionStatus(tx.ID, domain.Failed, dbTx)
		return nil, err
	}

	if err := s.balanceRepo.UpdateBalance(toUserID, amount); err != nil {
		s.transactionRepository.UpdateTransactionStatus(tx.ID, domain.Failed, dbTx)
		return nil, err
	}

	if err := s.transactionRepository.UpdateTransactionStatus(tx.ID, domain.Completed, dbTx); err != nil {
		return nil, err
	}

	if err := dbTx.Commit(); err != nil {
		return nil, err
	}

	return tx, nil
}

/* ---------------- READ ---------------- */

func (s *TransactionService) GetTransactionHistory(userID int64) ([]domain.Transaction, error) {
	return s.transactionRepository.GetUserTransactions(userID)
}

func (s *TransactionService) GetTransactionByID(id int64) (*domain.Transaction, error) {
	return s.transactionRepository.GetTransactionByID(id)
}
