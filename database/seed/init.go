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

func mapTypeToID(title interface{}) int {
	exerciseMap := map[interface{}]int{
		"Chest":     1,
		"Back":      2,
		"Shoulders": 3,
		"Biceps":    4,
		"Triceps":   5,
		"Forearms":  6,
		"Quads":     7,
		"Calves":    8,
		"Glutes":    9,
		"Abs":       10,
		"Cardio":    11,
		"Other":     12,
	}
	return exerciseMap[title]
}

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

	//Migrate down and back up
	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}
	_, err = migrate.Exec(db, "postgres", migrations, migrate.Down)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Migrated Down")
	_, err = migrate.Exec(testDb, "postgres", migrations, migrate.Down)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Migrated Down for test db")

	_, err = migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Migrated Up")
	_, err = migrate.Exec(testDb, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Migrated Up for test db")

	//Begin seeding data
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
		db.Exec("INSERT INTO exercise_types (title) VALUES ($1)", item["title"])
	}
	log.Println("Seeded Exercise Types")

	jsonFile, err = os.Open("database/data/exercises.json")
	defer jsonFile.Close()
	byteValue, _ = ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &result)
	for _, item := range result {
		ID := mapTypeToID(item["exercise_type"])
		db.Exec("INSERT INTO exercises (title, exercise_type) VALUES ($1, $2)", item["title"], ID)
	}
	log.Println("Seeded Exercises")

	generateFixturesForTesting(db)
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
