package domain

import (
	"errors"
	"time"
)

type TransactionStatus string
type TransactionType string
type SortOrder string

const (
	CreditTransaction   TransactionType   = "credit"
	DebitTransaction    TransactionType   = "debit"
	TransferTransaction TransactionType   = "transfer"
	SortAsc             SortOrder         = "asc"
	SortDesc            SortOrder         = "desc"
	Pending             TransactionStatus = "pending"
	Completed           TransactionStatus = "completed"
	Failed              TransactionStatus = "failed"
)

type Transaction struct {
	ID        int64
	FromUser  int64
	ToUser    *int64
	Amount    float64
	Type      TransactionType
	Status    TransactionStatus
	CreatedAt time.Time
}

type TransactionFilter struct {
	UserID int64

	FromTime *time.Time
	ToTime   *time.Time

	Types    []TransactionType
	Statuses []TransactionStatus

	MinAmount *float64
	MaxAmount *float64

	SortBy string
	Order  SortOrder

	Page  int
	Limit int
}

func (f *TransactionFilter) Offset() int {
	if f.Page <= 1 {
		return 0
	}
	return (f.Page - 1) * f.Limit
}

func (t *Transaction) Validate() error {
	if t.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	if t.Type == TransferTransaction && (t.ToUser == nil || *t.ToUser == t.FromUser) {
		return errors.New("invalid transfer transaction")
	}
	return nil
}
