package mocks

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const schemaFilePath = "../../config/schema.sql"

func SetupTestContainer_Postgres() (*sql.DB, func(), error) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"5432/tcp"},
		WaitingFor:   wait.ForListeningPort("5432/tcp"),
		Env: map[string]string{
			"POSTGRES_USER":     "testuser",
			"POSTGRES_PASSWORD": "testpassword",
			"POSTGRES_DB":       "testdb",
		},
	}

	postgresContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to start container: %w", err)
	}

	host, err := postgresContainer.Host(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get container host: %w", err)
	}

	port, err := postgresContainer.MappedPort(ctx, "5432")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get mapped port: %w", err)
	}

	dsn := fmt.Sprintf("postgres://testuser:testpassword@%s:%s/testdb?sslmode=disable", host, port.Port())
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Connect to db
	time.Sleep(time.Second * 5)

	// Create schema
	sqlBytes, err := os.ReadFile(schemaFilePath)
	if err != nil {
		return nil, nil, fmt.Errorf("failed reading SQL file: %w", err)
	}
	sqlContent := string(sqlBytes)

	_, err = db.Exec(sqlContent)
	if err != nil {
		return nil, nil, fmt.Errorf("failed during creation of table: %w", err)
	}

	teardown := func() {
		db.Close()
		postgresContainer.Terminate(ctx)
	}

	return db, teardown, nil
}
