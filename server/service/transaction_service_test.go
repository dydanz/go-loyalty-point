package service

import (
	"context"
	"testing"
	"time"

	"go-playground/server/domain"

	"github.com/google/uuid"
)

// fakeTransactionRepo is a minimal in-memory domain.TransactionRepository for wiring tests.
type fakeTransactionRepo struct{}

func (f *fakeTransactionRepo) Create(ctx context.Context, transaction *domain.Transaction) (*domain.Transaction, error) {
	transaction.TransactionID = uuid.New()
	return transaction, nil
}
func (f *fakeTransactionRepo) GetByID(ctx context.Context, id uuid.UUID) (*domain.Transaction, error) {
	return nil, nil
}
func (f *fakeTransactionRepo) GetByCustomerID(ctx context.Context, customerID uuid.UUID) ([]*domain.Transaction, error) {
	return nil, nil
}
func (f *fakeTransactionRepo) GetByCustomerIDWithPagination(ctx context.Context, customerID uuid.UUID, offset, limit int) ([]*domain.Transaction, int64, error) {
	return nil, 0, nil
}
func (f *fakeTransactionRepo) GetByMerchantIDWithPagination(ctx context.Context, merchantID uuid.UUID, offset, limit int) ([]*domain.Transaction, int64, error) {
	return nil, 0, nil
}
func (f *fakeTransactionRepo) GetByUserIDWithPagination(ctx context.Context, userID uuid.UUID, offset, limit int) ([]*domain.Transaction, int64, error) {
	return nil, 0, nil
}
func (f *fakeTransactionRepo) UpdateStatus(ctx context.Context, id uuid.UUID, status string) error {
	return nil
}

// fakeProgramRuleRepo returns a fixed set of rules for a given program ID.
type fakeProgramRuleRepo struct {
	rulesByProgram map[uuid.UUID][]*domain.ProgramRule
}

func (f *fakeProgramRuleRepo) Create(ctx context.Context, rule *domain.ProgramRule) error { return nil }
func (f *fakeProgramRuleRepo) GetByID(ctx context.Context, id uuid.UUID) (*domain.ProgramRule, error) {
	return nil, nil
}
func (f *fakeProgramRuleRepo) GetByProgramID(ctx context.Context, programID uuid.UUID) ([]*domain.ProgramRule, error) {
	return f.rulesByProgram[programID], nil
}
func (f *fakeProgramRuleRepo) Update(ctx context.Context, rule *domain.ProgramRule) error { return nil }
func (f *fakeProgramRuleRepo) Delete(ctx context.Context, id uuid.UUID) error             { return nil }
func (f *fakeProgramRuleRepo) GetActiveRules(ctx context.Context, programID uuid.UUID, timestamp time.Time) ([]*domain.ProgramRule, error) {
	return f.rulesByProgram[programID], nil
}

// fakePointsService records the points it was asked to earn/redeem.
type fakePointsService struct {
	lastEarnPoints int
	earnCalls      int
}

func (f *fakePointsService) GetLedger(ctx context.Context, customerID, programID uuid.UUID) ([]*domain.PointsLedger, error) {
	return nil, nil
}
func (f *fakePointsService) GetBalance(ctx context.Context, customerID, programID uuid.UUID) (*domain.PointsBalance, error) {
	return nil, nil
}
func (f *fakePointsService) EarnPoints(ctx context.Context, req *domain.PointsTransaction) (*domain.PointsTransaction, error) {
	f.lastEarnPoints = req.Points
	f.earnCalls++
	return req, nil
}
func (f *fakePointsService) RedeemPoints(ctx context.Context, req *domain.PointsTransaction) (*domain.PointsTransaction, error) {
	return req, nil
}

// fakeEventLoggerService is a no-op logger for wiring tests.
type fakeEventLoggerService struct{}

func (f *fakeEventLoggerService) SaveTransactionEvents(ctx context.Context, eventType domain.EventLogType, transaction *domain.Transaction, pointsEarned int) error {
	return nil
}
func (f *fakeEventLoggerService) SaveRedemptionEvents(ctx context.Context, eventType domain.EventLogType, redemption *domain.Redemption, reward *domain.Reward) error {
	return nil
}
func (f *fakeEventLoggerService) SaveUserUpdateEvents(ctx context.Context, eventType domain.EventLogType, user *domain.User) error {
	return nil
}
func (f *fakeEventLoggerService) SaveMerchantUpdateEvents(ctx context.Context, eventType domain.EventLogType, merchant *domain.Merchant) error {
	return nil
}
func (f *fakeEventLoggerService) SavePointUpdateEvents(ctx context.Context, eventType domain.EventLogType, ledger *domain.PointsLedger) error {
	return nil
}
func (f *fakeEventLoggerService) SaveProgramUpdateEvents(ctx context.Context, eventType domain.EventLogType, program *domain.Program) error {
	return nil
}
func (f *fakeEventLoggerService) SaveProgramRulesEvents(ctx context.Context, eventType domain.EventLogType, programRule *domain.ProgramRule) error {
	return nil
}

// fakeMerchantCustomersRepo resolves any customer to a fixed merchant.
type fakeMerchantCustomersRepo struct {
	merchantID uuid.UUID
}

func (f *fakeMerchantCustomersRepo) Create(ctx context.Context, customer *domain.MerchantCustomer) error {
	return nil
}
func (f *fakeMerchantCustomersRepo) GetByID(ctx context.Context, id uuid.UUID) (*domain.MerchantCustomer, error) {
	return &domain.MerchantCustomer{ID: id, MerchantID: f.merchantID}, nil
}
func (f *fakeMerchantCustomersRepo) GetByEmail(ctx context.Context, email string) (*domain.MerchantCustomer, error) {
	return nil, nil
}
func (f *fakeMerchantCustomersRepo) GetByPhone(ctx context.Context, phone string) (*domain.MerchantCustomer, error) {
	return nil, nil
}
func (f *fakeMerchantCustomersRepo) GetByMerchantID(ctx context.Context, merchantID uuid.UUID) ([]*domain.MerchantCustomer, error) {
	return nil, nil
}
func (f *fakeMerchantCustomersRepo) Update(ctx context.Context, customer *domain.MerchantCustomer) error {
	return nil
}

func newTestTransactionService(rules []*domain.ProgramRule, points *fakePointsService, merchantID uuid.UUID) (*TransactionService, uuid.UUID) {
	programID := uuid.New()
	ruleRepo := &fakeProgramRuleRepo{rulesByProgram: map[uuid.UUID][]*domain.ProgramRule{programID: rules}}
	svc := NewTransactionService(
		&fakeTransactionRepo{},
		points,
		&fakeEventLoggerService{},
		&fakeMerchantCustomersRepo{merchantID: merchantID},
		ruleRepo,
	)
	return svc, programID
}

func ruleFixture(conditionType, conditionValue string, multiplier float64, pointsAwarded int) *domain.ProgramRule {
	now := time.Now()
	return &domain.ProgramRule{
		ID:             uuid.New(),
		RuleName:       "fixture",
		ConditionType:  conditionType,
		ConditionValue: conditionValue,
		Multiplier:     multiplier,
		PointsAwarded:  pointsAwarded,
		EffectiveFrom:  now.AddDate(0, 0, -1),
	}
}

// TestCreate_SameTransactionDifferentRulesYieldsDifferentPoints is the FR-1.2 DoD
// regression: identical purchase transactions under different active rules must
// produce different point totals, proving points come from the rule engine and
// not a hardcoded formula.
func TestCreate_SameTransactionDifferentRulesYieldsDifferentPoints(t *testing.T) {
	merchantID := uuid.New()

	pointsA := &fakePointsService{}
	ruleA := ruleFixture("program_rule_transaction_amount", "50", 1.0, 0) // 1 point per dollar over $50
	svcA, programIDA := newTestTransactionService([]*domain.ProgramRule{ruleA}, pointsA, merchantID)

	pointsB := &fakePointsService{}
	ruleB := ruleFixture("program_rule_transaction_amount", "50", 5.0, 0) // 5 points per dollar over $50
	svcB, programIDB := newTestTransactionService([]*domain.ProgramRule{ruleB}, pointsB, merchantID)

	ctx := context.Background()
	customerID := uuid.New()

	req := &domain.CreateTransactionRequest{
		MerchantID:          merchantID,
		MerchantCustomersID: customerID,
		ProgramID:           programIDA,
		TransactionType:     "purchase",
		TransactionAmount:   100.0,
		TransactionDate:     time.Now(),
		Status:              "completed",
	}
	if _, err := svcA.Create(ctx, req); err != nil {
		t.Fatalf("svcA.Create() error = %v", err)
	}

	req.ProgramID = programIDB
	if _, err := svcB.Create(ctx, req); err != nil {
		t.Fatalf("svcB.Create() error = %v", err)
	}

	if pointsA.lastEarnPoints == pointsB.lastEarnPoints {
		t.Fatalf("expected different point totals under different rules, got %d for both", pointsA.lastEarnPoints)
	}
	if pointsA.lastEarnPoints != 100 {
		t.Errorf("ruleA: got %d points, want 100", pointsA.lastEarnPoints)
	}
	if pointsB.lastEarnPoints != 500 {
		t.Errorf("ruleB: got %d points, want 500", pointsB.lastEarnPoints)
	}
}

func TestCreate_NoActiveRuleYieldsZeroPoints(t *testing.T) {
	merchantID := uuid.New()
	points := &fakePointsService{}
	svc, programID := newTestTransactionService(nil, points, merchantID)

	req := &domain.CreateTransactionRequest{
		MerchantID:          merchantID,
		MerchantCustomersID: uuid.New(),
		ProgramID:           programID,
		TransactionType:     "purchase",
		TransactionAmount:   100.0,
		TransactionDate:     time.Now(),
		Status:              "completed",
	}
	if _, err := svc.Create(context.Background(), req); err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	if points.earnCalls != 0 {
		t.Errorf("expected no earn call when computed points is 0, got %d calls with %d points", points.earnCalls, points.lastEarnPoints)
	}
}

func TestCreate_OverlappingRulesSumPoints(t *testing.T) {
	merchantID := uuid.New()
	points := &fakePointsService{}
	amountRule := ruleFixture("program_rule_transaction_amount", "50", 1.0, 0)  // 100 points (100 * 1.0)
	typeRule := ruleFixture("program_rule_transaction_type", "purchase", 1.0, 25) // +25 flat bonus
	svc, programID := newTestTransactionService([]*domain.ProgramRule{amountRule, typeRule}, points, merchantID)

	req := &domain.CreateTransactionRequest{
		MerchantID:          merchantID,
		MerchantCustomersID: uuid.New(),
		ProgramID:           programID,
		TransactionType:     "purchase",
		TransactionAmount:   100.0,
		TransactionDate:     time.Now(),
		Status:              "completed",
	}
	if _, err := svc.Create(context.Background(), req); err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	if points.lastEarnPoints != 125 {
		t.Errorf("got %d points, want 125 (100 amount + 25 type bonus)", points.lastEarnPoints)
	}
}

func TestCreate_RefundKeepsHardcodedDeduction(t *testing.T) {
	merchantID := uuid.New()
	points := &fakePointsService{}
	svc, programID := newTestTransactionService(nil, points, merchantID)

	req := &domain.CreateTransactionRequest{
		MerchantID:          merchantID,
		MerchantCustomersID: uuid.New(),
		ProgramID:           programID,
		TransactionType:     "refund",
		TransactionAmount:   40.0,
		TransactionDate:     time.Now(),
		Status:              "completed",
	}
	if _, err := svc.Create(context.Background(), req); err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	if points.earnCalls != 0 {
		t.Errorf("refund should not call EarnPoints, got %d calls", points.earnCalls)
	}
}
