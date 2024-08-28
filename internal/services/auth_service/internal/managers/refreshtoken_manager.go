package managers

import (
	"database/sql"

	"github.com/peygy/nektoyou/internal/pkg/logger"
)

type IRefreshManager interface {
	AddToken(userId, refreshToken string) bool
	GetToken(userId string) string
	RefreshToken(userId, newRefreshToken string) (string, bool)
}

type refreshManager struct {
	db  *sql.DB
	log logger.ILogger
}

func NewRefreshManager(db *sql.DB, log logger.ILogger) IRefreshManager {
	return &refreshManager{db: db, log: log}
}

func (rm *refreshManager) AddToken(userId, refreshToken string) bool {
	query := `INSERT INTO users_tokens (user_id, token) VALUES ($1, $2) ON CONFLICT (token) DO NOTHING`

	_, err := rm.db.Exec(query, userId, refreshToken)
	if err != nil {
		rm.log.Error("Error inserting token: " + err.Error())
		return false
	}

	return true
}

func (rm *refreshManager) GetToken(userId string) string {
	var token string
	query := `SELECT token FROM users_tokens WHERE user_id = $1`

	err := rm.db.QueryRow(query, userId).Scan(&token)
	if err != nil {
		if err == sql.ErrNoRows {
			rm.log.Error("No token was found: " + err.Error())
			return ""
		}

		rm.log.Error("Error getting token: " + err.Error())
		return ""
	}

	rm.log.Error("Token is empty")
	return token
}

func (rm *refreshManager) RefreshToken(userId, newRefreshToken string) (string, bool) {
	tx, err := rm.db.Begin()
	if err != nil {
		rm.log.Error("Error starting transaction: " + err.Error())
		return "", false
	}
	defer tx.Rollback()

	var oldToken string
	query := `SELECT token FROM users_tokens WHERE user_id = $1 FOR UPDATE`

	err = tx.QueryRow(query, userId).Scan(&oldToken)
	if err != nil {
		if err == sql.ErrNoRows {
			rm.log.Error("No any token in database: " + err.Error())
			return "", false
		}

		rm.log.Error("Error retrieving old token: " + err.Error())
		return "", false
	}

	updateQuery := `UPDATE users_tokens SET token = $1 WHERE user_id = $2`
	_, err = tx.Exec(updateQuery, newRefreshToken, userId)
	if err != nil {
		rm.log.Error("Error updating token: " + err.Error())
		return "", false
	}

	err = tx.Commit()
	if err != nil {
		rm.log.Error("Error committing transaction: " + err.Error())
		return "", false
	}

	return newRefreshToken, true
}
