package managers

import (
	"database/sql"
	"errors"

	"github.com/peygy/nektoyou/internal/pkg/logger"
)

type IRefreshManager interface {
	AddToken(userId, refreshToken string) error
	GetToken(userId string) (string, error)
	RefreshToken(userId, newRefreshToken string) (bool, error)
}

type refreshManager struct {
	db  *sql.DB
	log logger.ILogger
}

func NewRefreshManager(db *sql.DB, log logger.ILogger) IRefreshManager {
	log.Info("RefreshManager created")
	return &refreshManager{db: db, log: log}
}

func (rm *refreshManager) AddToken(userId, refreshToken string) error {
	query := `INSERT INTO users_tokens (user_id, token) VALUES ($1, $2) ON CONFLICT (token) DO NOTHING`

	_, err := rm.db.Exec(query, userId, refreshToken)
	if err != nil {
		rm.log.Errorf("Can't inserts refresh token with error: %v", err)
		return errors.New("managers-refresh: can't inserts refresh token with error")
	}

	rm.log.Infof("Refresh token %s successfully added to user %s", refreshToken, userId)
	return nil
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

	rm.log.Info("Token was founded")

	return token, nil
}

func (rm *refreshManager) RefreshToken(userId, newRefreshToken string) (bool, error) {
	tx, err := rm.db.Begin()
	if err != nil {
		rm.log.Errorf("Error starting transaction: %v", err)
		return false, errors.New("managers-refresh: can't starting transaction")
	}
	rm.log.Info("Refresh-Refresh: Transaction is begining successfully")
	defer tx.Rollback()

	var oldToken string
	query := `SELECT token FROM users_tokens WHERE user_id = $1 FOR UPDATE`
	err = tx.QueryRow(query, userId).Scan(&oldToken)
	if err != nil {
		if err == sql.ErrNoRows {
			rm.log.Errorf("No any refresh token in database: %v", err)
			return false, errors.New("managers-refresh: no any refresh token in database")
		}

		rm.log.Errorf("Error retrieving old token: %v", err)
		return false, errors.New("managers-refresh: can't retrieving old token")
	}
	rm.log.Infof("User's %s refresh token was found", userId)

	updateQuery := `UPDATE users_tokens SET token = $1 WHERE user_id = $2`
	_, err = tx.Exec(updateQuery, newRefreshToken, userId)
	if err != nil {
		rm.log.Errorf("Error updating token: %v", err)
		return false, errors.New("managers-refresh: can't updating token")
	}
	rm.log.Infof("User's %s refresh token updated successfully", userId)

	if err = tx.Commit(); err != nil {
		rm.log.Errorf("Error committing transaction: %v", err)
		return false, errors.New("managers-refresh: can't committing transaction")
	}

	rm.log.Info("Refresh-Refresh: Transaction is commited successfully")
	return true, nil
}
