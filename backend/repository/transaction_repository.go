package repository

import (
	"database/sql"
	"errors"

	"myapp-backend/domain"
)

type ITransactionRepository interface {
	BeginTx() (*sql.Tx, error)

	CreateTransaction(transaction *domain.Transaction, tx *sql.Tx) error
	GetTransactionByID(id int64) (*domain.Transaction, error)
	UpdateTransactionStatus(id int64, status domain.TransactionStatus, tx *sql.Tx) error
	GetUserTransactions(userID int64) ([]domain.Transaction, error)
}

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) ITransactionRepository {
	return &TransactionRepository{db: db}
}

/* ---------- TX ---------- */

func (r *TransactionRepository) BeginTx() (*sql.Tx, error) {
	return r.db.Begin()
}

/* ---------- WRITE ---------- */

func (r *TransactionRepository) CreateTransaction(t *domain.Transaction, tx *sql.Tx) error {
	query := `
		INSERT INTO transactions (from_user_id, to_user_id, amount, type, status, created_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	res, err := tx.Exec(query, t.FromUser, t.ToUser, t.Amount, t.Type, t.Status, t.CreatedAt)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	t.ID = id
	return nil
}

func (r *TransactionRepository) UpdateTransactionStatus(id int64, status domain.TransactionStatus, tx *sql.Tx) error {
	if id == 0 {
		return errors.New("transaction id is zero")
	}

	_, err := tx.Exec(`UPDATE transactions SET status = ? WHERE id = ?`, status, id)
	return err
}

/* ---------- READ ---------- */

func (r *TransactionRepository) GetTransactionByID(id int64) (*domain.Transaction, error) {
	row := r.db.QueryRow(`
		SELECT id, from_user_id, to_user_id, amount, type, status, created_at
		FROM transactions
		WHERE id = ?
	`, id)

	var t domain.Transaction
	var toUser sql.NullInt64

	err := row.Scan(&t.ID, &t.FromUser, &toUser, &t.Amount, &t.Type, &t.Status, &t.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	if toUser.Valid {
		t.ToUser = &toUser.Int64
	}

	return &t, nil
}

func (r *TransactionRepository) GetUserTransactions(userID int64) ([]domain.Transaction, error) {
	rows, err := r.db.Query(`
		SELECT id, from_user_id, to_user_id, amount, type, status, created_at
		FROM transactions
		WHERE from_user_id = ? OR to_user_id = ?
		ORDER BY created_at DESC
	`, userID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []domain.Transaction

	for rows.Next() {
		var t domain.Transaction
		var toUser sql.NullInt64

		if err := rows.Scan(&t.ID, &t.FromUser, &toUser, &t.Amount, &t.Type, &t.Status, &t.CreatedAt); err != nil {
			return nil, err
		}

		if toUser.Valid {
			t.ToUser = &toUser.Int64
		}

		list = append(list, t)
	}

	return list, nil
}
