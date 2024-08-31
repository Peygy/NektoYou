package jwt_test

import (
	"testing"
	"time"

	"github.com/peygy/nektoyou/internal/services/auth_service/config"
	"github.com/peygy/nektoyou/internal/services/auth_service/internal/services/jwt"
	"github.com/peygy/nektoyou/internal/services/auth_service/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	roles = []string{"admin", "user"}
)

func TestNewAccessToken_Success(t *testing.T) {

}

func TestNewAccessToken_Error(t *testing.T) {

}

func TestNewRefreshToken_Success(t *testing.T) {

}

func TestNewRefreshToken_Error(t *testing.T) {

}

func TestVerifyAccessToken_Success(t *testing.T) {
	// Arrange
	tokenConfig := &config.TokenManagerConfig{SecretKey: "secretKey"}
	mockLog := new(mocks.LoggerMock)
	mockLog.On("Infof", "Access token of user %s created successfully", mock.Anything).Return()
	mockLog.On("Infof", "Access token parsed successfully for user %s", mock.Anything).Return()

	testUserId := "testUserId"

	// Act
	jwtService := jwt.NewTokenManager(tokenConfig, mockLog)
	accessToken, err := jwtService.NewAccessToken(testUserId, 1*time.Hour, roles...)
	assert.NoError(t, err, "Expected no error when create new access token")
	claims, err := jwtService.VerifyAccessToken(accessToken)

	// Accert
	assert.NoError(t, err, "Expected no error when create new access token")
	assert.Equal(t, testUserId, claims.UserId, "Expected an equal between userId and userId from claims")
}

func TestVerifyAccessToken_InvalidSigningMethod(t *testing.T) {

}

func TestVerifyAccessToken_ExpiredToken(t *testing.T) {

}

func TestVerifyAccessToken_InvalidToken(t *testing.T) {

}

func TestVerifyAccessToken_InvalidClaims(t *testing.T) {

}
