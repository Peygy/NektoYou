package managers

import (
	"database/sql"
	"log"
	"os"
	"sync"
	"testing"

	"github.com/peygy/nektoyou/internal/pkg/mocks"
)

var (
	dbs       map[string]*sql.DB
	teardowns map[string]func()
	once      sync.Once
)

const schemaFilePath = "../../config/schema.sql"

func setupDB() (*sql.DB, func(), error) {
	db, teardown, err := mocks.SetupTestContainer_Postgres(schemaFilePath)
	if err != nil {
		return nil, nil, err
	}
	return db, teardown, nil
}

func TestMain(m *testing.M) {
	once.Do(func() {
		dbs = make(map[string]*sql.DB)
		teardowns = make(map[string]func())

		services := []string{"IRoleManager", "IUserManager"}

		for _, service := range services {
			db, teardown, err := setupDB()
			if err != nil {
				log.Fatalf("Could not set up test container for %s: %v", service, err)
			}
			dbs[service] = db
			teardowns[service] = teardown
		}
	})

	code := m.Run()

	for _, teardown := range teardowns {
		if teardown != nil {
			teardown()
		}
	}

	os.Exit(code)
}
