package tokenManager

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"hexagonal-todo/internal/core/domain"
	"hexagonal-todo/internal/core/port"
	"sync"
)

// dummyTokenManager should not be used!!
type dummyTokenManager struct {
	tokenMap sync.Map
}

func NewDummyTokenManager() port.TokenManager {
	return &dummyTokenManager{
		tokenMap: sync.Map{},
	}
}

func (d *dummyTokenManager) GenerateToken(_ context.Context, token *domain.TokenData) (*string, error) {
	generatedKey := d.genKey()
	d.tokenMap.Store(generatedKey, token)
	return &generatedKey, nil
}

func (d *dummyTokenManager) DecodeToken(_ context.Context, token string) (*domain.TokenData, error) {
	val, ok := d.tokenMap.Load(token)
	if !ok {
		return nil, domain.AppErrInvalidToken
	}
	return val.(*domain.TokenData), nil
}

func (d *dummyTokenManager) genKey() string {
	buf := make([]byte, 32)
	_, err := rand.Read(buf)
	if err != nil {
		panic(err) // Out of randomness, should never happen
	}
	return hex.EncodeToString(buf)
}
