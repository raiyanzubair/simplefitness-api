package repository_test

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"gopkg.in/testfixtures.v2"
	"log"
	"simplefitnessApi/repository"
	"testing"
)

var (
	db       *sql.DB
	fixtures *testfixtures.Context
)

func prepTestRepo() *repository.Exercise {
	var err error
	db, err = sql.Open("postgres", "user=raiyanzubair dbname=simplefitness_test sslmode=disable")
	if err != nil {
		log.Fatalf("ERROR opening db %v", err)
	}

	fixtures, err = testfixtures.NewFiles(db, &testfixtures.PostgreSQL{},
		"../database/fixtures/exercise_types.yml",
		"../database/fixtures/exercises.yml",
	)
	if err != nil {
		log.Fatalf("NANI %v", err)
	}
	if err := fixtures.Load(); err != nil {
		log.Fatal(err)
	}

	repo := repository.NewExerciseRepo(sqlx.NewDb(db, "sqlmock"))
	return repo
}

func TestGetAll(t *testing.T) {
	repo := prepTestRepo()
	exercises, err := repo.GetAll()

	assert.NoError(t, err, "Should not error")
	assert.NotNil(t, exercises, "Should fetch exercises")
	assert.Len(t, exercises, 1299, "There should be 1299 exercises")
}

func TestGetByID(t *testing.T) {
	repo := prepTestRepo()
	for i := 1; i < 1300; i++ {
		go func(j int) {
			exercise, err := repo.GetByID(j)
			assert.NoError(t, err, "Should not error")
			assert.Equal(t, j, exercise.ID, "Should have matching IDs")
			assert.NotNil(t, exercise, "Should fetch exercise")
		}(i)
	}
}
