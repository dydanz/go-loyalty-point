package domain

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	TransactionID       uuid.UUID  `json:"transaction_id"`
	MerchantID          uuid.UUID  `json:"merchant_id"`
	MerchantCustomersID uuid.UUID  `json:"merchant_customers_id"`
	ProgramID           uuid.UUID  `json:"program_id"`
	TransactionType     string     `json:"transaction_type"` // purchase, refund, bonus
	TransactionAmount   float64    `json:"transaction_amount"`
	TransactionDate     time.Time  `json:"transaction_date"`
	BranchID            *uuid.UUID `json:"branch_id,omitempty"`
	Status              string     `json:"status"`
	CreatedAt           time.Time  `json:"created_at"`

	// Rule-evaluation context. Not persisted (no transactions table column);
	// populated by the caller before running the rewards engine so program
	// rules of these condition types (program_rule_transaction_category,
	// _merchant_group, _transaction_count, _tenure) can match. Zero value
	// when unset, so existing JSON responses are unaffected.
	Category         string `json:"category,omitempty"`
	MerchantGroupID  string `json:"merchant_group_id,omitempty"`
	TransactionCount int    `json:"transaction_count,omitempty"`
	MembershipTenure int    `json:"membership_tenure,omitempty"` // in days
}

type CreateTransactionRequest struct {
	MerchantID          uuid.UUID  `json:"merchant_id" binding:"required"`
	MerchantCustomersID uuid.UUID  `json:"merchant_customers_id" binding:"required"`
	ProgramID           uuid.UUID  `json:"program_id" binding:"required"`
	TransactionType     string     `json:"transaction_type" binding:"required,oneof=purchase refund bonus"`
	TransactionAmount   float64    `json:"transaction_amount" binding:"required,gt=0"`
	TransactionDate     time.Time  `json:"transaction_date" binding:"required"`
	BranchID            *uuid.UUID `json:"branch_id,omitempty"`
	Status              string     `json:"status" binding:"required,oneof=pending completed failed cancelled"`
}

type UpdateTransactionStatusRequest struct {
	Status string `json:"status" binding:"required,oneof=pending completed failed cancelled"`
}
