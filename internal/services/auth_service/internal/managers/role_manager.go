package managers

import (
	"database/sql"

	"github.com/peygy/nektoyou/internal/pkg/logger"
)

type IRoleManager interface {
	AddRolesToUser(userId string, roles ...string) bool

	DeleteRolesFromUser(userId string, roles ...string) bool
}

var roles = []string{
	"user",
	"admin",
}

type roleManger struct {
	db  *sql.DB
	log logger.ILogger
}

func NewRoleManager(db *sql.DB, log logger.ILogger) IRoleManager {
	for _, role := range roles {
		query := `INSERT INTO roles (role_name) VALUES ($1) ON CONFLICT (role_name) DO NOTHING`

		_, err := db.Exec(query, role)
		if err != nil {
			log.Error("Error during role " + role + " insertion: " + err.Error())
			return nil
		}
		log.Info("Role " + role + " inserts successful")
	}

	return &roleManger{db: db, log: log}
}

func (rm *roleManger) addRoleToUser(userId, role string) bool {
	roleId, ok := rm.checkRoleExists(role)
	if !ok {
		return false
	}

	query := `INSERT INTO users_roles (user_id, role_id) VALUES ($1, $2) ON CONFLICT DO NOTHING`
	_, err := rm.db.Exec(query, userId, roleId)
	if err != nil {
		rm.log.Warn("Error during inserts role " + role + " to user " + userId + ": " + err.Error())
		return false
	}
	return true
}

func (rm *roleManger) AddRolesToUser(userId string, roles ...string) bool {
	tx, err := rm.db.Begin()
	if err != nil {
		rm.log.Error(err.Error())
		return false
	}
	defer tx.Rollback()

	for _, role := range roles {
		added := rm.addRoleToUser(userId, role)
		if !added {
			return false
		}
	}

	if err = tx.Commit(); err != nil {
		rm.log.Error(err.Error())
		return false
	}
	return true
}

func (rm *roleManger) deleteRoleFromUser(userId, role string) bool {
	roleId, ok := rm.checkRoleExists(role)
	if !ok {
		return false
	}

	query := `DELETE FROM users_roles WHERE user_id = $1 AND role_id = $2`
	result, err := rm.db.Exec(query, userId, roleId)
	if err != nil {
		rm.log.Warn("Error during deletes role " + role + " from user " + userId + ": " + err.Error())
		return false
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		rm.log.Warn("Error during takes result from deleting " + role + " from user " + userId + ": " + err.Error())
		return false
	}

	return rowsAffected > 0
}

func (rm *roleManger) DeleteRolesFromUser(userId string, roles ...string) bool {
	tx, err := rm.db.Begin()
	if err != nil {
		rm.log.Error("Error starting transaction: " + err.Error())
		return false
	}
	defer tx.Rollback()

	for _, role := range roles {
		deleted := rm.deleteRoleFromUser(userId, role)
		if !deleted {
			rm.log.Error("role " + role + " not be deleted")
			return false
		}
	}

	if err = tx.Commit(); err != nil {
		rm.log.Error("Error committing transaction: " + err.Error())
		return false
	}
	return true
}

func (rm *roleManger) checkRoleExists(role string) (string, bool) {
	var roleId string
	query := `SELECT id FROM roles WHERE role_name = $1`

	if err := rm.db.QueryRow(query).Scan(&roleId); err != nil {
		if err == sql.ErrNoRows {
			rm.log.Warn("Not find role " + role + " in table roles")
		} else {
			rm.log.Warn("Error during find role " + role + ": " + err.Error())
		}

		return "", false
	}

	return roleId, true
}
