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
	role_serviceName = "IRoleManager"
)

func insertUser(t *testing.T, db *sql.DB, userName, hashedPassword string) string {
	var userId string
	query := `INSERT INTO users (username, password_hash) VALUES ($1, $2) RETURNING id`
	err := db.QueryRow(query, userName, hashedPassword).Scan(&userId)
	assert.NoError(t, err, "Expected no error when gets querying user")
	return userId
}

func getUserRoleId(t *testing.T, db *sql.DB, testRole string) string {
	var roleId string
	query := `SELECT id FROM roles WHERE role_name = $1`
	err := db.QueryRow(query, testRole).Scan(&roleId)
	assert.NoError(t, err, "Expected no error when gets querying role")
	return roleId
}

func insertRoleToUser(t *testing.T, db *sql.DB, userId, roleId string) {
	query := `INSERT INTO users_roles (user_id, role_id) VALUES ($1, $2) ON CONFLICT DO NOTHING`
	_, err := db.Exec(query, userId, roleId)
	assert.NoError(t, err, "Expected no error when gets querying user role")
}

func TestNewRoleManager_Success(t *testing.T) {
	// Arrange
	db := dbs[role_serviceName]
	mockLog := new(mocks.LoggerMock)
	mockLog.On("Infof", "Role %s inserts successful", mock.Anything).Twice().Return()
	mockLog.On("Info", "RoleManager created").Return()

	// Act
	roleManager := NewRoleManager(db, mockLog)

	// Assert
	assert.NotNil(t, roleManager, "Expected not nil when create new role manager")

	rows, err := db.Query(`SELECT role_name FROM roles`)
	assert.NoError(t, err, "Expected no error when querying roles")
	defer rows.Close()

	var roleNames []string
	for rows.Next() {
		var role string
		err = rows.Scan(&role)
		assert.NoError(t, err, "Expected no error when scanning role_name")
		roleNames = append(roleNames, role)
	}
	assert.Equal(t, 2, len(roleNames), "Expected equal between expected count and actual count of roles from database")
}

func TestNewRoleManager_Error(t *testing.T) {
	// Arrange
	mockDB, sqlm, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	mockLog := new(mocks.LoggerMock)
	mockLog.On("Errorf", "Error during role %s insertion: %v", mock.Anything, mock.Anything).Twice().Return()
	mockLog.On("Info", "RoleManager created").Return()

	testRole := "testRole"

	sqlm.ExpectExec(`INSERT INTO roles (role_name) VALUES ($1) ON CONFLICT (role_name) DO NOTHING`).
		WithArgs(testRole).
		WillReturnError(errors.New("create new role error"))

	// Act
	roleManager := NewRoleManager(mockDB, mockLog)

	// Assert
	assert.Nil(t, roleManager, "Expected nil when create new role manager")
}

func TestCheckRoleExists_Success(t *testing.T) {
	// Arrange
	db := dbs[role_serviceName]
	mockLog := new(mocks.LoggerMock)
	mockLog.On("Infof", "Role %s inserts successful", mock.Anything).Twice().Return()
	mockLog.On("Info", "RoleManager created").Return()
	mockLog.On("Infof", "Role %s exists", mock.Anything).Return()

	testRole := "user"

	// Act
	rm := NewRoleManager(db, mockLog)
	assert.NotNil(t, rm, "Expected not nil when create new role manager")

	roleManager := &roleManger{
		db:  db,
		log: mockLog,
	}
	roleId, result := roleManager.checkRoleExists(testRole)

	// Assert
	assert.NotEqual(t, "", roleId, "Expected no roleId when querying role")
	assert.True(t, result, "Expected true when querying roles")
}

func TestCheckRoleExists_NoRoleExists(t *testing.T) {
	// Arrange
	db := dbs[role_serviceName]
	mockLog := new(mocks.LoggerMock)
	mockLog.On("Warnf", "Not find role %s in table roles", mock.Anything).Return()

	testRole := "testRole"

	// Act
	roleManager := &roleManger{
		db:  db,
		log: mockLog,
	}
	roleId, result := roleManager.checkRoleExists(testRole)

	// Assert
	assert.Equal(t, "", roleId, "Expected no roleId when querying role")
	assert.False(t, result, "Expected false when querying roles")
}

func TestCheckRoleExists_AnotherError(t *testing.T) {
	// Arrange
	mockDB, sqlm, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	mockLog := new(mocks.LoggerMock)
	mockLog.On("Warnf", "Error during find role %s: %v", mock.Anything, mock.Anything).Return()

	testRole := "testRole"

	sqlm.ExpectQuery(`SELECT id FROM roles WHERE role_name = $1`).
		WithArgs(testRole).
		WillReturnError(errors.New("return roleId error"))

	// Act
	roleManager := &roleManger{
		db:  mockDB,
		log: mockLog,
	}
	roleId, result := roleManager.checkRoleExists(testRole)

	// Assert
	assert.Equal(t, "", roleId, "Expected no roleId when querying role")
	assert.False(t, result, "Expected false when querying roles")
}

func TestAddRoleToUser_Success(t *testing.T) {
	// Arrange
	db := dbs[role_serviceName]
	mockLog := new(mocks.LoggerMock)
	mockLog.On("Infof", "Role %s inserts successful", mock.Anything).Twice().Return()
	mockLog.On("Info", "RoleManager created").Return()
	mockLog.On("Infof", "Role %s exists", mock.Anything).Return()
	mockLog.On("Infof", "Role %s added successfully to user %s", mock.Anything, mock.Anything).Return()

	testUserName := "testUser1"
	testHashedPassword := "testPassword1"
	testRole := "user"

	userId := insertUser(t, db, testUserName, testHashedPassword)

	// Act
	rm := NewRoleManager(db, mockLog)
	assert.NotNil(t, rm, "Expected not nil when create new role manager")

	roleManager := &roleManger{
		db:  db,
		log: mockLog,
	}
	err := roleManager.addRoleToUser(userId, testRole)

	// Assert
	assert.NoError(t, err, "Expected no error when adding role")
}

func TestAddRoleToUser_NoRoleExist(t *testing.T) {
	// Arrange
	db := dbs[role_serviceName]
	mockLog := new(mocks.LoggerMock)
	mockLog.On("Infof", "Role %s inserts successful", mock.Anything).Twice().Return()
	mockLog.On("Info", "RoleManager created").Return()
	mockLog.On("Warnf", "Not find role %s in table roles", mock.Anything).Return()

	testUserId := "testUserId"
	testRole := "testRole"

	// Act
	roleManager := &roleManger{
		db:  db,
		log: mockLog,
	}
	err := roleManager.addRoleToUser(testUserId, testRole)

	// Assert
	assert.Error(t, err, "Expected an error when adding role")
}

func TestAddRoleToUser_NoUserIdExist(t *testing.T) {
	// Arrange
	db := dbs[role_serviceName]
	mockLog := new(mocks.LoggerMock)
	mockLog.On("Infof", "Role %s inserts successful", mock.Anything).Twice().Return()
	mockLog.On("Info", "RoleManager created").Return()
	mockLog.On("Infof", "Role %s exists", mock.Anything).Return()
	mockLog.On("Errorf", "Can't inserts role %s to user %s: %v", mock.Anything, mock.Anything, mock.Anything).Return()

	testUserId := "testUserId"
	testRole := "user"

	// Act
	roleManager := &roleManger{
		db:  db,
		log: mockLog,
	}
	rm := NewRoleManager(db, mockLog)
	assert.NotNil(t, rm, "Expected not nil when create new role manager")
	err := roleManager.addRoleToUser(testUserId, testRole)

	// Assert
	assert.Error(t, err, "Expected an error when adding role")
}

func TestAddRolesToUser_Success_OneRole(t *testing.T) {
	// Arrange
	db := dbs[role_serviceName]
	mockLog := new(mocks.LoggerMock)
	mockLog.On("Infof", "Role %s inserts successful", mock.Anything).Twice().Return()
	mockLog.On("Info", "RoleManager created").Return()

	mockLog.On("Info", "Role-Add: Transaction is begining").Return()
	mockLog.On("Info", "Role-Add: Transaction is begined successfully").Return()

	mockLog.On("Infof", "Role %s exists", mock.Anything).Return()
	mockLog.On("Infof", "Role %s added successfully to user %s", mock.Anything, mock.Anything).Return()

	mockLog.On("Infof", "All roles %v added to user %s successfully", mock.Anything, mock.Anything).Return()
	mockLog.On("Info", "Role-Add: Transaction is commited successfully").Return()

	testUserName := "testUser2"
	testHashedPassword := "testPassword2"

	userId := insertUser(t, db, testUserName, testHashedPassword)

	// Act
	roleManager := NewRoleManager(db, mockLog)
	assert.NotNil(t, roleManager, "Expected not nil when create new role manager")
	err := roleManager.AddRolesToUser(userId, "user")

	// Assert
	assert.NoError(t, err, "Expected no error when adding role")
}

func TestAddRolesToUser_Success_TwoRoles(t *testing.T) {
	// Arrange
	db := dbs[role_serviceName]
	mockLog := new(mocks.LoggerMock)
	mockLog.On("Infof", "Role %s inserts successful", mock.Anything).Twice().Return()
	mockLog.On("Info", "RoleManager created").Return()

	mockLog.On("Info", "Role-Add: Transaction is begining").Return()
	mockLog.On("Info", "Role-Add: Transaction is begined successfully").Return()

	mockLog.On("Infof", "Role %s exists", mock.Anything).Twice().Return()
	mockLog.On("Infof", "Role %s added successfully to user %s", mock.Anything, mock.Anything).Twice().Return()

	mockLog.On("Infof", "All roles %v added to user %s successfully", mock.Anything, mock.Anything).Return()
	mockLog.On("Info", "Role-Add: Transaction is commited successfully").Return()

	testUserName := "testUser3"
	testHashedPassword := "testPassword3"

	userId := insertUser(t, db, testUserName, testHashedPassword)

	// Act
	roleManager := NewRoleManager(db, mockLog)
	assert.NotNil(t, roleManager, "Expected not nil when create new role manager")
	err := roleManager.AddRolesToUser(userId, "user", "admin")

	// Assert
	assert.NoError(t, err, "Expected no error when adding roles")
}

func TestAddRolesToUser_NoUserIdExist(t *testing.T) {
	// Arrange
	db := dbs[role_serviceName]
	mockLog := new(mocks.LoggerMock)
	mockLog.On("Infof", "Role %s inserts successful", mock.Anything).Twice().Return()
	mockLog.On("Info", "RoleManager created").Return()

	mockLog.On("Info", "Role-Add: Transaction is begining").Return()
	mockLog.On("Info", "Role-Add: Transaction is begined successfully").Return()

	mockLog.On("Infof", "Role %s exists", mock.Anything).Return()
	mockLog.On("Errorf", "Can't inserts role %s to user %s: %v", mock.Anything, mock.Anything, mock.Anything).Return()

	testUserId := "testUserId"

	// Act
	roleManager := NewRoleManager(db, mockLog)
	assert.NotNil(t, roleManager, "Expected not nil when create new role manager")
	err := roleManager.AddRolesToUser(testUserId, "user")

	// Assert
	assert.Error(t, err, "Expected an error when adding roles")
}

func TestAddRolesToUser_NoSecondRoleExist(t *testing.T) {
	// Arrange
	db := dbs[role_serviceName]
	mockLog := new(mocks.LoggerMock)
	mockLog.On("Infof", "Role %s inserts successful", mock.Anything).Twice().Return()
	mockLog.On("Info", "RoleManager created").Return()

	mockLog.On("Info", "Role-Add: Transaction is begining").Return()
	mockLog.On("Info", "Role-Add: Transaction is begined successfully").Return()

	mockLog.On("Infof", "Role %s exists", mock.Anything).Return()
	mockLog.On("Infof", "Role %s added successfully to user %s", mock.Anything, mock.Anything).Return()

	mockLog.On("Infof", "Role %s exists", mock.Anything).Return()
	mockLog.On("Errorf", "Can't inserts role %s to user %s: %v", mock.Anything, mock.Anything, mock.Anything).Return()

	testUserId := "testUserId"
	testRole := "testRole"

	// Act
	roleManager := NewRoleManager(db, mockLog)
	assert.NotNil(t, roleManager, "Expected not nil when create new role manager")
	err := roleManager.AddRolesToUser(testUserId, "user", testRole)

	// Assert
	assert.Error(t, err, "Expected an error when adding roles")
}

func TestAddRolesToUser_BeginTransactionError(t *testing.T) {
	// Arrange
	mockDB, sqlm, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	mockLog := new(mocks.LoggerMock)
	mockLog.On("Info", "Role-Add: Transaction is begining").Return()
	mockLog.On("Errorf", "Can't starts transaction: %v", mock.Anything).Return()

	testUserId := "testUserId"
	testRole := "testRole"

	sqlm.ExpectBegin().WillReturnError(errors.New("begin transaction error"))

	// Act
	roleManager := &roleManger{
		db:  mockDB,
		log: mockLog,
	}
	err = roleManager.AddRolesToUser(testUserId, testRole)

	// Assert
	assert.Error(t, err, "Expected an error when begining transaction")
}

func TestAddRolesToUser_CommitTransactionError(t *testing.T) {
	// Arrange
	mockDB, sqlm, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	mockLog := new(mocks.LoggerMock)
	mockLog.On("Info", "Role-Add: Transaction is begining").Return()
	mockLog.On("Info", "Role-Add: Transaction is begined successfully").Return()

	mockLog.On("Infof", "All roles %v added to user %s successfully", mock.Anything, mock.Anything).Return()
	mockLog.On("Errorf", "Can't commits transaction: %v", mock.Anything).Return()

	sqlm.ExpectBegin()
	sqlm.ExpectCommit().WillReturnError(errors.New("begin transaction error"))

	testUserId := "testUserId"

	// Act
	roleManager := &roleManger{
		db:  mockDB,
		log: mockLog,
	}
	err = roleManager.AddRolesToUser(testUserId, []string{}...)

	// Assert
	assert.Error(t, err, "Expected an error when commiting transaction")
}

func TestDeleteRoleFromUser_Success(t *testing.T) {
	// Arrange
	db := dbs[role_serviceName]
	mockLog := new(mocks.LoggerMock)
	mockLog.On("Infof", "Role %s inserts successful", mock.Anything).Twice().Return()
	mockLog.On("Info", "RoleManager created").Return()
	mockLog.On("Infof", "Role %s exists", mock.Anything).Return()
	mockLog.On("Infof", "Role %s deleted successfully from user %s", mock.Anything, mock.Anything).Return()

	testUserName := "testUser4"
	testHashedPassword := "testPassword4"
	testRole := "user"

	rm := NewRoleManager(db, mockLog)
	assert.NotNil(t, rm, "Expected not nil when create new role manager")

	userId := insertUser(t, db, testUserName, testHashedPassword)
	roleId := getUserRoleId(t, db, testRole)
	insertRoleToUser(t, db, userId, roleId)

	// Act
	roleManager := &roleManger{
		db:  db,
		log: mockLog,
	}
	err := roleManager.deleteRoleFromUser(userId, testRole)

	// Assert
	assert.NoError(t, err, "Expected no error when deleting role")
}

func TestDeleteRoleFromUser_NoRoleExist(t *testing.T) {
	// Arrange
	db := dbs[role_serviceName]
	mockLog := new(mocks.LoggerMock)
	mockLog.On("Infof", "Role %s inserts successful", mock.Anything).Twice().Return()
	mockLog.On("Info", "RoleManager created").Return()
	mockLog.On("Warnf", "Not find role %s in table roles", mock.Anything).Return()

	testUserId := "testUserId"
	testRole := "testRole"

	// Act
	roleManager := &roleManger{
		db:  db,
		log: mockLog,
	}
	err := roleManager.deleteRoleFromUser(testUserId, testRole)

	// Assert
	assert.Error(t, err, "Expected an error when deleting role")
}

func TestDeleteRoleFromUser_NoRecordExist(t *testing.T) {
	// Arrange
	db := dbs[role_serviceName]
	mockLog := new(mocks.LoggerMock)
	mockLog.On("Infof", "Role %s inserts successful", mock.Anything).Twice().Return()
	mockLog.On("Info", "RoleManager created").Return()
	mockLog.On("Infof", "Role %s exists", mock.Anything).Return()
	mockLog.On("Errorf", "Can't deletes role %s from user %s: %v", mock.Anything, mock.Anything, mock.Anything).Return()

	testUserId := "testUserId"
	testRole := "user"

	// Act
	roleManager := &roleManger{
		db:  db,
		log: mockLog,
	}
	rm := NewRoleManager(db, mockLog)
	assert.NotNil(t, rm, "Expected not nil when create new role manager")
	err := roleManager.deleteRoleFromUser(testUserId, testRole)

	// Assert
	assert.Error(t, err, "Expected an error when deleting role")
}

func TestDeleteRolesFromUser_Success_OneRole(t *testing.T) {
	// Arrange
	db := dbs[role_serviceName]
	mockLog := new(mocks.LoggerMock)
	mockLog.On("Infof", "Role %s inserts successful", mock.Anything).Twice().Return()
	mockLog.On("Info", "RoleManager created").Return()

	mockLog.On("Info", "Role-Delete: Transaction is begining").Return()
	mockLog.On("Info", "Role-Delete: Transaction is begined successfully").Return()

	mockLog.On("Infof", "Role %s exists", mock.Anything).Return()
	mockLog.On("Infof", "Role %s deleted successfully from user %s", mock.Anything, mock.Anything).Return()

	mockLog.On("Infof", "All roles %v deleted from user %s successfully", mock.Anything, mock.Anything).Return()
	mockLog.On("Info", "Role-Delete: Transaction is commited successfully").Return()

	testUserName := "testUser5"
	testHashedPassword := "testPassword5"
	testRole := "user"

	roleManager := NewRoleManager(db, mockLog)
	assert.NotNil(t, roleManager, "Expected not nil when create new role manager")

	userId := insertUser(t, db, testUserName, testHashedPassword)
	roleId := getUserRoleId(t, db, testRole)
	insertRoleToUser(t, db, userId, roleId)

	// Act
	err := roleManager.DeleteRolesFromUser(userId, testRole)

	// Assert
	assert.NoError(t, err, "Expected no error when deleting role")
}

func TestDeleteRolesFromUser_Success_TwoRoles(t *testing.T) {
	// Arrange
	db := dbs[role_serviceName]
	mockLog := new(mocks.LoggerMock)
	mockLog.On("Infof", "Role %s inserts successful", mock.Anything).Twice().Return()
	mockLog.On("Info", "RoleManager created").Return()

	mockLog.On("Info", "Role-Delete: Transaction is begining").Return()
	mockLog.On("Info", "Role-Delete: Transaction is begined successfully").Return()

	mockLog.On("Infof", "Role %s exists", mock.Anything).Return()
	mockLog.On("Infof", "Role %s deleted successfully from user %s", mock.Anything, mock.Anything).Return()

	mockLog.On("Infof", "All roles %v deleted from user %s successfully", mock.Anything, mock.Anything).Return()
	mockLog.On("Info", "Role-Delete: Transaction is commited successfully").Return()

	testUserName := "testUser6"
	testHashedPassword := "testPassword6"
	testRole1 := "user"
	testRole2 := "admin"

	roleManager := NewRoleManager(db, mockLog)
	assert.NotNil(t, roleManager, "Expected not nil when create new role manager")

	userId := insertUser(t, db, testUserName, testHashedPassword)
	roleId1 := getUserRoleId(t, db, testRole1)
	roleId2 := getUserRoleId(t, db, testRole2)
	insertRoleToUser(t, db, userId, roleId1)
	insertRoleToUser(t, db, userId, roleId2)

	// Act
	err := roleManager.DeleteRolesFromUser(userId, testRole1, testRole2)

	// Assert
	assert.NoError(t, err, "Expected no error when deleting role")
}

func TestDeleteRolesFromUser_NoUserIdExist(t *testing.T) {
	// Arrange
	db := dbs[role_serviceName]
	mockLog := new(mocks.LoggerMock)
	mockLog.On("Infof", "Role %s inserts successful", mock.Anything).Twice().Return()
	mockLog.On("Info", "RoleManager created").Return()

	mockLog.On("Info", "Role-Delete: Transaction is begining").Return()
	mockLog.On("Info", "Role-Delete: Transaction is begined successfully").Return()

	mockLog.On("Infof", "Role %s exists", mock.Anything).Return()
	mockLog.On("Errorf", "Can't deletes role %s from user %s: %v", mock.Anything, mock.Anything, mock.Anything).Return()
	mockLog.On("Errorf", "Can't delete role %s from user %s", mock.Anything, mock.Anything).Return()

	testUserId := "testUserId"
	testRole := "user"

	// Act
	roleManager := NewRoleManager(db, mockLog)
	assert.NotNil(t, roleManager, "Expected not nil when create new role manager")
	err := roleManager.DeleteRolesFromUser(testUserId, testRole)

	// Assert
	assert.Error(t, err, "Expected an error when deleting roles")
}

func TestDeleteRolesFromUser_NoSecondRoleExist(t *testing.T) {
	// Arrange
	db := dbs[role_serviceName]
	mockLog := new(mocks.LoggerMock)
	mockLog.On("Infof", "Role %s inserts successful", mock.Anything).Twice().Return()
	mockLog.On("Info", "RoleManager created").Return()

	mockLog.On("Info", "Role-Delete: Transaction is begining").Return()
	mockLog.On("Info", "Role-Delete: Transaction is begined successfully").Return()

	mockLog.On("Infof", "Role %s exists", mock.Anything).Return()
	mockLog.On("Infof", "Role %s deleted successfully from user %s", mock.Anything, mock.Anything).Return()

	mockLog.On("Warnf", "Not find role %s in table roles", mock.Anything).Return()
	mockLog.On("Errorf", "Can't deletes role %s from user %s: %v", mock.Anything, mock.Anything, mock.Anything).Return()
	mockLog.On("Errorf", "Can't delete role %s from user %s", mock.Anything, mock.Anything).Return()

	testUserName := "testUser7"
	testHashedPassword := "testPassword7"
	testRole := "user"

	roleManager := NewRoleManager(db, mockLog)
	assert.NotNil(t, roleManager, "Expected not nil when create new role manager")

	userId := insertUser(t, db, testUserName, testHashedPassword)
	roleId := getUserRoleId(t, db, testRole)
	insertRoleToUser(t, db, userId, roleId)

	// Act
	err := roleManager.DeleteRolesFromUser(userId, testRole, "testUser")

	// Assert
	assert.Error(t, err, "Expected an error when deleting roles")
}

func TestDeleteRolesFromUser_BeginTransactionError(t *testing.T) {
	// Arrange
	mockDB, sqlm, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	mockLog := new(mocks.LoggerMock)
	mockLog.On("Info", "Role-Delete: Transaction is begining").Return()
	mockLog.On("Errorf", "Can't starts transaction: %v", mock.Anything).Return()

	testUserId := "testUserId"
	testRole := "testRole"

	sqlm.ExpectBegin().WillReturnError(errors.New("begin transaction error"))

	// Act
	roleManager := &roleManger{
		db:  mockDB,
		log: mockLog,
	}
	err = roleManager.DeleteRolesFromUser(testUserId, testRole)

	// Assert
	assert.Error(t, err, "Expected an error when begining transaction")
}

func TestDeleteRolesFromUser_CommitTransactionError(t *testing.T) {
	// Arrange
	mockDB, sqlm, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	mockLog := new(mocks.LoggerMock)
	mockLog.On("Info", "Role-Delete: Transaction is begining").Return()
	mockLog.On("Info", "Role-Delete: Transaction is begined successfully").Return()

	mockLog.On("Infof", "All roles %v deleted from user %s successfully", mock.Anything, mock.Anything).Return()
	mockLog.On("Errorf", "Can't commits transaction: %v", mock.Anything).Return()

	sqlm.ExpectBegin()
	sqlm.ExpectCommit().WillReturnError(errors.New("begin transaction error"))

	testUserId := "testUserId"

	// Act
	roleManager := &roleManger{
		db:  mockDB,
		log: mockLog,
	}
	err = roleManager.DeleteRolesFromUser(testUserId, []string{}...)

	// Assert
	assert.Error(t, err, "Expected an error when commiting transaction")
}
