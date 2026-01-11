package repository

import (
	"database/sql"

	"myapp-backend/domain"
)

type IAuditLogRepository interface {
	Insert(log domain.AuditLog) error
}

type AuditLogRepository struct {
	db *sql.DB
}

func NewAuditLogRepository(db *sql.DB) IAuditLogRepository {
	return &AuditLogRepository{db: db}
}

func (r *AuditLogRepository) Insert(log domain.AuditLog) error {
	query := `
		INSERT INTO audit_logs (entity_type, entity_id, action, details, created_at)
		VALUES (?, ?, ?, ?, ?)
	`
	_, err := r.db.Exec(query, log.EntityType, log.EntityID, log.Action, log.Details, log.CreatedAt)
	return err
}
