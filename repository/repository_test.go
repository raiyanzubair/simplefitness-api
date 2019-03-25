package repository_test

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gopkg.in/testfixtures.v2"
	"log"
)

func prepTestDB() *sqlx.DB {
	db, err := sql.Open("postgres", "user=raiyanzubair dbname=simplefitness_test sslmode=disable")
	if err != nil {
		log.Fatalf("ERROR opening db %v", err)
	}

	fixtures, err := testfixtures.NewFiles(db, &testfixtures.PostgreSQL{},
		"../database/fixtures/exercise_types.yml",
		"../database/fixtures/exercises.yml",
		"../database/fixtures/measurement_units.yml",
		"../database/fixtures/workouts.yml",
		"../database/fixtures/workout_exercises.yml",
		"../database/fixtures/workout_exercise_sets.yml",
	)
	if err != nil {
		log.Fatalf("NANI %v", err)
	}
	if err := fixtures.Load(); err != nil {
		log.Fatal(err)
	}

	return sqlx.NewDb(db, "postgres")
}
