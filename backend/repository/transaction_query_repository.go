package repository

import (
	"database/sql"
	"fmt"
	"strings"

	"myapp-backend/domain"
)

type ITransactionQueryRepository interface {
	GetTransactionHistory(filter domain.TransactionFilter) ([]domain.Transaction, int, error)
}

type TransactionQueryRepository struct {
	db *sql.DB
}

func NewTransactionQueryRepository(db *sql.DB) ITransactionQueryRepository {
	return &TransactionQueryRepository{db: db}
}

func (r *TransactionQueryRepository) GetTransactionHistory(filter domain.TransactionFilter) ([]domain.Transaction, int, error) {
	where := []string{"(from_user_id = ? OR to_user_id = ?)"}
	args := []interface{}{filter.UserID, filter.UserID}

	if filter.FromTime != nil {
		where = append(where, "created_at >= ?")
		args = append(args, *filter.FromTime)
	}
	if filter.ToTime != nil {
		where = append(where, "created_at <= ?")
		args = append(args, *filter.ToTime)
	}
	if len(filter.Types) > 0 {
		placeholders := make([]string, len(filter.Types))
		for i, t := range filter.Types {
			placeholders[i] = "?"
			args = append(args, t)
		}
		where = append(where, fmt.Sprintf("type IN (%s)", strings.Join(placeholders, ",")))
	}
	if len(filter.Statuses) > 0 {
		placeholders := make([]string, len(filter.Statuses))
		for i, s := range filter.Statuses {
			placeholders[i] = "?"
			args = append(args, s)
		}
		where = append(where, fmt.Sprintf("status IN (%s)", strings.Join(placeholders, ",")))
	}
	if filter.MinAmount != nil {
		where = append(where, "amount >= ?")
		args = append(args, *filter.MinAmount)
	}
	if filter.MaxAmount != nil {
		where = append(where, "amount <= ?")
		args = append(args, *filter.MaxAmount)
	}

	whereSQL := "WHERE " + strings.Join(where, " AND ")

	sortBy := "created_at"
	if filter.SortBy != "" {
		sortBy = filter.SortBy
	}

	order := "DESC"
	if filter.Order == domain.SortAsc {
		order = "ASC"
	}

	limit := filter.Limit
	if limit <= 0 {
		limit = 20
	}

	query := fmt.Sprintf(`
		SELECT id, from_user_id, to_user_id, amount, type, status, created_at
		FROM transactions
		%s
		ORDER BY %s %s
		LIMIT ? OFFSET ?
	`, whereSQL, sortBy, order)

	argsWithPaging := append(args, limit, filter.Offset())

	rows, err := r.db.Query(query, argsWithPaging...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var transactions []domain.Transaction

	for rows.Next() {
		var tx domain.Transaction
		var toUser sql.NullInt64
		err := rows.Scan(&tx.ID, &tx.FromUser, &toUser, &tx.Amount, &tx.Type, &tx.Status, &tx.CreatedAt)
		if err != nil {
			return nil, 0, err
		}
		if toUser.Valid {
			tx.ToUser = &toUser.Int64
		}
		transactions = append(transactions, tx)
	}

	countQuery := fmt.Sprintf(`SELECT COUNT(*) FROM transactions %s`, whereSQL)
	row := r.db.QueryRow(countQuery, args...)
	var total int
	if err := row.Scan(&total); err != nil {
		return nil, 0, err
	}

	return transactions, total, nil
}
