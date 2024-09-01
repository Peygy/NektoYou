package data

import (
	"database/sql"
	"errors"
	"os"

	_ "github.com/lib/pq"
	"github.com/peygy/nektoyou/internal/pkg/logger"
)

const schemaFilePath = "./config/schema.sql"

func InitDatabaseSchema(db *sql.DB, log logger.ILogger) error {
	sqlBytes, err := os.ReadFile(schemaFilePath)
	if err != nil {
		log.Fatalf("Error reading SQL file: %v", err)
		return errors.New("data: can't reads SQL config file")
	}
	sqlContent := string(sqlBytes)

	_, err = db.Exec(sqlContent)
	if err != nil {
		log.Fatalf("Error during creation of table: %v", err)
		return errors.New("data: can't creates tables in the database")
	}

	log.Info("Tables users, roles, users_roles created successfully")
	return nil
}
