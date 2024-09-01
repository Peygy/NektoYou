package jwt_test

import (
	"errors"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/peygy/nektoyou/internal/pkg/mocks"
	"github.com/peygy/nektoyou/internal/services/token_service/config"
	jwtService "github.com/peygy/nektoyou/internal/services/token_service/internal/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	roles     = []string{"admin", "user"}
	secretKey = "secretKey"
)

type fakeJwtClaims struct {
	Id string `json:"id"`
	jwt.StandardClaims
}

func parseAccessToken(t *testing.T, accessToken string) error {
	_, err := jwt.ParseWithClaims(accessToken, &jwtService.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwtService.ErrInvalidToken
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, jwtService.ErrExpiredToken) {
			return jwtService.ErrExpiredToken
		}

		return jwtService.ErrInvalidToken
	}

	return nil
}

func generateToken(secretKey string, userId string, ttl time.Duration, roles ...string) string {
	claims := jwtService.JwtClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ttl).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: userId,
		Roles:  roles,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, _ := token.SignedString([]byte(secretKey))
	return accessToken
}

func TestNewAccessToken_Success(t *testing.T) {
	// Arrange
	tokenConfig := &config.TokenConfig{SecretKey: secretKey}
	mockLog := new(mocks.LoggerMock)
	mockLog.On("Info", "TokenManager created").Return()
	mockLog.On("Infof", "Access token of user %s created successfully", mock.Anything).Return()
	mockLog.On("Infof", "Access token parsed successfully for user %s", mock.Anything).Return()

	testUserId := "testUserId"

	// Act
	jwtManager := jwtService.NewTokenManager(tokenConfig, mockLog)
	accessToken, err := jwtManager.NewAccessToken(testUserId, 1*time.Hour, roles...)
	_, verifyErr := jwtManager.VerifyAccessToken(accessToken)
	assert.NoError(t, verifyErr, "Expected no error when verify access token")

	// Accert
	assert.NoError(t, err, "Expected no error when create new access token")
	assert.NoError(t, parseAccessToken(t, accessToken), "Expected no error when parse new access token")
}

func TestNewAccessToken_ErrorUserIdIsEmpty(t *testing.T) {
	// Arrange
	tokenConfig := &config.TokenConfig{SecretKey: secretKey}
	mockLog := new(mocks.LoggerMock)
	mockLog.On("Info", "TokenManager created").Return()

	testUserId := ""

	// Act
	jwtManager := jwtService.NewTokenManager(tokenConfig, mockLog)
	accessToken, err := jwtManager.NewAccessToken(testUserId, 1*time.Hour, roles...)

	// Accert
	assert.Error(t, err, "Expected an error when create new access token")
	assert.Empty(t, accessToken, "Expected empty when create new access token")
}

func TestNewAccessToken_ErrorDurationIsLessThanZero(t *testing.T) {
	// Arrange
	tokenConfig := &config.TokenConfig{SecretKey: secretKey}
	mockLog := new(mocks.LoggerMock)
	mockLog.On("Info", "TokenManager created").Return()

	testUserId := "testUserId"

	// Act
	jwtManager := jwtService.NewTokenManager(tokenConfig, mockLog)
	accessToken, err := jwtManager.NewAccessToken(testUserId, -1*time.Hour, roles...)

	// Accert
	assert.Error(t, err, "Expected an error when create new access token")
	assert.Empty(t, accessToken, "Expected empty when create new access token")
}

func TestNewRefreshToken_Success(t *testing.T) {
	// Arrange
	tokenConfig := &config.TokenConfig{SecretKey: secretKey}
	mockLog := new(mocks.LoggerMock)
	mockLog.On("Info", "TokenManager created").Return()
	mockLog.On("Info", "Refresh token created successfully").Return()

	// Act
	jwtManager := jwtService.NewTokenManager(tokenConfig, mockLog)
	refreshToken, err := jwtManager.NewRefreshToken()

	// Accert
	assert.NoError(t, err, "Expected an error when create new refresh token")
	assert.NotEmpty(t, refreshToken, "Expected no empty when create new refresh token")
}

func TestVerifyAccessToken_Success(t *testing.T) {
	// Arrange
	tokenConfig := &config.TokenConfig{SecretKey: "secretKey"}
	mockLog := new(mocks.LoggerMock)
	mockLog.On("Info", "TokenManager created").Return()
	mockLog.On("Infof", "Access token parsed successfully for user %s", mock.Anything).Return()

	testUserId := "testUserId"

	// Act
	jwtManager := jwtService.NewTokenManager(tokenConfig, mockLog)
	claims, err := jwtManager.VerifyAccessToken(
		generateToken(tokenConfig.SecretKey, testUserId, 1*time.Hour, roles...),
	)

	// Accert
	assert.NoError(t, err, "Expected no error when verify access token")
	assert.Equal(t, testUserId, claims.UserId, "Expected an equal between userId and userId from claims")
}

func TestVerifyAccessToken_ExpiredToken(t *testing.T) {
	// Arrange
	tokenConfig := &config.TokenConfig{SecretKey: "secretKey"}
	mockLog := new(mocks.LoggerMock)
	mockLog.On("Info", "TokenManager created").Return()
	mockLog.On("Errorf", "Access token is expired: %v", mock.Anything).Return()

	claims := &jwtService.JwtClaims{
		UserId: "user123",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(-1 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(secretKey))

	// Act
	jwtManager := jwtService.NewTokenManager(tokenConfig, mockLog)
	_, err := jwtManager.VerifyAccessToken(tokenString)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, jwtService.ErrExpiredToken, err)
}

func TestVerifyAccessToken_InvalidToken(t *testing.T) {
	// Arrange
	tokenConfig := &config.TokenConfig{SecretKey: "secretKey"}
	mockLog := new(mocks.LoggerMock)
	mockLog.On("Info", "TokenManager created").Return()
	mockLog.On("Errorf", "Invalid access token: %v", mock.Anything).Return()

	token := "token"

	// Act
	jwtManager := jwtService.NewTokenManager(tokenConfig, mockLog)
	_, err := jwtManager.VerifyAccessToken(token)

	// Accert
	assert.Error(t, err, "Expected no error when verify access token")
	assert.Equal(t, err, jwtService.ErrInvalidToken)
}

func TestVerifyAccessToken_InvalidClaims(t *testing.T) {
	// Arrange
	tokenConfig := &config.TokenConfig{SecretKey: "secretKey"}
	mockLog := new(mocks.LoggerMock)
	mockLog.On("Info", "TokenManager created").Return()
	mockLog.On("Errorf", "Invalid access token: %v", mock.Anything).Return()

	claims := &jwtService.JwtClaims{
		UserId: "user123",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, _ := token.SignedString([]byte("wrongkey"))

	// Act
	jwtManager := jwtService.NewTokenManager(tokenConfig, mockLog)
	_, err := jwtManager.VerifyAccessToken(tokenString)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, jwtService.ErrInvalidToken, err)
}
