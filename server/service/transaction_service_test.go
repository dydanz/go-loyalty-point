package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"go-playground/server/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// mockTransactionRepo implements domain.TransactionRepository.
type mockTransactionRepo struct {
	mock.Mock
}

func (m *mockTransactionRepo) Create(ctx context.Context, transaction *domain.Transaction) (*domain.Transaction, error) {
	args := m.Called(ctx, transaction)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Transaction), args.Error(1)
}
func (m *mockTransactionRepo) GetByID(ctx context.Context, id uuid.UUID) (*domain.Transaction, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Transaction), args.Error(1)
}
func (m *mockTransactionRepo) GetByCustomerID(ctx context.Context, customerID uuid.UUID) ([]*domain.Transaction, error) {
	args := m.Called(ctx, customerID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Transaction), args.Error(1)
}
func (m *mockTransactionRepo) GetByCustomerIDWithPagination(ctx context.Context, customerID uuid.UUID, offset, limit int) ([]*domain.Transaction, int64, error) {
	args := m.Called(ctx, customerID, offset, limit)
	if args.Get(0) == nil {
		return nil, args.Get(1).(int64), args.Error(2)
	}
	return args.Get(0).([]*domain.Transaction), args.Get(1).(int64), args.Error(2)
}
func (m *mockTransactionRepo) GetByMerchantIDWithPagination(ctx context.Context, merchantID uuid.UUID, offset, limit int) ([]*domain.Transaction, int64, error) {
	args := m.Called(ctx, merchantID, offset, limit)
	if args.Get(0) == nil {
		return nil, args.Get(1).(int64), args.Error(2)
	}
	return args.Get(0).([]*domain.Transaction), args.Get(1).(int64), args.Error(2)
}
func (m *mockTransactionRepo) GetByUserIDWithPagination(ctx context.Context, userID uuid.UUID, offset, limit int) ([]*domain.Transaction, int64, error) {
	args := m.Called(ctx, userID, offset, limit)
	if args.Get(0) == nil {
		return nil, args.Get(1).(int64), args.Error(2)
	}
	return args.Get(0).([]*domain.Transaction), args.Get(1).(int64), args.Error(2)
}
func (m *mockTransactionRepo) UpdateStatus(ctx context.Context, id uuid.UUID, status string) error {
	args := m.Called(ctx, id, status)
	return args.Error(0)
}

// mockPointsServiceIface implements domain.PointsService.
type mockPointsServiceIface struct {
	mock.Mock
}

func (m *mockPointsServiceIface) GetLedger(ctx context.Context, customerID, programID uuid.UUID) ([]*domain.PointsLedger, error) {
	args := m.Called(ctx, customerID, programID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.PointsLedger), args.Error(1)
}
func (m *mockPointsServiceIface) GetBalance(ctx context.Context, customerID, programID uuid.UUID) (*domain.PointsBalance, error) {
	args := m.Called(ctx, customerID, programID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.PointsBalance), args.Error(1)
}
func (m *mockPointsServiceIface) EarnPoints(ctx context.Context, req *domain.PointsTransaction) (*domain.PointsTransaction, error) {
	args := m.Called(ctx, req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.PointsTransaction), args.Error(1)
}
func (m *mockPointsServiceIface) RedeemPoints(ctx context.Context, req *domain.PointsTransaction) (*domain.PointsTransaction, error) {
	args := m.Called(ctx, req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.PointsTransaction), args.Error(1)
}

// mockEventLoggerService implements domain.EventLoggerService.
type mockEventLoggerService struct {
	mock.Mock
}

func (m *mockEventLoggerService) SaveTransactionEvents(ctx context.Context, eventType domain.EventLogType, transaction *domain.Transaction, pointsEarned int) error {
	args := m.Called(ctx, eventType, transaction, pointsEarned)
	return args.Error(0)
}
func (m *mockEventLoggerService) SaveRedemptionEvents(ctx context.Context, eventType domain.EventLogType, redemption *domain.Redemption, reward *domain.Reward) error {
	args := m.Called(ctx, eventType, redemption, reward)
	return args.Error(0)
}
func (m *mockEventLoggerService) SaveUserUpdateEvents(ctx context.Context, eventType domain.EventLogType, user *domain.User) error {
	args := m.Called(ctx, eventType, user)
	return args.Error(0)
}
func (m *mockEventLoggerService) SaveMerchantUpdateEvents(ctx context.Context, eventType domain.EventLogType, merchant *domain.Merchant) error {
	args := m.Called(ctx, eventType, merchant)
	return args.Error(0)
}
func (m *mockEventLoggerService) SaveProgramUpdateEvents(ctx context.Context, eventType domain.EventLogType, program *domain.Program) error {
	args := m.Called(ctx, eventType, program)
	return args.Error(0)
}
func (m *mockEventLoggerService) SaveProgramRulesEvents(ctx context.Context, eventType domain.EventLogType, programRule *domain.ProgramRule) error {
	args := m.Called(ctx, eventType, programRule)
	return args.Error(0)
}
func (m *mockEventLoggerService) SavePointUpdateEvents(ctx context.Context, eventType domain.EventLogType, ledger *domain.PointsLedger) error {
	args := m.Called(ctx, eventType, ledger)
	return args.Error(0)
}

// mockMerchantCustomersRepo implements domain.MerchantCustomersRepository.
type mockMerchantCustomersRepo struct {
	mock.Mock
}

func (m *mockMerchantCustomersRepo) Create(ctx context.Context, customer *domain.MerchantCustomer) error {
	args := m.Called(ctx, customer)
	return args.Error(0)
}
func (m *mockMerchantCustomersRepo) GetByID(ctx context.Context, id uuid.UUID) (*domain.MerchantCustomer, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.MerchantCustomer), args.Error(1)
}
func (m *mockMerchantCustomersRepo) GetByEmail(ctx context.Context, email string) (*domain.MerchantCustomer, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.MerchantCustomer), args.Error(1)
}
func (m *mockMerchantCustomersRepo) GetByPhone(ctx context.Context, phone string) (*domain.MerchantCustomer, error) {
	args := m.Called(ctx, phone)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.MerchantCustomer), args.Error(1)
}
func (m *mockMerchantCustomersRepo) GetByMerchantID(ctx context.Context, merchantID uuid.UUID) ([]*domain.MerchantCustomer, error) {
	args := m.Called(ctx, merchantID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.MerchantCustomer), args.Error(1)
}
func (m *mockMerchantCustomersRepo) Update(ctx context.Context, customer *domain.MerchantCustomer) error {
	args := m.Called(ctx, customer)
	return args.Error(0)
}

// mockProgramRuleRepo implements domain.ProgramRuleRepository.
type mockProgramRuleRepo struct {
	mock.Mock
}

func (m *mockProgramRuleRepo) Create(ctx context.Context, rule *domain.ProgramRule) error {
	args := m.Called(ctx, rule)
	return args.Error(0)
}
func (m *mockProgramRuleRepo) GetByID(ctx context.Context, id uuid.UUID) (*domain.ProgramRule, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.ProgramRule), args.Error(1)
}
func (m *mockProgramRuleRepo) GetByProgramID(ctx context.Context, programID uuid.UUID) ([]*domain.ProgramRule, error) {
	args := m.Called(ctx, programID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.ProgramRule), args.Error(1)
}
func (m *mockProgramRuleRepo) Update(ctx context.Context, rule *domain.ProgramRule) error {
	args := m.Called(ctx, rule)
	return args.Error(0)
}
func (m *mockProgramRuleRepo) Delete(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
func (m *mockProgramRuleRepo) GetActiveRules(ctx context.Context, programID uuid.UUID, timestamp time.Time) ([]*domain.ProgramRule, error) {
	args := m.Called(ctx, programID, timestamp)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.ProgramRule), args.Error(1)
}

// TransactionServiceTestSuite defines the test suite.
type TransactionServiceTestSuite struct {
	suite.Suite
	transactionRepo *mockTransactionRepo
	pointsService   *mockPointsServiceIface
	eventLogger     *mockEventLoggerService
	merchantRepo    *mockMerchantCustomersRepo
	ruleRepo        *mockProgramRuleRepo
	service         *TransactionService
}

func (s *TransactionServiceTestSuite) SetupTest() {
	s.transactionRepo = new(mockTransactionRepo)
	s.pointsService = new(mockPointsServiceIface)
	s.eventLogger = new(mockEventLoggerService)
	s.merchantRepo = new(mockMerchantCustomersRepo)
	s.ruleRepo = new(mockProgramRuleRepo)
	s.service = NewTransactionService(s.transactionRepo, s.pointsService, s.eventLogger, s.merchantRepo, s.ruleRepo)
	s.eventLogger.On("SaveTransactionEvents", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil).Maybe()
}

func TestTransactionServiceTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionServiceTestSuite))
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

func (s *TransactionServiceTestSuite) baseRequest(merchantID, customerID, programID uuid.UUID, txType string, amount float64) *domain.CreateTransactionRequest {
	return &domain.CreateTransactionRequest{
		MerchantID:          merchantID,
		MerchantCustomersID: customerID,
		ProgramID:           programID,
		TransactionType:     txType,
		TransactionAmount:   amount,
		TransactionDate:     time.Now(),
		Status:              "completed",
	}
}

// TestCreate_EarnHappyPath_Purchase covers the happy path: a purchase transaction
// under an active rule computes rule-derived points and calls EarnPoints with
// exactly that amount.
func (s *TransactionServiceTestSuite) TestCreate_EarnHappyPath_Purchase() {
	ctx := context.Background()
	merchantID := uuid.New()
	customerID := uuid.New()
	programID := uuid.New()
	req := s.baseRequest(merchantID, customerID, programID, "purchase", 100.0)

	s.merchantRepo.On("GetByID", ctx, customerID).Return(&domain.MerchantCustomer{ID: customerID, MerchantID: merchantID}, nil)
	s.ruleRepo.On("GetActiveRules", ctx, programID, mock.Anything).Return([]*domain.ProgramRule{
		ruleFixture("program_rule_transaction_amount", "50", 1.0, 0), // 1 point per dollar over $50
	}, nil)
	s.transactionRepo.On("Create", ctx, mock.AnythingOfType("*domain.Transaction")).Return(&domain.Transaction{
		TransactionID:       uuid.New(),
		MerchantID:          merchantID,
		MerchantCustomersID: customerID,
		ProgramID:           programID,
		TransactionType:     "purchase",
		TransactionAmount:   100.0,
	}, nil)
	s.pointsService.On("EarnPoints", ctx, mock.MatchedBy(func(pt *domain.PointsTransaction) bool {
		return pt.Points == 100 && pt.CustomerID == customerID.String() && pt.ProgramID == programID.String()
	})).Return(&domain.PointsTransaction{Points: 100}, nil)

	result, err := s.service.Create(ctx, req)

	s.NoError(err)
	s.NotNil(result)
	s.pointsService.AssertCalled(s.T(), "EarnPoints", ctx, mock.Anything)
}

// TestCreate_RuleDrivenAmounts_DifferentRules is the FR-1.2 DoD regression: the
// same transaction under two different active rules yields two different point
// totals, proving points come from the rule engine and not a hardcoded formula.
func (s *TransactionServiceTestSuite) TestCreate_RuleDrivenAmounts_DifferentRules() {
	ctx := context.Background()
	merchantID := uuid.New()

	runWithRule := func(rule *domain.ProgramRule) int {
		s.SetupTest()
		customerID := uuid.New()
		programID := uuid.New()
		req := s.baseRequest(merchantID, customerID, programID, "purchase", 100.0)

		s.merchantRepo.On("GetByID", ctx, customerID).Return(&domain.MerchantCustomer{ID: customerID, MerchantID: merchantID}, nil)
		s.ruleRepo.On("GetActiveRules", ctx, programID, mock.Anything).Return([]*domain.ProgramRule{rule}, nil)
		s.transactionRepo.On("Create", ctx, mock.AnythingOfType("*domain.Transaction")).Return(&domain.Transaction{
			TransactionID: uuid.New(), MerchantID: merchantID, MerchantCustomersID: customerID, ProgramID: programID,
			TransactionType: "purchase", TransactionAmount: 100.0,
		}, nil)

		var gotPoints int
		s.pointsService.On("EarnPoints", ctx, mock.MatchedBy(func(pt *domain.PointsTransaction) bool {
			gotPoints = pt.Points
			return true
		})).Return(&domain.PointsTransaction{}, nil)

		_, err := s.service.Create(ctx, req)
		s.Require().NoError(err)
		return gotPoints
	}

	pointsA := runWithRule(ruleFixture("program_rule_transaction_amount", "50", 1.0, 0))
	pointsB := runWithRule(ruleFixture("program_rule_transaction_amount", "50", 5.0, 0))

	s.NotEqual(pointsA, pointsB)
	s.Equal(100, pointsA)
	s.Equal(500, pointsB)
}

func (s *TransactionServiceTestSuite) TestCreate_NoActiveRule_NoEarnCall() {
	ctx := context.Background()
	merchantID := uuid.New()
	customerID := uuid.New()
	programID := uuid.New()
	req := s.baseRequest(merchantID, customerID, programID, "purchase", 100.0)

	s.merchantRepo.On("GetByID", ctx, customerID).Return(&domain.MerchantCustomer{ID: customerID, MerchantID: merchantID}, nil)
	s.ruleRepo.On("GetActiveRules", ctx, programID, mock.Anything).Return([]*domain.ProgramRule{}, nil)
	s.transactionRepo.On("Create", ctx, mock.AnythingOfType("*domain.Transaction")).Return(&domain.Transaction{
		TransactionID: uuid.New(), MerchantID: merchantID, MerchantCustomersID: customerID, ProgramID: programID,
		TransactionType: "purchase", TransactionAmount: 100.0,
	}, nil)

	_, err := s.service.Create(ctx, req)

	s.NoError(err)
	s.pointsService.AssertNotCalled(s.T(), "EarnPoints", mock.Anything, mock.Anything)
	s.pointsService.AssertNotCalled(s.T(), "RedeemPoints", mock.Anything, mock.Anything)
}

func (s *TransactionServiceTestSuite) TestCreate_Refund_CallsRedeemPoints() {
	ctx := context.Background()
	merchantID := uuid.New()
	customerID := uuid.New()
	programID := uuid.New()
	req := s.baseRequest(merchantID, customerID, programID, "refund", 40.0)

	s.merchantRepo.On("GetByID", ctx, customerID).Return(&domain.MerchantCustomer{ID: customerID, MerchantID: merchantID}, nil)
	s.transactionRepo.On("Create", ctx, mock.AnythingOfType("*domain.Transaction")).Return(&domain.Transaction{
		TransactionID: uuid.New(), MerchantID: merchantID, MerchantCustomersID: customerID, ProgramID: programID,
		TransactionType: "refund", TransactionAmount: 40.0,
	}, nil)
	s.pointsService.On("RedeemPoints", ctx, mock.MatchedBy(func(pt *domain.PointsTransaction) bool {
		return pt.Points == -40
	})).Return(&domain.PointsTransaction{}, nil)

	_, err := s.service.Create(ctx, req)

	s.NoError(err)
	s.pointsService.AssertCalled(s.T(), "RedeemPoints", ctx, mock.Anything)
	s.ruleRepo.AssertNotCalled(s.T(), "GetActiveRules", mock.Anything, mock.Anything, mock.Anything)
}

func (s *TransactionServiceTestSuite) TestCreate_ValidationError_NonPositiveAmount() {
	ctx := context.Background()
	req := s.baseRequest(uuid.New(), uuid.New(), uuid.New(), "purchase", 0)

	result, err := s.service.Create(ctx, req)

	s.Nil(result)
	s.Error(err)
	var valErr domain.ValidationError
	s.True(errors.As(err, &valErr))
	s.merchantRepo.AssertNotCalled(s.T(), "GetByID", mock.Anything, mock.Anything)
}

func (s *TransactionServiceTestSuite) TestCreate_MerchantCustomerNotFound() {
	ctx := context.Background()
	customerID := uuid.New()
	req := s.baseRequest(uuid.New(), customerID, uuid.New(), "purchase", 100.0)

	s.merchantRepo.On("GetByID", ctx, customerID).Return(nil, nil)

	result, err := s.service.Create(ctx, req)

	s.Nil(result)
	s.Error(err)
	var sysErr domain.SystemError
	s.True(errors.As(err, &sysErr))
	var notFoundErr domain.ResourceNotFoundError
	s.True(errors.As(sysErr.Err, &notFoundErr))
}

func (s *TransactionServiceTestSuite) TestCreate_MerchantCustomerRepoError() {
	ctx := context.Background()
	customerID := uuid.New()
	req := s.baseRequest(uuid.New(), customerID, uuid.New(), "purchase", 100.0)

	s.merchantRepo.On("GetByID", ctx, customerID).Return(nil, errors.New("db down"))

	result, err := s.service.Create(ctx, req)

	s.Nil(result)
	s.Error(err)
	var sysErr domain.SystemError
	s.True(errors.As(err, &sysErr))
}

func (s *TransactionServiceTestSuite) TestCreate_TransactionRepoCreateError() {
	ctx := context.Background()
	merchantID := uuid.New()
	customerID := uuid.New()
	programID := uuid.New()
	req := s.baseRequest(merchantID, customerID, programID, "purchase", 100.0)

	s.merchantRepo.On("GetByID", ctx, customerID).Return(&domain.MerchantCustomer{ID: customerID, MerchantID: merchantID}, nil)
	s.transactionRepo.On("Create", ctx, mock.AnythingOfType("*domain.Transaction")).Return(nil, errors.New("insert failed"))

	result, err := s.service.Create(ctx, req)

	s.Nil(result)
	s.Error(err)
	var sysErr domain.SystemError
	s.True(errors.As(err, &sysErr))
}

func (s *TransactionServiceTestSuite) TestCreate_ProgramRuleRepoError() {
	ctx := context.Background()
	merchantID := uuid.New()
	customerID := uuid.New()
	programID := uuid.New()
	req := s.baseRequest(merchantID, customerID, programID, "purchase", 100.0)

	s.merchantRepo.On("GetByID", ctx, customerID).Return(&domain.MerchantCustomer{ID: customerID, MerchantID: merchantID}, nil)
	s.transactionRepo.On("Create", ctx, mock.AnythingOfType("*domain.Transaction")).Return(&domain.Transaction{
		TransactionID: uuid.New(), MerchantID: merchantID, MerchantCustomersID: customerID, ProgramID: programID,
		TransactionType: "purchase", TransactionAmount: 100.0,
	}, nil)
	s.ruleRepo.On("GetActiveRules", ctx, programID, mock.Anything).Return(nil, errors.New("rule query failed"))

	result, err := s.service.Create(ctx, req)

	s.Nil(result)
	s.Error(err)
	var sysErr domain.SystemError
	s.True(errors.As(err, &sysErr))
	s.pointsService.AssertNotCalled(s.T(), "EarnPoints", mock.Anything, mock.Anything)
}

func (s *TransactionServiceTestSuite) TestCreate_EarnPointsError() {
	ctx := context.Background()
	merchantID := uuid.New()
	customerID := uuid.New()
	programID := uuid.New()
	req := s.baseRequest(merchantID, customerID, programID, "purchase", 100.0)

	s.merchantRepo.On("GetByID", ctx, customerID).Return(&domain.MerchantCustomer{ID: customerID, MerchantID: merchantID}, nil)
	s.ruleRepo.On("GetActiveRules", ctx, programID, mock.Anything).Return([]*domain.ProgramRule{
		ruleFixture("program_rule_transaction_amount", "50", 1.0, 0),
	}, nil)
	s.transactionRepo.On("Create", ctx, mock.AnythingOfType("*domain.Transaction")).Return(&domain.Transaction{
		TransactionID: uuid.New(), MerchantID: merchantID, MerchantCustomersID: customerID, ProgramID: programID,
		TransactionType: "purchase", TransactionAmount: 100.0,
	}, nil)
	s.pointsService.On("EarnPoints", ctx, mock.Anything).Return(nil, errors.New("ledger write failed"))

	result, err := s.service.Create(ctx, req)

	s.Nil(result)
	s.Error(err)
	var sysErr domain.SystemError
	s.True(errors.As(err, &sysErr))
}

func (s *TransactionServiceTestSuite) TestCreate_RedeemPointsError() {
	ctx := context.Background()
	merchantID := uuid.New()
	customerID := uuid.New()
	programID := uuid.New()
	req := s.baseRequest(merchantID, customerID, programID, "refund", 40.0)

	s.merchantRepo.On("GetByID", ctx, customerID).Return(&domain.MerchantCustomer{ID: customerID, MerchantID: merchantID}, nil)
	s.transactionRepo.On("Create", ctx, mock.AnythingOfType("*domain.Transaction")).Return(&domain.Transaction{
		TransactionID: uuid.New(), MerchantID: merchantID, MerchantCustomersID: customerID, ProgramID: programID,
		TransactionType: "refund", TransactionAmount: 40.0,
	}, nil)
	s.pointsService.On("RedeemPoints", ctx, mock.Anything).Return(nil, errors.New("ledger write failed"))

	result, err := s.service.Create(ctx, req)

	s.Nil(result)
	s.Error(err)
	var sysErr domain.SystemError
	s.True(errors.As(err, &sysErr))
}

func (s *TransactionServiceTestSuite) TestGetByID_Success() {
	ctx := context.Background()
	id := uuid.New()
	expected := &domain.Transaction{TransactionID: id}
	s.transactionRepo.On("GetByID", ctx, id).Return(expected, nil)

	result, err := s.service.GetByID(ctx, id)

	s.NoError(err)
	s.Equal(expected, result)
}

func (s *TransactionServiceTestSuite) TestGetByID_NotFound() {
	ctx := context.Background()
	id := uuid.New()
	s.transactionRepo.On("GetByID", ctx, id).Return(nil, nil)

	result, err := s.service.GetByID(ctx, id)

	s.Nil(result)
	s.Error(err)
	var notFoundErr domain.ResourceNotFoundError
	s.True(errors.As(err, &notFoundErr))
}

func (s *TransactionServiceTestSuite) TestGetByID_RepoError() {
	ctx := context.Background()
	id := uuid.New()
	s.transactionRepo.On("GetByID", ctx, id).Return(nil, errors.New("db down"))

	result, err := s.service.GetByID(ctx, id)

	s.Nil(result)
	s.Error(err)
	var sysErr domain.SystemError
	s.True(errors.As(err, &sysErr))
}

func (s *TransactionServiceTestSuite) TestGetByCustomerID_Success() {
	ctx := context.Background()
	customerID := uuid.New()
	expected := []*domain.Transaction{{TransactionID: uuid.New()}}
	s.transactionRepo.On("GetByCustomerID", ctx, customerID).Return(expected, nil)

	result, err := s.service.GetByCustomerID(ctx, customerID)

	s.NoError(err)
	s.Equal(expected, result)
}

func (s *TransactionServiceTestSuite) TestGetByCustomerID_Empty() {
	ctx := context.Background()
	customerID := uuid.New()
	s.transactionRepo.On("GetByCustomerID", ctx, customerID).Return([]*domain.Transaction{}, nil)

	result, err := s.service.GetByCustomerID(ctx, customerID)

	s.NoError(err)
	s.Empty(result)
}

func (s *TransactionServiceTestSuite) TestGetByCustomerID_RepoError() {
	ctx := context.Background()
	customerID := uuid.New()
	s.transactionRepo.On("GetByCustomerID", ctx, customerID).Return(nil, errors.New("db down"))

	result, err := s.service.GetByCustomerID(ctx, customerID)

	s.Nil(result)
	s.Error(err)
	var sysErr domain.SystemError
	s.True(errors.As(err, &sysErr))
}

func (s *TransactionServiceTestSuite) TestGetByCustomerIDWithPagination() {
	ctx := context.Background()
	customerID := uuid.New()
	expected := []*domain.Transaction{{TransactionID: uuid.New()}}
	s.transactionRepo.On("GetByCustomerIDWithPagination", ctx, customerID, 0, 10).Return(expected, int64(1), nil)

	result, total, err := s.service.GetByCustomerIDWithPagination(ctx, customerID, 0, 10)

	s.NoError(err)
	s.Equal(expected, result)
	s.Equal(int64(1), total)
}

func (s *TransactionServiceTestSuite) TestGetByMerchantIDWithPagination() {
	ctx := context.Background()
	merchantID := uuid.New()
	expected := []*domain.Transaction{{TransactionID: uuid.New()}}
	s.transactionRepo.On("GetByMerchantIDWithPagination", ctx, merchantID, 0, 10).Return(expected, int64(1), nil)

	result, total, err := s.service.GetByMerchantIDWithPagination(ctx, merchantID, 0, 10)

	s.NoError(err)
	s.Equal(expected, result)
	s.Equal(int64(1), total)
}

func (s *TransactionServiceTestSuite) TestGetByUserIDWithPagination() {
	ctx := context.Background()
	userID := uuid.New()
	expected := []*domain.Transaction{{TransactionID: uuid.New()}}
	s.transactionRepo.On("GetByUserIDWithPagination", ctx, userID, 0, 10).Return(expected, int64(1), nil)

	result, total, err := s.service.GetByUserIDWithPagination(ctx, userID, 0, 10)

	s.NoError(err)
	s.Equal(expected, result)
	s.Equal(int64(1), total)
}

func (s *TransactionServiceTestSuite) TestUpdateStatus_Success() {
	ctx := context.Background()
	id := uuid.New()
	s.transactionRepo.On("UpdateStatus", ctx, id, "completed").Return(nil)

	err := s.service.UpdateStatus(ctx, id.String(), "completed")

	s.NoError(err)
}

func (s *TransactionServiceTestSuite) TestUpdateStatus_InvalidID() {
	ctx := context.Background()

	err := s.service.UpdateStatus(ctx, "not-a-uuid", "completed")

	s.Error(err)
	var sysErr domain.SystemError
	s.True(errors.As(err, &sysErr))
}

func (s *TransactionServiceTestSuite) TestUpdateStatus_RepoError() {
	ctx := context.Background()
	id := uuid.New()
	s.transactionRepo.On("UpdateStatus", ctx, id, "completed").Return(errors.New("db down"))

	err := s.service.UpdateStatus(ctx, id.String(), "completed")

	s.Error(err)
}

func (s *TransactionServiceTestSuite) TestSetPointsService() {
	newSvc := new(mockPointsServiceIface)
	s.service.SetPointsService(newSvc)
	s.Equal(newSvc, s.service.pointsService)
}
