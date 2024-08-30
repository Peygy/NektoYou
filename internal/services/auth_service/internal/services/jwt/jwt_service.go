package jwt

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/peygy/nektoyou/internal/pkg/logger"
	"github.com/peygy/nektoyou/internal/services/auth_service/config"
)

type ITokenManager interface {
	NewAccessToken(userId string, ttl time.Duration, roles ...string) (string, error)
	NewRefreshToken() (string, error)
}

type tokenManager struct {
	secretKey string
	log       logger.ILogger
}

type customClaims struct {
	userId string   `json:"user_id"`
	roles  []string `json:"roles"`
	jwt.StandardClaims
}

func NewTokenManager(tknCfg *config.TokenManagerConfig, logger logger.ILogger) ITokenManager {
	return &tokenManager{secretKey: tknCfg.SecretKey, log: logger}
}

func (m *tokenManager) NewAccessToken(userId string, ttl time.Duration, roles ...string) (string, error) {
	claims := customClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ttl).Unix(),
			IssuedAt:  time.Now().Unix(),
			Subject:   userId,
		},
		userId: userId,
		roles:  roles,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := token.SignedString([]byte(m.secretKey))
	if err != nil {
		m.log.Errorf("Can't creates access token: %v", err)
		return "", errors.New("jwt: can't creates access token")
	}

	m.log.Infof("Access token of user %s created successfully", userId)
	return accessToken, nil
}

func (m *tokenManager) NewRefreshToken() (string, error) {
	buffer := make([]byte, 32)

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	_, err := r.Read(buffer)
	if err != nil {
		m.log.Errorf("Can't creates refresh token: %v", err)
		return "", errors.New("jwt: can't creates refresh token")
	}

	m.log.Infof("Refresh token created successfully")
	return fmt.Sprintf("%x", buffer), nil
}
