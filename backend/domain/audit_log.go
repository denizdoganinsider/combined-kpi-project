package domain

import "time"

type AuditAction string

const (
	AuditUserCreated AuditAction = "USER_CREATED"
	AuditUserLogin   AuditAction = "USER_LOGIN"
	AuditCredit      AuditAction = "BALANCE_CREDIT"
	AuditDebit       AuditAction = "BALANCE_DEBIT"
	AuditTransfer    AuditAction = "MONEY_TRANSFER"
	AuditRollback    AuditAction = "TRANSACTION_ROLLBACK"
)

type AuditLog struct {
	ID         int64
	EntityType string
	EntityID   int64
	Action     AuditAction
	Details    string
	CreatedAt  time.Time
}
