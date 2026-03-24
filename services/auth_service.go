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
	"time"

	"github.com/google/uuid"
)

type AuthService interface {
	Register(ctx context.Context, request requests.RegisterRequest) (*responses.AuthResponse, error)
	Login(ctx context.Context, request requests.LoginRequest) (*responses.AuthResponse, error)
	ChangePassword(ctx context.Context, userID uuid.UUID, request requests.ChangePasswordRequest) error
	ForgetPassword(ctx context.Context, request requests.ForgotPasswordRequest) error
	ResetPassword(ctx context.Context, request requests.ResetPasswordRequest) error
}

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{userRepo: userRepo}
}

func (auth *authService) Register(ctx context.Context, request requests.RegisterRequest) (*responses.AuthResponse, error) {
	emailExists, err := auth.userRepo.ExistsByEmail(ctx, request.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to check email: %w", err)
	}
	if emailExists {
		return nil, errors.New("an account with this email already exists")
	}
	phoneExists, err := auth.userRepo.ExistsByPhone(ctx, request.Phone)
	if err != nil {
		return nil, fmt.Errorf("failed to check phone: %w", err)
	}
	if phoneExists {
		return nil, errors.New("an account with this phone already exists")
	}
	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}
	user := &models.User{
		Name:     request.Name,
		Email:    request.Email,
		Phone:    request.Phone,
		Password: hashedPassword,
		Role:     request.Role,
		Status:   models.Active,
	}
	if err := auth.userRepo.Create(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	return auth.buildAuthResponse(user)
}

func (auth *authService) Login(ctx context.Context, request requests.LoginRequest) (*responses.AuthResponse, error) {
	user, err := auth.userRepo.FindByEmail(ctx, request.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user: %w", err)
	}
	if user == nil || utils.CheckPassword(request.Password, user.Password) != nil {
		return nil, errors.New("invalid email or password")
	}
	if user.Status == models.Suspended {
		return nil, errors.New("your account has been suspended, pls contact your support")
	}
	if user.Status == models.Inactive {
		return nil, errors.New("your account is inactive, please contact support")
	}
	return auth.buildAuthResponse(user)
}

func (auth *authService) ChangePassword(ctx context.Context, userID uuid.UUID, request requests.ChangePasswordRequest) error {
	user, err := auth.userRepo.FindById(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to fetch user: %w", err)
	}
	if user == nil {
		return errors.New("user not found")
	}

	if utils.CheckPassword(request.CurrentPassword, user.Password) != nil {
		return errors.New("current password is incorrect")
	}

	hashed, err := utils.HashPassword(request.NewPassword)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	user.Password = hashed
	return auth.userRepo.Update(ctx, user)
}

func (auth *authService) ForgetPassword(ctx context.Context, request requests.ForgotPasswordRequest) error {
	user, err := auth.userRepo.FindByEmail(ctx, request.Email)
	if err != nil {
		return fmt.Errorf("failed to fetch user: %w", err)
	}
	if user == nil {
		return nil
	}
	otp := utils.GenerateOTP()
	key := fmt.Sprintf("otp:%s", request.Email)
	if err := utils.SetString(key, otp, 15*time.Minute); err != nil {
		return fmt.Errorf("failed to store otp: %w", err)
	}
	fmt.Printf("DEBUG OTP for %s: %s\n", request.Email, otp)
	return nil
}

func (auth *authService) ResetPassword(ctx context.Context, request requests.ResetPasswordRequest) error {
	key := fmt.Sprintf("otp:%s", request.Email)
	stored, found, err := utils.GetString(key)
	if err != nil {
		return fmt.Errorf("failed to retrieve otp: %w", err)
	}
	if !found || stored != request.OTP {
		return errors.New("invalid or expired otp")
	}
	user, err := auth.userRepo.FindByEmail(ctx, request.Email)
	if err != nil {
		return fmt.Errorf("failed to fetch user: %w", err)
	}
	if user == nil {
		return errors.New("user not found")
	}
	hashed, err := utils.HashPassword(request.NewPassword)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}
	user.Password = hashed
	if err := auth.userRepo.Update(ctx, user); err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	_ = utils.DeleteKey(key)
	return nil
}

func (auth *authService) buildAuthResponse(user *models.User) (*responses.AuthResponse, error) {
	token, err := utils.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}
	refreshToken, err := utils.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to generate refresh token: %w", err)
	}
	return &responses.AuthResponse{
		AccessToken:  token,
		RefreshToken: refreshToken,
		ExpiresIn:    24 * 60 * 60,
		User:         utils.ToUserResponse(*user),
	}, nil
}
