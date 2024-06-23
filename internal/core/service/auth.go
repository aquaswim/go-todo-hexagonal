package service

import (
	"context"
	"hexagonal-todo/internal/core/domain"
	"hexagonal-todo/internal/core/helpers"
	"hexagonal-todo/internal/core/port"
)

type authService struct {
	userRepo     port.UserRepository
	tokenManager port.TokenManager
}

func NewAuthService(
	userRepo port.UserRepository,
	tokenManager port.TokenManager,
) port.AuthService {
	return &authService{
		userRepo:     userRepo,
		tokenManager: tokenManager,
	}
}

func (a authService) Login(ctx context.Context, credential *domain.LoginCredential) (*domain.LoginResponse, error) {
	// get user by email
	user, err := a.userRepo.GetUserByEmail(ctx, credential.Email)
	if err != nil {
		return nil, domain.AppErrUserNotFound
	}

	// compare password with bcrypt
	passwordValid, err := helpers.ComparePasswordWithHash(credential.Password, user.Password)
	if !passwordValid {
		return nil, domain.AppErrUserNotFound
	}
	if err != nil {
		return nil, domain.FromError(domain.ErrCodeInternal, err)
	}

	token, err := a.tokenManager.GenerateToken(ctx, &domain.TokenData{Id: user.Id})
	if err != nil {
		return nil, domain.FromError(domain.ErrCodeInternal, err)
	}

	return &domain.LoginResponse{
		Token: *token,
	}, nil
}

func (a authService) Register(ctx context.Context, userData *domain.UserData) (*domain.UserDataWithID, error) {
	// check for email existance
	_, err := a.userRepo.GetUserByEmail(ctx, userData.Email)
	if err == nil && !domain.ErrCodeNotFound.IsErrEqual(err) {
		return nil, domain.AppErrEmailAlreadyExists
	}

	// hash password
	hashedPassword, err := helpers.HashPassword(userData.Password)
	if err != nil {
		return nil, domain.FromError(domain.ErrCodeInternal, err)
	}
	// insert to db
	createdUser, err := a.userRepo.CreateUser(ctx, &domain.UserData{
		LoginCredential: domain.LoginCredential{
			Email:    userData.Email,
			Password: hashedPassword,
		},
		FullName: userData.FullName,
	})
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}

func (a authService) ValidateToken(ctx context.Context, token string) (*domain.TokenData, error) {
	decodeToken, err := a.tokenManager.DecodeToken(ctx, token)
	if err != nil {
		return nil, err
	}
	return decodeToken, nil
}

func (a authService) MyProfile(ctx context.Context) (*domain.UserDataWithID, error) {
	authCtx, err := helpers.GetAuthCtx(ctx)
	if err != nil {
		return nil, err
	}
	userData, err := a.userRepo.GetUserById(ctx, authCtx.Id)
	if err != nil {
		return nil, err
	}
	return userData, nil
}
