package service

import (
	"fmt"
	"time"

	"go-playground/server/domain"
)

func evaluateRule(rule *domain.ProgramRule, tx domain.Transaction) (bool, float64) {
	switch rule.ConditionType {
	case "program_rule_tenure":
		// Check if membership tenure meets the condition
		tenure := tx.MembershipTenure
		return float64(tenure) > parseConditionValue(rule.ConditionValue), rule.Multiplier * float64(rule.PointsAwarded)

	case "program_rule_transaction_amount":
		// Check if transaction amount meets the condition
		if tx.TransactionAmount > parseConditionValue(rule.ConditionValue) {
			if rule.PointsAwarded == 0 {
				return true, rule.Multiplier * tx.TransactionAmount
			}
			return true, rule.Multiplier * float64(rule.PointsAwarded)
		}
	case "program_rule_transaction_count":
		// Check if transaction count meets the condition
		count := tx.TransactionCount
		return float64(count) > parseConditionValue(rule.ConditionValue), rule.Multiplier * float64(rule.PointsAwarded)

	case "program_rule_transaction_type":
		// Check if transaction type matches the condition
		return tx.TransactionType == rule.ConditionValue, rule.Multiplier * float64(rule.PointsAwarded)

	case "program_rule_transaction_category":
		// Check if transaction category matches the condition
		return tx.Category == rule.ConditionValue, rule.Multiplier * float64(rule.PointsAwarded)

	case "program_rule_transaction_merchant":
		// Check if transaction merchant matches the condition
		return tx.MerchantID.String() == rule.ConditionValue, rule.Multiplier * float64(rule.PointsAwarded)

	case "program_rule_transaction_merchant_group":
		// Check if transaction merchant group matches the condition
		return tx.MerchantGroupID == rule.ConditionValue, rule.Multiplier * float64(rule.PointsAwarded)
	}
	return false, 0
}

// Helper function to parse condition value (e.g., "> 100")
func parseConditionValue(condition string) float64 {
	// Implement logic to parse condition strings like "> 100"
	// For simplicity, assume the condition is a numeric value
	var value float64
	fmt.Sscanf(condition, "%f", &value)
	return value
}

// calculatePoints evaluates every active rule against tx and sums the points
// awarded. Transaction-amount rules are summed as base points, all other
// condition types as bonus points, so an amount rule and a bonus rule both
// matching contribute independently rather than the second overwriting the
// first.
func calculatePoints(rules []*domain.ProgramRule, tx domain.Transaction) float64 {
	totalPoints := 0.0
	basePoints := 0.0
	bonusPoints := 0.0
	now := time.Now()

	// First pass: Calculate base points from transaction amount rules
	for _, rule := range rules {
		if rule.ConditionType == "program_rule_transaction_amount" {
			// Skip expired or future rules
			if now.Before(rule.EffectiveFrom) || (rule.EffectiveTo != nil && now.After(*rule.EffectiveTo)) {
				continue
			}

			matches, points := evaluateRule(rule, tx)
			if matches {
				basePoints += points
			}
		}
	}

	// Second pass: Calculate bonus points from other rules
	for _, rule := range rules {
		// Skip expired or future rules
		if now.Before(rule.EffectiveFrom) || (rule.EffectiveTo != nil && now.After(*rule.EffectiveTo)) {
			continue
		}

		// Skip transaction amount rules as they were handled in first pass
		if rule.ConditionType != "program_rule_transaction_amount" {
			matches, points := evaluateRule(rule, tx)
			if matches {
				bonusPoints += points
			}
		}
	}

	totalPoints = basePoints + bonusPoints
	return totalPoints
}
