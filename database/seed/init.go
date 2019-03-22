package main

import (
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
	"github.com/rubenv/sql-migrate"
	"gopkg.in/testfixtures.v2"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	connStr := os.Getenv("DATABASE_URL")
	env := os.Getenv("GO_ENV")
	if connStr == "" || env == "development" {
		log.Println("We in dev")
		connStr = "user=raiyanzubair dbname=simplefitness sslmode=disable"
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	testDb, err := sql.Open("postgres", "user=raiyanzubair dbname=simplefitness_test sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer testDb.Close()

	performMigrations(db)
	performMigrations(testDb)

	performSeeding(db)

	generateFixturesForTesting(db)
}

func performMigrations(db *sql.DB) {
	//Migrate down and back up
	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}
	_, err := migrate.Exec(db, "postgres", migrations, migrate.Down)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Migrated Down")

	_, err = migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Migrated Up")
}

func performSeeding(db *sql.DB) {
	jsonFile, err := os.Open("database/data/exercise_types.json")
	defer jsonFile.Close()
	if err != nil {
		log.Println(err)
		return
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result []map[string]interface{}
	json.Unmarshal(byteValue, &result)
	for _, item := range result {
		db.Exec("INSERT INTO exercise_types (id, title) VALUES ($1, $2)", item["id"], item["title"])
	}

	jsonFile, err = os.Open("database/data/exercises.json")
	byteValue, _ = ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &result)
	for _, item := range result {
		db.Exec("INSERT INTO exercises (id, title, exercise_type) VALUES ($1, $2, $3)", item["id"], item["title"], item["exercise_type"])
	}

	jsonFile, err = os.Open("database/data/measurement_units.json")
	byteValue, _ = ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &result)
	for _, item := range result {
		db.Exec("INSERT INTO measurement_units (id, title) VALUES ($1, $2)", item["id"], item["title"])
	}
}

func generateFixturesForTesting(db *sql.DB) {
	err := testfixtures.GenerateFixturesForTables(
		db,
		[]*testfixtures.TableInfo{
			&testfixtures.TableInfo{Name: "exercise_types"},
			&testfixtures.TableInfo{Name: "exercises"},
		},
		&testfixtures.PostgreSQL{},
		"database/fixtures/",
	)
	if err != nil {
		log.Fatalf("Error generating fixtures: %v", err)
	}
	log.Print("GENERATED FIXTURES")
}
