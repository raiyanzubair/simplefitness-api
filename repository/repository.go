package repository

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gopkg.in/testfixtures.v2"
	"log"
)

func PrepTestDB() *sqlx.DB {
	db, err := sql.Open("postgres", "user=raiyanzubair dbname=simplefitness_test sslmode=disable")
	if err != nil {
		log.Fatalf("ERROR opening db %v", err)
	}

	fixtures, err := testfixtures.NewFiles(db, &testfixtures.PostgreSQL{},
		"../fixtures/exercise_types.yml",
		"../fixtures/exercises.yml",
		"../fixtures/measurement_units.yml",
		"../fixtures/users.yml",
		"../fixtures/workouts.yml",
		"../fixtures/workout_exercises.yml",
		"../fixtures/workout_exercise_sets.yml",
	)
	if err != nil {
		log.Print(err)
		log.Fatalf("NANI %v", err)
	}
	if err := fixtures.Load(); err != nil {
		log.Fatal(err)
	}

	return sqlx.NewDb(db, "postgres")
}
