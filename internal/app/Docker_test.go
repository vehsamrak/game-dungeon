package app_test

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/ory/dockertest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"log"
	"testing"
)

func TestDocker(test *testing.T) {
	suite.Run(test, &dockerTest{})
}

type dockerTest struct {
	suite.Suite
	database *sql.DB
	resource *dockertest.Resource
	pool     *dockertest.Pool
}

func (suite *dockerTest) SetupSuite() {
	suite.database, suite.resource, suite.pool = suite.startPostgres()
}

func (suite *dockerTest) TearDownSuite() {
	suite.stopPostgres(suite.resource, suite.pool)
}

func (suite *dockerTest) Test_PostgresSelect_dockerWithPostgres_SQLSelectSucceeded() {
	var selectResult int
	rows := suite.database.QueryRow("SELECT 10")
	err := rows.Scan(&selectResult)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 10, selectResult)
}

func (suite *dockerTest) stopPostgres(resource *dockertest.Resource, pool *dockertest.Pool) {
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
}

func (suite *dockerTest) startPostgres() (database *sql.DB, resource *dockertest.Resource, pool *dockertest.Pool) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	resource, err = pool.Run("postgres", "9.6", []string{"POSTGRES_PASSWORD=secret", "POSTGRES_DB=test"})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	if err = pool.Retry(func() error {
		var err error
		database, err = sql.Open(
			"postgres",
			fmt.Sprintf(
				"postgres://postgres:secret@localhost:%s/%s?sslmode=disable",
				resource.GetPort("5432/tcp"),
				"test",
			),
		)
		if err != nil {
			return err
		}
		return database.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	return
}
