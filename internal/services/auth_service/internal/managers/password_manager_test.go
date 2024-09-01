package managers

import (
	"testing"

	"math/rand"

	"github.com/peygy/nektoyou/internal/pkg/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

func randomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

var (
	_password = "password123"
)

func TestHashPassword_Success(t *testing.T) {
	// Arrange
	mockLog := new(mocks.LoggerMock)
	mockLog.On("Info", "PasswordManager created").Return()
	mockLog.On("Info", "Password is hashed successfully").Return()

	// Act
	passwordManager := newPasswordManager(7, mockLog)
	hash, err := passwordManager.hashPassword(_password)

	// Assert
	assert.NoError(t, err, "Expected no error during hashing")
	assert.NotEmpty(t, hash, "Expected a non-empty hashed password")

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(_password))
	assert.NoError(t, err, "Expected bcrypt to confirm the password hash")
}

func TestHashPassword_Error(t *testing.T) {
	// Arrange
	mockLog := new(mocks.LoggerMock)
	mockLog.On("Info", "PasswordManager created").Return()
	mockLog.On("Errorf", "Can't create hashed password with error: %v", mock.Anything).Return()

	password := randomString(73)

	// Act
	passwordManager := newPasswordManager(7, mockLog)
	hash, err := passwordManager.hashPassword(password)

	// Assert
	assert.Error(t, err, "Expected an error during hashing")
	assert.Empty(t, hash, "Expected an empty hashed password when an error occurs")
}

func TestValidPassword_Success(t *testing.T) {
	// Arrange
	mockLog := new(mocks.LoggerMock)
	mockLog.On("Info", "PasswordManager created").Return()
	mockLog.On("Info", "Password is valided").Return()

	// Act
	passwordManager := newPasswordManager(7, mockLog)
	err := passwordManager.validPassword(_password)

	// Assert
	assert.NoError(t, err, "Expected no error due to password being too short")
}

func TestValidPassword_Error(t *testing.T) {
	// Arrange
	mockLog := new(mocks.LoggerMock)
	mockLog.On("Info", "PasswordManager created").Return()
	mockLog.On("Error", "Password length less than minimum length").Return()

	// Act
	passwordManager := newPasswordManager(15, mockLog)
	err := passwordManager.validPassword(_password)

	// Assert
	assert.Error(t, err, "Expected an error due to password being too short")
}

func TestCheckPasswordHash_True(t *testing.T) {
	// Arrange
	mockLog := new(mocks.LoggerMock)
	mockLog.On("Info", "PasswordManager created").Return()
	bytes, _ := bcrypt.GenerateFromPassword([]byte(_password), 14)

	// Act
	passwordManager := newPasswordManager(7, mockLog)
	result := passwordManager.checkPasswordHash(_password, string(bytes))

	// Assert
	assert.True(t, result, "Expected true after comparing passwords")
}

func TestCheckPasswordHash_False(t *testing.T) {
	// Arrange
	mockLog := new(mocks.LoggerMock)
	mockLog.On("Info", "PasswordManager created").Return()

	// Act
	passwordManager := newPasswordManager(7, mockLog)
	result := passwordManager.checkPasswordHash(_password, string([]byte("test")))

	// Assert
	assert.False(t, result, "Expected false after comparing passwords")
}
