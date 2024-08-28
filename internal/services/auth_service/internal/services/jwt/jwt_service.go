package jwt

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/peygy/nektoyou/internal/pkg/logger"
	"github.com/peygy/nektoyou/internal/services/auth_service/config"
)

type ITokenManager interface {
	NewAccessToken(userId string, ttl time.Duration) (string, error)
	NewRefreshToken() (string, error)
}

type tokenManager struct {
	secretKey string
	logger    logger.ILogger
}

func NewTokenManager(tknCfg *config.TokenManagerConfig, logger logger.ILogger) ITokenManager {
	return &tokenManager{secretKey: tknCfg.SecretKey, logger: logger}
}

func (m *tokenManager) NewAccessToken(userId string, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(ttl).Unix(),
		Subject:   userId,
	})

	return token.SignedString([]byte(m.secretKey))
}

func (m *tokenManager) NewRefreshToken() (string, error) {
	buffer := make([]byte, 32)

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	_, err := r.Read(buffer)
	if err != nil {
		m.logger.Error("error during creation of refresh token: " + err.Error())
		return "", err
	}

	return fmt.Sprintf("%x", buffer), nil
}
