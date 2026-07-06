package service

import (
	"testing"
	"time"

	"go-playground/server/domain"

	"github.com/google/uuid"
)

func TestCalculatePoints(t *testing.T) {
	// Helper function to create a time pointer
	timePtr := func(t time.Time) *time.Time {
		return &t
	}

	// Base time for testing
	now := time.Now()
	yesterday := now.AddDate(0, 0, -1)
	tomorrow := now.AddDate(0, 0, 1)

	tests := []struct {
		name     string
		rules    []*domain.ProgramRule
		tx       domain.Transaction
		expected float64
	}{
		{
			name: "Transaction-Based Points - Minimum Spend",
			rules: []*domain.ProgramRule{
				{
					RuleName:       "Spend $100 get 10 points",
					ConditionType:  "program_rule_transaction_amount",
					ConditionValue: "100",
					Multiplier:     1.0,
					PointsAwarded:  10,
					EffectiveFrom:  yesterday,
					EffectiveTo:    timePtr(tomorrow),
				},
			},
			tx: domain.Transaction{
				TransactionAmount: 150.0,
			},
			expected: 10.0,
		},
		{
			name: "Transaction-Based Points - Category Multiplier",
			rules: []*domain.ProgramRule{
				{
					RuleName:       "2x points on dining",
					ConditionType:  "program_rule_transaction_category",
					ConditionValue: "dining",
					Multiplier:     2.0,
					PointsAwarded:  100,
					EffectiveFrom:  yesterday,
					EffectiveTo:    timePtr(tomorrow),
				},
			},
			tx: domain.Transaction{
				TransactionAmount: 100.0,
				Category:          "dining",
			},
			expected: 200.0,
		},
		{
			name: "Frequency-Based Points - Transaction Count",
			rules: []*domain.ProgramRule{
				{
					RuleName:       "Bonus for 5+ transactions",
					ConditionType:  "program_rule_transaction_count",
					ConditionValue: "5",
					Multiplier:     1.0,
					PointsAwarded:  500,
					EffectiveFrom:  yesterday,
					EffectiveTo:    timePtr(tomorrow),
				},
			},
			tx: domain.Transaction{
				TransactionAmount: 100.0,
				TransactionCount:  6,
			},
			expected: 500.0,
		},
		{
			name: "Loyalty Milestone - Membership Tenure",
			rules: []*domain.ProgramRule{
				{
					RuleName:       "1 year membership bonus",
					ConditionType:  "program_rule_tenure",
					ConditionValue: "365",
					Multiplier:     1.0,
					PointsAwarded:  1000,
					EffectiveFrom:  yesterday,
					EffectiveTo:    timePtr(tomorrow),
				},
			},
			tx: domain.Transaction{
				TransactionAmount: 100.0,
				MembershipTenure:  366,
			},
			expected: 1000.0,
		},
		{
			name: "Category-Based Points - Merchant Group",
			rules: []*domain.ProgramRule{
				{
					RuleName:       "Partner merchant group bonus",
					ConditionType:  "program_rule_transaction_merchant_group",
					ConditionValue: "premium_partners",
					Multiplier:     3.0,
					PointsAwarded:  100,
					EffectiveFrom:  yesterday,
					EffectiveTo:    timePtr(tomorrow),
				},
			},
			tx: domain.Transaction{
				TransactionAmount: 100.0,
				MerchantGroupID:   "premium_partners",
			},
			expected: 300.0,
		},
		{
			name: "Multiple Rules Combined",
			rules: []*domain.ProgramRule{
				{
					RuleName:       "Base points per amount",
					ConditionType:  "program_rule_transaction_amount",
					ConditionValue: "50",
					Multiplier:     0.1, // 0.1 points per dollar
					PointsAwarded:  0,
					EffectiveFrom:  yesterday,
					EffectiveTo:    timePtr(tomorrow),
				},
				{
					RuleName:       "Category bonus",
					ConditionType:  "program_rule_transaction_category",
					ConditionValue: "electronics",
					Multiplier:     2.0,
					PointsAwarded:  50,
					EffectiveFrom:  yesterday,
					EffectiveTo:    timePtr(tomorrow),
				},
			},
			tx: domain.Transaction{
				TransactionAmount: 100.0,
				Category:          "electronics",
			},
			expected: 110.0, // 10 points from amount (100 * 0.1) + 100 points from category bonus (50 * 2)
		},
		{
			name: "Promotional Points - Limited Time Offer",
			rules: []*domain.ProgramRule{
				{
					RuleName:       "Holiday Season Double Points",
					ConditionType:  "program_rule_transaction_amount",
					ConditionValue: "0",
					Multiplier:     2.0,
					PointsAwarded:  0,
					EffectiveFrom:  yesterday,
					EffectiveTo:    timePtr(tomorrow),
				},
			},
			tx: domain.Transaction{
				TransactionAmount: 100.0,
			},
			expected: 200.0,
		},
		{
			name: "Category Specific with Minimum Spend",
			rules: []*domain.ProgramRule{
				{
					RuleName:       "Electronics Category Base",
					ConditionType:  "program_rule_transaction_category",
					ConditionValue: "electronics",
					Multiplier:     1.0,
					PointsAwarded:  50,
					EffectiveFrom:  yesterday,
					EffectiveTo:    timePtr(tomorrow),
				},
				{
					RuleName:       "High Value Purchase Bonus",
					ConditionType:  "program_rule_transaction_amount",
					ConditionValue: "500",
					Multiplier:     2.0,
					PointsAwarded:  100,
					EffectiveFrom:  yesterday,
					EffectiveTo:    timePtr(tomorrow),
				},
			},
			tx: domain.Transaction{
				TransactionAmount: 600.0,
				Category:          "electronics",
			},
			expected: 250.0, // 50 (category) + 200 (high value)
		},
		{
			name: "Expired rule does not contribute",
			rules: []*domain.ProgramRule{
				{
					RuleName:       "Expired promo",
					ConditionType:  "program_rule_transaction_amount",
					ConditionValue: "0",
					Multiplier:     5.0,
					PointsAwarded:  0,
					EffectiveFrom:  now.AddDate(0, -1, 0),
					EffectiveTo:    timePtr(yesterday),
				},
			},
			tx: domain.Transaction{
				TransactionAmount: 100.0,
			},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calculatePoints(tt.rules, tt.tx)
			if got != tt.expected {
				t.Errorf("calculatePoints() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestEvaluateRule(t *testing.T) {
	merchantID := uuid.New()

	tests := []struct {
		name        string
		rule        *domain.ProgramRule
		tx          domain.Transaction
		wantMatches bool
		wantPoints  float64
	}{
		{
			name: "Transaction Amount Rule - Above Threshold",
			rule: &domain.ProgramRule{
				ConditionType:  "program_rule_transaction_amount",
				ConditionValue: "100",
				Multiplier:     0.1,
				PointsAwarded:  0,
			},
			tx: domain.Transaction{
				TransactionAmount: 150.0,
			},
			wantMatches: true,
			wantPoints:  15.0, // 150 * 0.1
		},
		{
			name: "Transaction Amount Rule - Below Threshold",
			rule: &domain.ProgramRule{
				ConditionType:  "program_rule_transaction_amount",
				ConditionValue: "100",
				Multiplier:     0.1,
				PointsAwarded:  0,
			},
			tx: domain.Transaction{
				TransactionAmount: 50.0,
			},
			wantMatches: false,
			wantPoints:  0,
		},
		{
			name: "Transaction Type Rule - Matching",
			rule: &domain.ProgramRule{
				ConditionType:  "program_rule_transaction_type",
				ConditionValue: "purchase",
				Multiplier:     2.0,
				PointsAwarded:  50,
			},
			tx: domain.Transaction{
				TransactionType: "purchase",
			},
			wantMatches: true,
			wantPoints:  100.0, // 50 * 2
		},
		{
			name: "Membership Tenure Rule - Exceeding",
			rule: &domain.ProgramRule{
				ConditionType:  "program_rule_tenure",
				ConditionValue: "365",
				Multiplier:     1.0,
				PointsAwarded:  1000,
			},
			tx: domain.Transaction{
				MembershipTenure: 400,
			},
			wantMatches: true,
			wantPoints:  1000.0,
		},
		{
			name: "Merchant Rule - Matching by ID",
			rule: &domain.ProgramRule{
				ConditionType:  "program_rule_transaction_merchant",
				ConditionValue: merchantID.String(),
				Multiplier:     1.0,
				PointsAwarded:  75,
			},
			tx: domain.Transaction{
				MerchantID: merchantID,
			},
			wantMatches: true,
			wantPoints:  75.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMatches, gotPoints := evaluateRule(tt.rule, tt.tx)
			if gotMatches != tt.wantMatches {
				t.Errorf("evaluateRule() matches = %v, want %v", gotMatches, tt.wantMatches)
			}
			if gotPoints != tt.wantPoints {
				t.Errorf("evaluateRule() points = %v, want %v", gotPoints, tt.wantPoints)
			}
		})
	}
}
