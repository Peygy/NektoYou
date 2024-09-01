package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/peygy/nektoyou/internal/pkg/logger"
)

type IDatabaseServer interface {
	Run(ctx context.Context) error
}

type databaseServer struct {
	db  *sql.DB
	log logger.ILogger
}

type PostgresConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbname"`
}

func NewDatabaseConnection(cfg *PostgresConfig, log logger.ILogger) (IDatabaseServer, *sql.DB) {
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
		log.Fatalf("Error during connection %s to the database: %v", psqlconn, err)
		return nil, nil
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Error during ping the database: %v", err)
		return nil, nil
	}

	log.Info("Database connection structure is created")
	return &databaseServer{db: db, log: log}, db
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
