package services

import (
	"context"
	"errors"
	"fmt"
	"logisticApp/data/models"
	"logisticApp/data/repository"
	"logisticApp/dtos/requests"
	"logisticApp/dtos/responses"
	"logisticApp/utils"

	"github.com/google/uuid"
)

type RiderProfileService interface {
	Create(ctx context.Context, userID uuid.UUID, request requests.CreateRiderProfileRequest) (*responses.RiderProfileResponse, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) (*responses.RiderProfileResponse, error)
	Update(ctx context.Context, userID uuid.UUID, request requests.UpdateRiderProfileRequest) (*responses.RiderProfileResponse, error)
	Approve(ctx context.Context, riderID uuid.UUID) error
	ListAvailable(ctx context.Context) ([]responses.RiderPublicResponse, error)
}

type riderProfileService struct {
	riderRepo repository.RiderProfileRepository
}

func NewRiderProfileService(riderRepo repository.RiderProfileRepository) RiderProfileService {
	return &riderProfileService{riderRepo: riderRepo}
}

func (rider *riderProfileService) Create(ctx context.Context, userID uuid.UUID, request requests.CreateRiderProfileRequest) (*responses.RiderProfileResponse, error) {
	existing, err := rider.riderRepo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to check existing profile: %w\", err")
	}
	if existing != nil {
		return nil, fmt.Errorf("rider profile already exists for this account")
	}
	profile := &models.RiderProfile{
		UserID:        userID.String(),
		VehicleType:   request.VehicleType,
		LicenseNumber: request.LicenseNumber,
		Status:        models.Pending,
	}
	if err := rider.riderRepo.Create(ctx, profile); err != nil {
		return nil, fmt.Errorf("failed to create rider profile: %w", err)
	}
	response := utils.ToRiderProfileResponse(*profile)
	return &response, nil
}

func (rider *riderProfileService) GetByUserID(ctx context.Context, userID uuid.UUID) (*responses.RiderProfileResponse, error) {
	profile, err := rider.riderRepo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch profile: %w", err)
	}
	if profile == nil {
		return nil, fmt.Errorf("rider profile not found")
	}
	response := utils.ToRiderProfileResponse(*profile)
	return &response, nil
}

func (rider *riderProfileService) Update(ctx context.Context, userID uuid.UUID, request requests.UpdateRiderProfileRequest) (*responses.RiderProfileResponse, error) {
	profile, err := rider.riderRepo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch profile: %w", err)
	}
	if profile == nil {
		return nil, fmt.Errorf("rider profile not found: %w", err)
	}
	if request.VehicleType != "" {
		profile.VehicleType = request.VehicleType
	}
	if request.LicenseNumber != "" {
		profile.LicenseNumber = request.LicenseNumber
	}
	if err := rider.riderRepo.Update(ctx, profile); err != nil {
		return nil, fmt.Errorf("failed to update profile: %w", err)
	}
	response := utils.ToRiderProfileResponse(*profile)
	return &response, nil
}

func (rider *riderProfileService) Approve(ctx context.Context, riderID uuid.UUID) error {
	profile, err := rider.riderRepo.FindByID(ctx, riderID)
	if err != nil {
		return fmt.Errorf("failed to fetch profile: %w", err)
	}
	if profile == nil {
		return fmt.Errorf("rider profile not found")
	}
	if profile.Status == models.Active {
		return errors.New("rider is already approved")
	}
	return rider.riderRepo.UpdateStatus(ctx, riderID, models.Active)
}

func (rider *riderProfileService) ListAvailable(ctx context.Context) ([]responses.RiderPublicResponse, error) {
	riders, err := rider.riderRepo.ListAvailable(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list riders: %w", err)
	}

	response := make([]responses.RiderPublicResponse, len(riders))
	for i, r := range riders {
		response[i] = utils.ToRiderPublicResponse(r, r.User)
	}
	return response, nil
}
