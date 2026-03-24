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

type BusinessProfileService interface {
	Create(ctx context.Context, userID uuid.UUID, request requests.CreateBusinessProfileRequest) (*responses.BusinessProfileResponse, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) (*responses.BusinessProfileResponse, error)
	Update(ctx context.Context, userID uuid.UUID, request requests.UpdateBusinessProfileRequest) (*responses.BusinessProfileResponse, error)
}

type businessProfileService struct {
	businessRepo repository.BusinessProfileRepository
}

func NewBusinessProfileService(businessRepo repository.BusinessProfileRepository) BusinessProfileService {
	return &businessProfileService{businessRepo: businessRepo}
}

func (bs *businessProfileService) Create(ctx context.Context, userID uuid.UUID, request requests.CreateBusinessProfileRequest) (*responses.BusinessProfileResponse, error) {
	existing, err := bs.businessRepo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to check existing profile: %w", err)
	}
	if existing != nil {
		return nil, fmt.Errorf("business profile already exists for this account")
	}
	profile := &models.BusinessProfile{
		UserID:          userID.String(),
		BusinessName:    request.BusinessName,
		BusinessAddress: request.BusinessAddress,
	}
	if err != nil {
		return nil, fmt.Errorf("failed to create business profile: %w", err)
	}
	response := utils.ToBusinessProfileResponse(*profile)
	return &response, nil
}

func (bs *businessProfileService) GetByUserID(ctx context.Context, userID uuid.UUID) (*responses.BusinessProfileResponse, error) {
	profile, err := bs.businessRepo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch profile: %w", err)
	}
	if profile == nil {
		return nil, errors.New("business profile not found")
	}

	resp := utils.ToBusinessProfileResponse(*profile)
	return &resp, nil
}

func (bs *businessProfileService) Update(ctx context.Context, userID uuid.UUID, request requests.UpdateBusinessProfileRequest) (*responses.BusinessProfileResponse, error) {
	profile, err := bs.businessRepo.FindByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch profile: %w", err)
	}
	if profile == nil {
		return nil, errors.New("business profile not found")
	}

	if request.BusinessName != "" {
		profile.BusinessName = request.BusinessName
	}
	if request.BusinessAddress != "" {
		profile.BusinessAddress = request.BusinessAddress
	}
	if request.SubscriptionPlan != "" {
		profile.SubscriptionPlan = request.SubscriptionPlan
	}

	if err := bs.businessRepo.Update(ctx, profile); err != nil {
		return nil, fmt.Errorf("failed to update profile: %w", err)
	}

	resp := utils.ToBusinessProfileResponse(*profile)
	return &resp, nil
}
