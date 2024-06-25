package tokenManager

import (
	"context"
	"github.com/golang-jwt/jwt"
	"github.com/rs/zerolog/log"
	"hexagonal-todo/internal/adapter/config"
	"hexagonal-todo/internal/core/domain"
	"hexagonal-todo/internal/core/port"
	"strconv"
	"time"
)

type jwtTokenManager struct {
	cfg *config.JwtConfig
}

func NewJwtTokenManager(cfg *config.JwtConfig) port.TokenManager {
	log.Debug().Msg("initializing jwt token manager")

	return &jwtTokenManager{
		cfg: cfg,
	}
}

func (j jwtTokenManager) GenerateToken(_ context.Context, tokenData *domain.TokenData) (*string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(j.cfg.Duration).Unix(),
		//Id:        "",
		IssuedAt: time.Now().Unix(),
		//Issuer:    "",
		//NotBefore: 0,
		Subject: strconv.FormatInt(tokenData.Id, 10),
	})
	signedString, err := t.SignedString([]byte(j.cfg.Secret))
	if err != nil {
		return nil, domain.FromError(domain.ErrCodeInternal, err)
	}
	return &signedString, nil
}

func (j jwtTokenManager) DecodeToken(_ context.Context, tokenStr string) (*domain.TokenData, error) {
	claims := &jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.cfg.Secret), nil
	})
	if err != nil {
		return nil, domain.FromError(domain.ErrCodeForbidden, err)
	}
	if !token.Valid {
		return nil, domain.AppErrInvalidToken
	}
	id, err := strconv.ParseInt(claims.Subject, 10, 64)
	if err != nil {
		return nil, domain.FromError(domain.ErrCodeInternal, err)
	}
	return &domain.TokenData{
		Id: id,
	}, nil
}
