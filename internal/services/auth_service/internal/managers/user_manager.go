package managers

import (
	"database/sql"
	"errors"

	"github.com/peygy/nektoyou/internal/pkg/logger"
)

// Provides the service for managing user persistence store.
type IUserManager interface {
	// Adds new user to the database.
	// Returns new user id or error.
	AddUser(user UserRecord) (string, error)
	// Gets user from the database by id.
	// Returns user model or error.
	GetUserById(userId string) (UserRecord, error)
	// Updates user from the database by id.
	// Returns count of updated records or error.
	UpdateUserById(userId string, newUser UserRecord) (int, error)
	// Deletes user from the database by id.
	// Returns count of deleted records or error.
	DeleteUserById(userId string) (int, error)
}

type UserRecord struct {
	Id       string
	UserName string
	Password string
}

type userManger struct {
	iPasswordManager

	db  *sql.DB
	log logger.ILogger
}

func NewUserManager(db *sql.DB, log logger.ILogger) IUserManager {
	passwordManager := newPasswordManager(7, log)

	return &userManger{
		iPasswordManager: passwordManager,
		db:               db,
		log:              log,
	}
}

func (um *userManger) AddUser(user UserRecord) (string, error) {
	if err := um.validPassword(user.Password); err != nil {
		return "", err
	}

	hashedPassword, err := um.hashPassword(user.Password)
	if err != nil {
		return "", err
	}

	var userId string
	query := `INSERT INTO users (username, password_hash) VALUES ($1, $2) RETURNING id`
	if err := um.db.QueryRow(query, user.UserName, hashedPassword).Scan(&userId); err != nil {
		um.log.Errorf("Can't adds new user (%s, %s) to the database: %v", user.UserName, hashedPassword, err)
		return "", errors.New("managers-user: can't creates user in the database")
	}

	return userId, nil
}

func (um *userManger) GetUserById(userId string) (UserRecord, error) {
	f := new(UserRecord)
	return *f, nil
}

func (um *userManger) UpdateUserById(userId string, newUser UserRecord) (int, error) {
	return 0, nil
}

func (um *userManger) DeleteUserById(userId string) (int, error) {
	return 0, nil
}
