package data

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/peygy/nektoyou/internal/pkg/logger"
	"github.com/peygy/nektoyou/internal/services/auth_service/config"
)

const schemaFilePath = "./config/schema.sql"

type IDatabaseServer interface {
	Run(ctx context.Context) error
}

type databaseServer struct {
	db  *sql.DB
	log logger.ILogger
}

func NewDatabaseConnection(cfg *config.DatabaseConfig, log logger.ILogger) (IDatabaseServer, *sql.DB) {
	psqlconn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.DbName,
	)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal("Error during connection " + psqlconn + " to the database: " + err.Error())
		return nil, nil
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Error during ping the database: " + err.Error())
		return nil, nil
	}

	log.Info("Database connection structure is created")
	return &databaseServer{db: db, log: log}, db
}

func InitDatabaseSchema(db *sql.DB, log logger.ILogger) error {
	sqlBytes, err := os.ReadFile(schemaFilePath)
	if err != nil {
		log.Fatal("Error reading SQL file: " + err.Error())
		return err
	}
	sqlContent := string(sqlBytes)

	_, err = db.Exec(sqlContent)
	if err != nil {
		log.Fatal("Error during creation of table: " + err.Error())
		return err
	}

	log.Info("Tables users, roles, users_roles, users_tokens created successful")
	return nil
}

func (dc *databaseServer) Run(ctx context.Context) error {
	go func() {
		for {
			select {
			case <-ctx.Done():
				dc.log.Info("Shutting down Database")
				dc.db.Close()
				dc.log.Info("Database exited properly")
				return
			}
		}
	}()

	return nil
}
