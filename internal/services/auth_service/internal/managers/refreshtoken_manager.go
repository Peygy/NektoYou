package managers

import (
	"database/sql"
	"errors"

	"github.com/peygy/nektoyou/internal/pkg/logger"
)

type IRefreshManager interface {
	AddToken(userId, refreshToken string) (bool, error)
	GetToken(userId string) (string, error)
	RefreshToken(userId, newRefreshToken string) (bool, error)
}

type refreshManager struct {
	db  *sql.DB
	log logger.ILogger
}

func NewRefreshManager(db *sql.DB, log logger.ILogger) IRefreshManager {
	return &refreshManager{db: db, log: log}
}

func (rm *refreshManager) AddToken(userId, refreshToken string) (bool, error) {
	query := `INSERT INTO users_tokens (user_id, token) VALUES ($1, $2) ON CONFLICT (token) DO NOTHING`

	_, err := rm.db.Exec(query, userId, refreshToken)
	if err != nil {
		rm.log.Errorf("Can't inserts refresh token with error: %v", err)
		return false, errors.New("managers-refresh: can't inserts refresh token with error")
	}

	return true, nil
}

func (rm *refreshManager) GetToken(userId string) (string, error) {
	var token string
	query := `SELECT token FROM users_tokens WHERE user_id = $1`

	err := rm.db.QueryRow(query, userId).Scan(&token)
	if err != nil {
		if err == sql.ErrNoRows {
			rm.log.Errorf("No refresh token was found: %v", err)
			return "", errors.New("managers-refresh: no refresh token was found")
		}

		rm.log.Errorf("Can't gets refresh token: %v", err)
		return "", errors.New("managers-refresh: can't gets refresh token")
	}

	if token == "" {
		rm.log.Info("Token is empty")
	} else {
		rm.log.Info("Token was founded")
	}

	return token, nil
}

func (rm *refreshManager) RefreshToken(userId, newRefreshToken string) (bool, error) {
	tx, err := rm.db.Begin()
	if err != nil {
		rm.log.Errorf("Error starting transaction: %v", err)
		return false, errors.New("managers-refresh: can't starting transaction")
	}
	defer tx.Rollback()

	var oldToken string
	query := `SELECT token FROM users_tokens WHERE user_id = $1 FOR UPDATE`

	err = tx.QueryRow(query, userId).Scan(&oldToken)
	if err != nil {
		if err == sql.ErrNoRows {
			rm.log.Errorf("No any refresh token in database: %v", err)
			return false, errors.New("managers-refresh: no any refresh token in database")
		}

		rm.log.Errorf("Error retrieving old token: ", err)
		return false, errors.New("managers-refresh: can't retrieving old token")
	}

	updateQuery := `UPDATE users_tokens SET token = $1 WHERE user_id = $2`
	_, err = tx.Exec(updateQuery, newRefreshToken, userId)
	if err != nil {
		rm.log.Errorf("Error updating token: ", err)
		return false, errors.New("managers-refresh: can't updating token")
	}

	err = tx.Commit()
	if err != nil {
		rm.log.Errorf("Error committing transaction: ", err)
		return false, errors.New("managers-refresh: can't committing transaction")
	}

	return true, nil
}
