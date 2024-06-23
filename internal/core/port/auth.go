package port

import (
	"context"
	"hexagonal-todo/internal/core/domain"
)

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*domain.UserDataWithID, error)
	GetUserById(ctx context.Context, id int64) (*domain.UserDataWithID, error)
	CreateUser(ctx context.Context, user *domain.UserData) (*domain.UserDataWithID, error)
}

type AuthService interface {
	Login(ctx context.Context, credential *domain.LoginCredential) (*domain.LoginResponse, error)
	Register(ctx context.Context, userData *domain.UserData) (*domain.UserDataWithID, error)
	ValidateToken(ctx context.Context, token string) (*domain.TokenData, error)
	MyProfile(ctx context.Context) (*domain.UserDataWithID, error)
}

type TokenManager interface {
	GenerateToken(ctx context.Context, token *domain.TokenData) (*string, error)
	DecodeToken(ctx context.Context, token string) (*domain.TokenData, error)
}
