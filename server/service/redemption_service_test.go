package service

import (
	"context"
	"errors"
	"testing"

	"go-playground/server/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// mockRedemptionRepo implements domain.RedemptionRepository.
type mockRedemptionRepo struct {
	mock.Mock
}

func (m *mockRedemptionRepo) Create(ctx context.Context, redemption *domain.Redemption) ([]*domain.Redemption, error) {
	args := m.Called(ctx, redemption)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Redemption), args.Error(1)
}
func (m *mockRedemptionRepo) GetByID(ctx context.Context, id uuid.UUID) (*domain.Redemption, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Redemption), args.Error(1)
}
func (m *mockRedemptionRepo) GetByUserID(ctx context.Context, userID uuid.UUID) ([]*domain.Redemption, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Redemption), args.Error(1)
}
func (m *mockRedemptionRepo) Update(ctx context.Context, redemption *domain.Redemption) error {
	args := m.Called(ctx, redemption)
	return args.Error(0)
}

// mockRewardsRepo implements domain.RewardsRepository.
type mockRewardsRepo struct {
	mock.Mock
}

func (m *mockRewardsRepo) Create(ctx context.Context, reward *domain.Reward) (*domain.Reward, error) {
	args := m.Called(ctx, reward)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Reward), args.Error(1)
}
func (m *mockRewardsRepo) GetByID(ctx context.Context, id uuid.UUID) (*domain.Reward, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Reward), args.Error(1)
}
func (m *mockRewardsRepo) Update(ctx context.Context, reward *domain.Reward) (*domain.Reward, error) {
	args := m.Called(ctx, reward)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Reward), args.Error(1)
}
func (m *mockRewardsRepo) Delete(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
func (m *mockRewardsRepo) GetAll(ctx context.Context, activeOnly bool) ([]domain.Reward, error) {
	args := m.Called(ctx, activeOnly)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]domain.Reward), args.Error(1)
}
func (m *mockRewardsRepo) GetByProgramID(ctx context.Context, programID uuid.UUID) ([]*domain.Reward, error) {
	args := m.Called(ctx, programID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Reward), args.Error(1)
}

// mockTransactionServiceIface implements domain.TransactionService.
type mockTransactionServiceIface struct {
	mock.Mock
}

func (m *mockTransactionServiceIface) Create(ctx context.Context, req *domain.CreateTransactionRequest) (*domain.Transaction, error) {
	args := m.Called(ctx, req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Transaction), args.Error(1)
}
func (m *mockTransactionServiceIface) GetByID(ctx context.Context, id uuid.UUID) (*domain.Transaction, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Transaction), args.Error(1)
}
func (m *mockTransactionServiceIface) GetByCustomerID(ctx context.Context, customerID uuid.UUID) ([]*domain.Transaction, error) {
	args := m.Called(ctx, customerID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Transaction), args.Error(1)
}
func (m *mockTransactionServiceIface) GetByCustomerIDWithPagination(ctx context.Context, customerID uuid.UUID, offset, limit int) ([]*domain.Transaction, int64, error) {
	args := m.Called(ctx, customerID, offset, limit)
	if args.Get(0) == nil {
		return nil, args.Get(1).(int64), args.Error(2)
	}
	return args.Get(0).([]*domain.Transaction), args.Get(1).(int64), args.Error(2)
}
func (m *mockTransactionServiceIface) UpdateStatus(ctx context.Context, id string, status string) error {
	args := m.Called(ctx, id, status)
	return args.Error(0)
}
func (m *mockTransactionServiceIface) GetByMerchantIDWithPagination(ctx context.Context, merchantID uuid.UUID, offset, limit int) ([]*domain.Transaction, int64, error) {
	args := m.Called(ctx, merchantID, offset, limit)
	if args.Get(0) == nil {
		return nil, args.Get(1).(int64), args.Error(2)
	}
	return args.Get(0).([]*domain.Transaction), args.Get(1).(int64), args.Error(2)
}
func (m *mockTransactionServiceIface) GetByUserIDWithPagination(ctx context.Context, userID uuid.UUID, offset, limit int) ([]*domain.Transaction, int64, error) {
	args := m.Called(ctx, userID, offset, limit)
	if args.Get(0) == nil {
		return nil, args.Get(1).(int64), args.Error(2)
	}
	return args.Get(0).([]*domain.Transaction), args.Get(1).(int64), args.Error(2)
}

// RedemptionServiceTestSuite defines the test suite.
type RedemptionServiceTestSuite struct {
	suite.Suite
	redemptionRepo *mockRedemptionRepo
	rewardsRepo    *mockRewardsRepo
	pointsService  *mockPointsServiceIface
	txService      *mockTransactionServiceIface
	eventLogger    *mockEventLoggerService
	service        *RedemptionService
}

func (s *RedemptionServiceTestSuite) SetupTest() {
	s.redemptionRepo = new(mockRedemptionRepo)
	s.rewardsRepo = new(mockRewardsRepo)
	s.pointsService = new(mockPointsServiceIface)
	s.txService = new(mockTransactionServiceIface)
	s.eventLogger = new(mockEventLoggerService)
	s.service = NewRedemptionService(s.redemptionRepo, s.rewardsRepo, s.pointsService, s.txService, s.eventLogger)
	s.eventLogger.On("SaveRedemptionEvents", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil).Maybe()
}

func TestRedemptionServiceTestSuite(t *testing.T) {
	suite.Run(t, new(RedemptionServiceTestSuite))
}

func activeReward(programID uuid.UUID, pointsRequired int) *domain.Reward {
	return &domain.Reward{
		ID:             uuid.New(),
		ProgramID:      programID,
		Name:           "fixture reward",
		PointsRequired: pointsRequired,
		IsActive:       true,
	}
}

// TestCreate_Success covers the happy path: sufficient balance deducts exactly
// reward.PointsRequired via a paired redemption transaction (FR-1.3).
func (s *RedemptionServiceTestSuite) TestCreate_Success() {
	ctx := context.Background()
	customerID := uuid.New()
	programID := uuid.New()
	reward := activeReward(programID, 100)

	redemption := &domain.Redemption{
		ID:                  uuid.New(),
		MerchantCustomersID: customerID,
		RewardID:            reward.ID,
	}

	s.rewardsRepo.On("GetByID", ctx, reward.ID).Return(reward, nil)
	s.pointsService.On("GetBalance", ctx, customerID, programID).Return(&domain.PointsBalance{Balance: 150}, nil)
	s.redemptionRepo.On("Create", ctx, mock.MatchedBy(func(r *domain.Redemption) bool {
		return r.PointsUsed == 100
	})).Return([]*domain.Redemption{redemption}, nil)
	s.txService.On("Create", ctx, mock.MatchedBy(func(req *domain.CreateTransactionRequest) bool {
		return req.TransactionType == "redemption" && req.TransactionAmount == 100
	})).Return(&domain.Transaction{TransactionID: uuid.New()}, nil)

	err := s.service.Create(ctx, redemption)

	s.NoError(err)
	s.txService.AssertCalled(s.T(), "Create", ctx, mock.Anything)
}

// TestCreate_InsufficientBalance_DeductsNothing is the FR-1.3 DoD: a failed
// redemption (insufficient balance) must not create a redemption record or a
// deducting transaction.
func (s *RedemptionServiceTestSuite) TestCreate_InsufficientBalance_DeductsNothing() {
	ctx := context.Background()
	customerID := uuid.New()
	programID := uuid.New()
	reward := activeReward(programID, 100)

	redemption := &domain.Redemption{MerchantCustomersID: customerID, RewardID: reward.ID}

	s.rewardsRepo.On("GetByID", ctx, reward.ID).Return(reward, nil)
	s.pointsService.On("GetBalance", ctx, customerID, programID).Return(&domain.PointsBalance{Balance: 50}, nil)

	err := s.service.Create(ctx, redemption)

	s.Error(err)
	var bizErr domain.BusinessLogicError
	s.True(errors.As(err, &bizErr))
	s.Equal("INSUFFICIENT_POINTS", bizErr.Code)
	s.redemptionRepo.AssertNotCalled(s.T(), "Create", mock.Anything, mock.Anything)
	s.txService.AssertNotCalled(s.T(), "Create", mock.Anything, mock.Anything)
}

func (s *RedemptionServiceTestSuite) TestCreate_RewardNotFound() {
	ctx := context.Background()
	rewardID := uuid.New()
	redemption := &domain.Redemption{MerchantCustomersID: uuid.New(), RewardID: rewardID}

	s.rewardsRepo.On("GetByID", ctx, rewardID).Return(nil, nil)

	err := s.service.Create(ctx, redemption)

	s.Error(err)
	var notFoundErr domain.ResourceNotFoundError
	s.True(errors.As(err, &notFoundErr))
}

func (s *RedemptionServiceTestSuite) TestCreate_RewardRepoError() {
	ctx := context.Background()
	rewardID := uuid.New()
	redemption := &domain.Redemption{MerchantCustomersID: uuid.New(), RewardID: rewardID}

	s.rewardsRepo.On("GetByID", ctx, rewardID).Return(nil, errors.New("db down"))

	err := s.service.Create(ctx, redemption)

	s.Error(err)
	var sysErr domain.SystemError
	s.True(errors.As(err, &sysErr))
}

func (s *RedemptionServiceTestSuite) TestCreate_RewardInactive() {
	ctx := context.Background()
	reward := activeReward(uuid.New(), 100)
	reward.IsActive = false
	redemption := &domain.Redemption{MerchantCustomersID: uuid.New(), RewardID: reward.ID}

	s.rewardsRepo.On("GetByID", ctx, reward.ID).Return(reward, nil)

	err := s.service.Create(ctx, redemption)

	s.Error(err)
	var bizErr domain.BusinessLogicError
	s.True(errors.As(err, &bizErr))
	s.Equal("REWARD_INACTIVE", bizErr.Code)
}

func (s *RedemptionServiceTestSuite) TestCreate_GetBalanceError() {
	ctx := context.Background()
	customerID := uuid.New()
	programID := uuid.New()
	reward := activeReward(programID, 100)
	redemption := &domain.Redemption{MerchantCustomersID: customerID, RewardID: reward.ID}

	s.rewardsRepo.On("GetByID", ctx, reward.ID).Return(reward, nil)
	s.pointsService.On("GetBalance", ctx, customerID, programID).Return(nil, errors.New("cache down"))

	err := s.service.Create(ctx, redemption)

	s.Error(err)
	var sysErr domain.SystemError
	s.True(errors.As(err, &sysErr))
}

func (s *RedemptionServiceTestSuite) TestCreate_RedemptionRepoCreateError() {
	ctx := context.Background()
	customerID := uuid.New()
	programID := uuid.New()
	reward := activeReward(programID, 100)
	redemption := &domain.Redemption{MerchantCustomersID: customerID, RewardID: reward.ID}

	s.rewardsRepo.On("GetByID", ctx, reward.ID).Return(reward, nil)
	s.pointsService.On("GetBalance", ctx, customerID, programID).Return(&domain.PointsBalance{Balance: 150}, nil)
	s.redemptionRepo.On("Create", ctx, mock.Anything).Return(nil, errors.New("insert failed"))

	err := s.service.Create(ctx, redemption)

	s.Error(err)
	var sysErr domain.SystemError
	s.True(errors.As(err, &sysErr))
	s.txService.AssertNotCalled(s.T(), "Create", mock.Anything, mock.Anything)
}

func (s *RedemptionServiceTestSuite) TestCreate_TransactionServiceError() {
	ctx := context.Background()
	customerID := uuid.New()
	programID := uuid.New()
	reward := activeReward(programID, 100)
	redemption := &domain.Redemption{ID: uuid.New(), MerchantCustomersID: customerID, RewardID: reward.ID}

	s.rewardsRepo.On("GetByID", ctx, reward.ID).Return(reward, nil)
	s.pointsService.On("GetBalance", ctx, customerID, programID).Return(&domain.PointsBalance{Balance: 150}, nil)
	s.redemptionRepo.On("Create", ctx, mock.Anything).Return([]*domain.Redemption{redemption}, nil)
	s.txService.On("Create", ctx, mock.Anything).Return(nil, errors.New("ledger write failed"))

	err := s.service.Create(ctx, redemption)

	s.Error(err)
	var sysErr domain.SystemError
	s.True(errors.As(err, &sysErr))
}

func (s *RedemptionServiceTestSuite) TestGetByID_Success() {
	id := uuid.New()
	expected := &domain.Redemption{ID: id}
	s.redemptionRepo.On("GetByID", mock.Anything, id).Return(expected, nil)

	result, err := s.service.GetByID(id.String())

	s.NoError(err)
	s.Equal(expected, result)
}

func (s *RedemptionServiceTestSuite) TestGetByID_NotFound() {
	id := uuid.New()
	s.redemptionRepo.On("GetByID", mock.Anything, id).Return(nil, nil)

	result, err := s.service.GetByID(id.String())

	s.Nil(result)
	s.Error(err)
	var notFoundErr domain.ResourceNotFoundError
	s.True(errors.As(err, &notFoundErr))
}

func (s *RedemptionServiceTestSuite) TestGetByID_RepoError() {
	id := uuid.New()
	s.redemptionRepo.On("GetByID", mock.Anything, id).Return(nil, errors.New("db down"))

	result, err := s.service.GetByID(id.String())

	s.Nil(result)
	s.Error(err)
	var sysErr domain.SystemError
	s.True(errors.As(err, &sysErr))
}

func (s *RedemptionServiceTestSuite) TestGetByUserID_Success() {
	userID := uuid.New()
	expected := []*domain.Redemption{{ID: uuid.New()}}
	s.redemptionRepo.On("GetByUserID", mock.Anything, userID).Return(expected, nil)

	result, err := s.service.GetByUserID(userID.String())

	s.NoError(err)
	s.Equal(expected, result)
}

func (s *RedemptionServiceTestSuite) TestGetByUserID_Empty() {
	userID := uuid.New()
	s.redemptionRepo.On("GetByUserID", mock.Anything, userID).Return([]*domain.Redemption{}, nil)

	result, err := s.service.GetByUserID(userID.String())

	s.NoError(err)
	s.Empty(result)
}

func (s *RedemptionServiceTestSuite) TestGetByUserID_RepoError() {
	userID := uuid.New()
	s.redemptionRepo.On("GetByUserID", mock.Anything, userID).Return(nil, errors.New("db down"))

	result, err := s.service.GetByUserID(userID.String())

	s.Nil(result)
	s.Error(err)
	var sysErr domain.SystemError
	s.True(errors.As(err, &sysErr))
}

func (s *RedemptionServiceTestSuite) TestUpdateStatus_NotFound() {
	ctx := context.Background()
	id := uuid.New()
	s.redemptionRepo.On("GetByID", mock.Anything, id).Return(nil, nil)

	err := s.service.UpdateStatus(ctx, id.String(), "completed")

	s.Error(err)
	var notFoundErr domain.ResourceNotFoundError
	s.True(errors.As(err, &notFoundErr))
}

func (s *RedemptionServiceTestSuite) TestUpdateStatus_GetRepoError() {
	ctx := context.Background()
	id := uuid.New()
	s.redemptionRepo.On("GetByID", mock.Anything, id).Return(nil, errors.New("db down"))

	err := s.service.UpdateStatus(ctx, id.String(), "completed")

	s.Error(err)
	var sysErr domain.SystemError
	s.True(errors.As(err, &sysErr))
}

// TestUpdateStatus_CancelPendingRefundsPoints covers the pending -> canceled
// path, which must refund exactly reward.PointsRequired via EarnPoints.
func (s *RedemptionServiceTestSuite) TestUpdateStatus_CancelPendingRefundsPoints() {
	ctx := context.Background()
	customerID := uuid.New()
	programID := uuid.New()
	reward := activeReward(programID, 100)
	id := uuid.New()
	redemption := &domain.Redemption{
		ID:                  id,
		MerchantCustomersID: customerID,
		RewardID:            reward.ID,
		Status:              domain.RedemptionStatusPending,
	}

	s.redemptionRepo.On("GetByID", mock.Anything, id).Return(redemption, nil)
	s.rewardsRepo.On("GetByID", ctx, reward.ID).Return(reward, nil)
	s.pointsService.On("EarnPoints", ctx, mock.MatchedBy(func(pt *domain.PointsTransaction) bool {
		return pt.Points == 100 && pt.Type == "refund"
	})).Return(&domain.PointsTransaction{}, nil)
	s.redemptionRepo.On("Update", ctx, mock.MatchedBy(func(r *domain.Redemption) bool {
		return r.Status == "canceled"
	})).Return(nil)

	err := s.service.UpdateStatus(ctx, id.String(), "canceled")

	s.NoError(err)
	s.pointsService.AssertCalled(s.T(), "EarnPoints", ctx, mock.Anything)
}

func (s *RedemptionServiceTestSuite) TestUpdateStatus_CancelPending_RewardNotFound() {
	ctx := context.Background()
	rewardID := uuid.New()
	id := uuid.New()
	redemption := &domain.Redemption{ID: id, RewardID: rewardID, Status: domain.RedemptionStatusPending}

	s.redemptionRepo.On("GetByID", mock.Anything, id).Return(redemption, nil)
	s.rewardsRepo.On("GetByID", ctx, rewardID).Return(nil, nil)

	err := s.service.UpdateStatus(ctx, id.String(), "canceled")

	s.Error(err)
	var notFoundErr domain.ResourceNotFoundError
	s.True(errors.As(err, &notFoundErr))
}

func (s *RedemptionServiceTestSuite) TestUpdateStatus_CancelPending_RewardRepoError() {
	ctx := context.Background()
	rewardID := uuid.New()
	id := uuid.New()
	redemption := &domain.Redemption{ID: id, RewardID: rewardID, Status: domain.RedemptionStatusPending}

	s.redemptionRepo.On("GetByID", mock.Anything, id).Return(redemption, nil)
	s.rewardsRepo.On("GetByID", ctx, rewardID).Return(nil, errors.New("db down"))

	err := s.service.UpdateStatus(ctx, id.String(), "canceled")

	s.Error(err)
	var sysErr domain.SystemError
	s.True(errors.As(err, &sysErr))
}

func (s *RedemptionServiceTestSuite) TestUpdateStatus_CancelPending_EarnPointsError() {
	ctx := context.Background()
	programID := uuid.New()
	reward := activeReward(programID, 100)
	id := uuid.New()
	redemption := &domain.Redemption{ID: id, MerchantCustomersID: uuid.New(), RewardID: reward.ID, Status: domain.RedemptionStatusPending}

	s.redemptionRepo.On("GetByID", mock.Anything, id).Return(redemption, nil)
	s.rewardsRepo.On("GetByID", ctx, reward.ID).Return(reward, nil)
	s.pointsService.On("EarnPoints", ctx, mock.Anything).Return(nil, errors.New("ledger write failed"))

	err := s.service.UpdateStatus(ctx, id.String(), "canceled")

	s.Error(err)
	var sysErr domain.SystemError
	s.True(errors.As(err, &sysErr))
	s.redemptionRepo.AssertNotCalled(s.T(), "Update", mock.Anything, mock.Anything)
}

func (s *RedemptionServiceTestSuite) TestUpdateStatus_NonCancelTransition_NoRefund() {
	ctx := context.Background()
	id := uuid.New()
	redemption := &domain.Redemption{ID: id, Status: domain.RedemptionStatusPending}

	s.redemptionRepo.On("GetByID", mock.Anything, id).Return(redemption, nil)
	s.redemptionRepo.On("Update", ctx, mock.MatchedBy(func(r *domain.Redemption) bool {
		return r.Status == "completed"
	})).Return(nil)

	err := s.service.UpdateStatus(ctx, id.String(), "completed")

	s.NoError(err)
	s.rewardsRepo.AssertNotCalled(s.T(), "GetByID", mock.Anything, mock.Anything)
	s.pointsService.AssertNotCalled(s.T(), "EarnPoints", mock.Anything, mock.Anything)
}

func (s *RedemptionServiceTestSuite) TestUpdateStatus_UpdateRepoError() {
	ctx := context.Background()
	id := uuid.New()
	redemption := &domain.Redemption{ID: id, Status: domain.RedemptionStatusPending}

	s.redemptionRepo.On("GetByID", mock.Anything, id).Return(redemption, nil)
	s.redemptionRepo.On("Update", ctx, mock.Anything).Return(errors.New("db down"))

	err := s.service.UpdateStatus(ctx, id.String(), "completed")

	s.Error(err)
	var sysErr domain.SystemError
	s.True(errors.As(err, &sysErr))
}

func (s *RedemptionServiceTestSuite) TestSetPointsService() {
	newSvc := new(mockPointsServiceIface)
	s.service.SetPointsService(newSvc)
	s.Equal(newSvc, s.service.pointsService)
}
