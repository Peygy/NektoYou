package jwt

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/peygy/nektoyou/internal/pkg/logger"
	"github.com/peygy/nektoyou/internal/services/token_service/config"
)

var (
	ErrInvalidToken = errors.New("jwt: invalid access token")
	ErrExpiredToken = errors.New("jwt: expired access token")
)

type ITokenManager interface {
	NewAccessToken(userId string, ttl time.Duration, roles ...string) (string, error)
	NewRefreshToken() (string, error)
	VerifyAccessToken(accessToken string) (*JwtClaims, error)
}

type tokenManager struct {
	secretKey string
	log       logger.ILogger
}

type JwtClaims struct {
	UserId string   `json:"user_id"`
	Roles  []string `json:"roles"`
	jwt.StandardClaims
}

func NewTokenManager(tknCfg *config.TokenConfig, logger logger.ILogger) ITokenManager {
	logger.Info("TokenManager created")
	return &tokenManager{secretKey: tknCfg.SecretKey, log: logger}
}

func (m *tokenManager) NewAccessToken(userId string, ttl time.Duration, roles ...string) (string, error) {
	if userId == "" {
		return "", errors.New("jwt: userId cannot be empty")
	}

	if ttl <= 0 {
		return "", errors.New("jwt: token ttl must be greater than zero")
	}

	claims := JwtClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ttl).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: userId,
		Roles:  roles,
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

	m.log.Info("Refresh token created successfully")
	return fmt.Sprintf("%x", buffer), nil
}

func (m *tokenManager) VerifyAccessToken(accessToken string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			m.log.Errorf("Unexpected signing method: %v", token.Header["alg"])
			return nil, ErrInvalidToken
		}
		return []byte(m.secretKey), nil
	})

	if err != nil {
		if validationErr, ok := err.(*jwt.ValidationError); ok {
			if validationErr.Errors&jwt.ValidationErrorExpired != 0 {
				m.log.Errorf("Access token is expired: %v", err)
				return nil, ErrExpiredToken
			}
			m.log.Errorf("Invalid access token: %v", err)
			return nil, ErrInvalidToken
		}

		m.log.Errorf("Failed to parse access token: %v", err)
		return nil, ErrInvalidToken
	}

	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
		m.log.Infof("Access token parsed successfully for user %s", claims.UserId)
		return claims, nil
	}

	m.log.Error("Invalid access token claims")
	return nil, ErrInvalidToken
}
