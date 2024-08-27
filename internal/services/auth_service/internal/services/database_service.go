package services

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/peygy/nektoyou/internal/pkg/logger"
	"github.com/peygy/nektoyou/internal/services/auth_service/config"
	_ "github.com/lib/pq"
)

type DatabaseServer struct {
	DB *sql.DB
	log logger.ILogger
}

func NewDatabaseConnection(cfg *config.DatabaseConfig, log logger.ILogger) *DatabaseServer {
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
        return nil
    }

	log.Info("Database connection structure is created")
	return &DatabaseServer{DB: db, log: log}
}

func (dc *DatabaseServer) Run(ctx context.Context) error {
	go func () {
		for {
			select {
			case <-ctx.Done():
				dc.log.Info("Shutting down Database")
				dc.DB.Close()
				dc.log.Info("Database exited properly")
				return
			}
		}
	} ()

    createTablesQuery := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        username VARCHAR(50) UNIQUE NOT NULL,
        password_hash VARCHAR(255) NOT NULL,
        created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
    );

    CREATE TABLE IF NOT EXISTS roles (
        id SERIAL PRIMARY KEY,
        role_name VARCHAR(50) UNIQUE NOT NULL,
        created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
    );

    CREATE TABLE IF NOT EXISTS users_roles (
        user_id INT REFERENCES users(id) ON DELETE CASCADE,
        role_id INT REFERENCES roles(id) ON DELETE CASCADE,
        PRIMARY KEY (user_id, role_id)
    );

    CREATE TABLE IF NOT EXISTS users_tokens (
        id SERIAL PRIMARY KEY,
        user_id INT REFERENCES users(id) ON DELETE CASCADE,
        token VARCHAR(255) UNIQUE NOT NULL,
        created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
        expires_at TIMESTAMP WITH TIME ZONE NOT NULL
    );`

    _, err := dc.DB.Exec(createTablesQuery)
    if err != nil {
		dc.log.Fatal("Error during creation of table: " + err.Error())
        return err
    }
    dc.log.Info("Tables users, roles, users_roles, users_tokens created successful")

	return nil
}