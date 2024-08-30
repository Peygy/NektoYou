package managers

import (
	"errors"
	"testing"

	"github.com/peygy/nektoyou/internal/services/auth_service/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	user_serviceName = "IUserManager"
)

type PasswordManagerMock struct {
	mock.Mock
}

func (m *PasswordManagerMock) hashPassword(password string) (string, error) {
	args := m.Called(password)
	return args.String(0), args.Error(1)
}

func (m *PasswordManagerMock) validPassword(password string) error {
	args := m.Called(password)
	return args.Error(0)
}

func (m *PasswordManagerMock) checkPasswordHash(password, hash string) bool {
	args := m.Called(password, hash)
	return args.Bool(0)
}

func TestNewUserManager_Success(t *testing.T) {
	// Arrange
	db := dbs[user_serviceName]
	mockLog := new(mocks.LoggerMock)
	mockLog.On("Info", "PasswordManager created").Return()
	mockLog.On("Info", "UserManager created").Return()

	// Act
	userManger := NewUserManager(db, mockLog)

	// Accert
	assert.NotNil(t, userManger, "Expected not nil when creating user manager")
}

func TestInsertUser_Success(t *testing.T) {
	// Arrange
	db := dbs[user_serviceName]
	mockLog := new(mocks.LoggerMock)
	mockPM := new(PasswordManagerMock)
	mockPM.On("validPassword", mock.AnythingOfType("string")).Return(nil)
	mockPM.On("hashPassword", mock.AnythingOfType("string")).Return("", nil)
	mockLog.On("Infof", "User %s created successfully", mock.Anything).Return()

	testUserName := "testUserName1"

	// Act
	userManger := &userManger{
		db:               db,
		log:              mockLog,
		iPasswordManager: mockPM,
	}
	userId, err := userManger.InsertUser(UserRecord{UserName: testUserName})

	// Accert
	assert.NoError(t, err, "Expected no error when inserting new user")
	mockPM.AssertCalled(t, "validPassword", "")
	mockPM.AssertCalled(t, "hashPassword", "")

	var dbUserId string
	query := `SELECT id FROM users WHERE username = $1`
	err = db.QueryRow(query, testUserName).Scan(&dbUserId)
	assert.NoError(t, err, "Expected no error when gets querying role")
	assert.Equal(t, userId, dbUserId, "Expected equal between return userId and userId from database")
}

func TestInsertUser_NoValidPassword(t *testing.T) {
	// Arrange
	db := dbs[user_serviceName]
	mockLog := new(mocks.LoggerMock)
	mockPM := new(PasswordManagerMock)
	mockPM.On("validPassword", mock.AnythingOfType("string")).Return(errors.New("password is not valid"))

	// Act
	userManger := &userManger{
		db:               db,
		log:              mockLog,
		iPasswordManager: mockPM,
	}
	_, err := userManger.InsertUser(UserRecord{})

	// Accert
	assert.Error(t, err, "Expected an error when inserting new user")
	mockPM.AssertCalled(t, "validPassword", "")
}

func TestInsertUser_NoHashedPassword(t *testing.T) {
	// Arrange
	db := dbs[user_serviceName]
	mockLog := new(mocks.LoggerMock)
	mockPM := new(PasswordManagerMock)
	mockPM.On("validPassword", mock.AnythingOfType("string")).Return(nil)
	mockPM.On("hashPassword", mock.AnythingOfType("string")).Return("", errors.New("password is not hashed"))

	// Act
	userManger := &userManger{
		db:               db,
		log:              mockLog,
		iPasswordManager: mockPM,
	}
	_, err := userManger.InsertUser(UserRecord{})

	// Accert
	assert.Error(t, err, "Expected an error when inserting new user")
	mockPM.AssertCalled(t, "validPassword", "")
	mockPM.AssertCalled(t, "hashPassword", "")
}

func TestInsertUser_ErrorInsertUser(t *testing.T) {
	// Arrange
	db := dbs[user_serviceName]
	mockLog := new(mocks.LoggerMock)
	mockPM := new(PasswordManagerMock)
	mockPM.On("validPassword", mock.AnythingOfType("string")).Return(nil)
	mockPM.On("hashPassword", mock.AnythingOfType("string")).Return("", nil)
	mockLog.On("Errorf", "Can't adds new user (%s, %s) to the database: %v", mock.Anything, mock.Anything, mock.Anything).Return()

	testUserName := "testUserName2"

	// Act
	userManger := &userManger{
		db:               db,
		log:              mockLog,
		iPasswordManager: mockPM,
	}
	query := `INSERT INTO users (username, password_hash) VALUES ($1, $2)`
	_, err := db.Exec(query, testUserName, "")
	assert.NoError(t, err, "Expected no error when inserting new user")
	_, err = userManger.InsertUser(UserRecord{UserName: testUserName})

	// Accert
	assert.Error(t, err, "Expected an error when inserting new user")
	mockPM.AssertCalled(t, "validPassword", "")
	mockPM.AssertCalled(t, "hashPassword", "")
}
