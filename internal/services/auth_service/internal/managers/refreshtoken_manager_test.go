package managers

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/peygy/nektoyou/internal/services/auth_service/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	refresh_serviceName = "IRefreshManager"
)

func insertUser(t *testing.T, db *sql.DB, userName, hashedPassword string) string {
	var userId string
	query := `INSERT INTO users (username, password_hash) VALUES ($1, $2) RETURNING id`
	err := db.QueryRow(query, userName, hashedPassword).Scan(&userId)
	assert.NoError(t, err, "Expected no error when gets querying user")
	return userId
}

func checkRefreshToken(t *testing.T, db *sql.DB, userId, testRefreshToken string) {
	var token string
	err := db.QueryRow(`SELECT token FROM users_tokens WHERE user_id = $1`, userId).Scan(&token)
	assert.NoError(t, err, "Expected no error when gets querying token")
	assert.Equal(t, testRefreshToken, token, "Expected equal between refresh token and token from database")
}

func TestAddToken_Success_AddingNewRefreshToken(t *testing.T) {
	// Arrange
	db := dbs[refresh_serviceName]
	mockLog := new(mocks.LoggerMock)
	mockLog.On("Info", "RefreshManager created").Return()
	mockLog.On("Infof", "Refresh token %s successfully added to user %s", mock.Anything, mock.Anything).Return()

	testUserName := "testUser1"
	testHashedPassword := "testPassword1"
	testRefreshToken := "testToken1"

	userId := insertUser(t, db, testUserName, testHashedPassword)

	// Act
	refreshManager := NewRefreshManager(db, mockLog)
	err := refreshManager.AddToken(userId, testRefreshToken)

	// Assert
	assert.NoError(t, err, "Expected no error when adding refresh token")
	checkRefreshToken(t, db, userId, testRefreshToken)
}

func TestAddToken_Success_AddingRefreshTokenToUserIdWithRefreshToken(t *testing.T) {
	// Arrange
	db := dbs[refresh_serviceName]
	mockLog := new(mocks.LoggerMock)
	mockLog.On("Info", "RefreshManager created").Return()
	mockLog.On("Infof", "Refresh token %s successfully added to user %s", mock.Anything, mock.Anything).Return()

	testUserName := "testUser2"
	testHashedPassword := "testPassword2"
	testRefreshToken := "testToken2"
	newTestRefreshToken := "newTestToken2"

	userId := insertUser(t, db, testUserName, testHashedPassword)

	// Act
	refreshManager := NewRefreshManager(db, mockLog)
	err := refreshManager.AddToken(userId, testRefreshToken)
	assert.NoError(t, err, "Expected no error when adding refresh token")
	err = refreshManager.AddToken(userId, newTestRefreshToken)

	// Assert
	assert.NoError(t, err, "Expected no error when adding refresh token")

	checkRefreshToken(t, db, userId, testRefreshToken)
}

func TestAddToken_NoFindUserIdInUsersTable(t *testing.T) {
	// Arrange
	db := dbs[refresh_serviceName]
	mockLog := new(mocks.LoggerMock)
	mockLog.On("Info", "RefreshManager created").Return()
	mockLog.On("Errorf", "Can't inserts refresh token with error: %v", mock.Anything).Return()

	testRefreshToken := "testToken3"
	userId := "testId3"

	// Act
	refreshManager := NewRefreshManager(db, mockLog)
	err := refreshManager.AddToken(userId, testRefreshToken)

	// Assert
	assert.Error(t, err, "Expected an error when adding refresh token")
}

func TestAddToken_DuplicateRefreshToken(t *testing.T) {
	// Arrange
	db := dbs[refresh_serviceName]
	mockLog := new(mocks.LoggerMock)
	mockLog.On("Info", "RefreshManager created").Return()
	mockLog.On("Infof", "Refresh token %s successfully added to user %s", mock.Anything, mock.Anything).Return()

	testUserName := "testUser4"
	testHashedPassword := "testPassword4"
	testRefreshToken := "testToken4"

	userId := insertUser(t, db, testUserName, testHashedPassword)

	// Act
	refreshManager := NewRefreshManager(db, mockLog)
	err := refreshManager.AddToken(userId, testRefreshToken)
	assert.NoError(t, err, "Expected no error when adding refresh token")
	err = refreshManager.AddToken(userId, testRefreshToken)

	// Assert
	assert.NoError(t, err, "Expected no error when adding refresh token")

	checkRefreshToken(t, db, userId, testRefreshToken)
}

func TestGetToken_SuccessGetting(t *testing.T) {
	// Arrange
	db := dbs[refresh_serviceName]
	mockLog := new(mocks.LoggerMock)
	mockLog.On("Info", "RefreshManager created").Return()
	mockLog.On("Info", "Token was founded").Return()

	testUserName := "testUser5"
	testHashedPassword := "testPassword5"
	testRefreshToken := "testToken5"

	userId := insertUser(t, db, testUserName, testHashedPassword)

	// Act
	refreshManager := NewRefreshManager(db, mockLog)

	query := `INSERT INTO users_tokens (user_id, token) VALUES ($1, $2) ON CONFLICT (token) DO NOTHING`
	_, err := db.Exec(query, userId, testRefreshToken)
	assert.NoError(t, err, "Expected no error when adding refresh token")

	refreshToken, err := refreshManager.GetToken(userId)

	// Assert
	assert.NoError(t, err, "Expected no error when getting refresh token")
	assert.Equal(t, testRefreshToken, refreshToken, "Expected equal between refresh token and token from database")
}

func TestGetToken_NoAnyRefreshToken(t *testing.T) {
	// Arrange
	db := dbs[refresh_serviceName]
	mockLog := new(mocks.LoggerMock)
	mockLog.On("Info", "RefreshManager created").Return()
	mockLog.On("Errorf", "No refresh token was found: %v", mock.Anything).Return()

	testUserName := "testUser6"
	testHashedPassword := "testPassword6"

	userId := insertUser(t, db, testUserName, testHashedPassword)

	// Act
	refreshManager := NewRefreshManager(db, mockLog)
	_, err := refreshManager.GetToken(userId)

	// Assert
	assert.Error(t, err, "Expected an error when getting refresh token")
}

func TestGetToken_NoFindUserIdInUsersTable(t *testing.T) {
	// Arrange
	db := dbs[refresh_serviceName]
	mockLog := new(mocks.LoggerMock)
	mockLog.On("Info", "RefreshManager created").Return()
	mockLog.On("Errorf", "Can't gets refresh token: %v", mock.Anything).Return()

	userId := "testUserId7"

	// Act
	refreshManager := NewRefreshManager(db, mockLog)
	_, err := refreshManager.GetToken(userId)

	// Assert
	assert.Error(t, err, "Expected an error when getting refresh token")
}

func TestRefreshToken_Success(t *testing.T) {
	// Arrange
	db := dbs[refresh_serviceName]
	mockLog := new(mocks.LoggerMock)
	mockLog.On("Info", "RefreshManager created").Return()
	mockLog.On("Info", "Refresh-Refresh: Transaction is begining successfully").Return()
	mockLog.On("Infof", "User's %s refresh token was found", mock.Anything).Return()
	mockLog.On("Infof", "User's %s refresh token updated successfully", mock.Anything).Return()
	mockLog.On("Info", "Refresh-Refresh: Transaction is commited successfully").Return()

	testUserName := "testUser8"
	testHashedPassword := "testPassword8"
	testRefreshToken := "testToken8"
	newTestRefreshToken := "newTestToken8"

	userId := insertUser(t, db, testUserName, testHashedPassword)

	query := `INSERT INTO users_tokens (user_id, token) VALUES ($1, $2) ON CONFLICT (token) DO NOTHING`
	_, err := db.Exec(query, userId, testRefreshToken)
	assert.NoError(t, err, "Expected no error when adding refresh token")

	// Act
	refreshManager := NewRefreshManager(db, mockLog)
	result, err := refreshManager.RefreshToken(userId, newTestRefreshToken)

	// Assert
	assert.True(t, result, "Expected true when updating refresh token")
	assert.NoError(t, err, "Expected no error when updating refresh token")

	checkRefreshToken(t, db, userId, newTestRefreshToken)
}

func TestRefreshToken_NoAnyOldRefreshToken(t *testing.T) {
	// Arrange
	db := dbs[refresh_serviceName]
	mockLog := new(mocks.LoggerMock)
	mockLog.On("Info", "RefreshManager created").Return()
	mockLog.On("Info", "Refresh-Refresh: Transaction is begining successfully").Return()
	mockLog.On("Errorf", "No any refresh token in database: %v", mock.Anything).Return()

	testUserName := "testUser9"
	testHashedPassword := "testPassword9"
	newTestRefreshToken := "newTestToken9"

	userId := insertUser(t, db, testUserName, testHashedPassword)

	// Act
	refreshManager := NewRefreshManager(db, mockLog)
	result, err := refreshManager.RefreshToken(userId, newTestRefreshToken)

	// Assert
	assert.False(t, result, "Expected false when updating refresh token")
	assert.Error(t, err, "Expected an error when updating refresh token")
}

func TestRefreshToken_NoFindUserIdInUsersTable(t *testing.T) {
	// Arrange
	db := dbs[refresh_serviceName]
	mockLog := new(mocks.LoggerMock)
	mockLog.On("Info", "RefreshManager created").Return()
	mockLog.On("Info", "Refresh-Refresh: Transaction is begining successfully").Return()
	mockLog.On("Errorf", "Error retrieving old token: %v", mock.Anything).Return()

	userId := "testUserId10"
	newTestRefreshToken := "newTestToken10"

	// Act
	refreshManager := NewRefreshManager(db, mockLog)
	result, err := refreshManager.RefreshToken(userId, newTestRefreshToken)

	// Assert
	assert.False(t, result, "Expected false when updating refresh token")
	assert.Error(t, err, "Expected an error when updating refresh token")
}

func TestRefreshToken_BeginTransactionError(t *testing.T) {
	// Arrange
	mockDB, sqlm, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	mockLog := new(mocks.LoggerMock)
	mockLog.On("Info", "RefreshManager created").Return()
	mockLog.On("Errorf", "Error starting transaction: %v", mock.Anything).Return()

	userId := "testUserId11"
	newTestRefreshToken := "newTestToken11"

	sqlm.ExpectBegin().WillReturnError(errors.New("begin transaction error"))

	// Act
	refreshManager := NewRefreshManager(mockDB, mockLog)
	result, err := refreshManager.RefreshToken(userId, newTestRefreshToken)

	// Assert
	assert.False(t, result, "Expected false when starting transaction fails")
	assert.Error(t, err, "Expected error when starting transaction fails")
}

func TestRefreshToken_UpdateTokenError(t *testing.T) {
	// Arrange
	mockDB, sqlm, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	mockLog := new(mocks.LoggerMock)
	mockLog.On("Info", "RefreshManager created").Return()
	mockLog.On("Info", "Refresh-Refresh: Transaction is begining successfully").Return()
	mockLog.On("Infof", "User's %s refresh token was found", mock.Anything).Return()
	mockLog.On("Errorf", "Error updating token: %v", mock.Anything).Return()

	userId := "testUserId12"
	newTestRefreshToken := "newTestToken12"

	sqlm.ExpectBegin()
	sqlm.ExpectQuery(`SELECT token FROM users_tokens WHERE user_id = \$1 FOR UPDATE`).
		WithArgs(userId).
		WillReturnRows(sqlmock.NewRows([]string{"token"}).AddRow("oldToken"))
	sqlm.ExpectExec(`UPDATE users_tokens SET token = \$1 WHERE user_id = \$2`).
		WithArgs(newTestRefreshToken, userId).
		WillReturnError(errors.New("update token error"))
	sqlm.ExpectRollback()

	// Act
	refreshManager := NewRefreshManager(mockDB, mockLog)
	result, err := refreshManager.RefreshToken(userId, newTestRefreshToken)

	// Assert
	assert.False(t, result, "Expected false when starting transaction fails")
	assert.Error(t, err, "Expected error when starting transaction fails")
}

func TestRefreshToken_CommitTransactionError(t *testing.T) {
	// Arrange
	mockDB, sqlm, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	mockLog := new(mocks.LoggerMock)
	mockLog.On("Info", "RefreshManager created").Return()
	mockLog.On("Info", "Refresh-Refresh: Transaction is begining successfully").Return()
	mockLog.On("Infof", "User's %s refresh token was found", mock.Anything).Return()
	mockLog.On("Infof", "User's %s refresh token updated successfully", mock.Anything).Return()
	mockLog.On("Errorf", "Error committing transaction: %v", mock.Anything).Return()

	userId := "testUserId13"
	newTestRefreshToken := "newTestToken13"

	sqlm.ExpectBegin()
	sqlm.ExpectQuery(`SELECT token FROM users_tokens WHERE user_id = \$1 FOR UPDATE`).
		WithArgs(userId).
		WillReturnRows(sqlmock.NewRows([]string{"token"}).AddRow("oldToken"))
	sqlm.ExpectExec(`UPDATE users_tokens SET token = \$1 WHERE user_id = \$2`).
		WithArgs(newTestRefreshToken, userId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlm.ExpectCommit().WillReturnError(errors.New("commit transaction error"))

	// Act
	refreshManager := NewRefreshManager(mockDB, mockLog)
	result, err := refreshManager.RefreshToken(userId, newTestRefreshToken)

	// Assert
	assert.False(t, result, "Expected false when starting transaction fails")
	assert.Error(t, err, "Expected error when starting transaction fails")
}
