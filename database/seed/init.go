package main

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rubenv/sql-migrate"
	"io/ioutil"
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
		fmt.Println("We in dev")
		connStr = "user=raiyanzubair dbname=simplefitness sslmode=disable"
	}


	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()

	//Migrate down and back up
	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}
	n, err := migrate.Exec(db.DB, "postgres", migrations, migrate.Down)
	if err != nil {
		fmt.Print(n, err)
	}
	fmt.Println("Migrated Down")

	n, err = migrate.Exec(db.DB, "postgres", migrations, migrate.Up)
	if err != nil {
		fmt.Print(n, err)
	}
	fmt.Println("Migrated Up")


	//Begin seeding data
	jsonFile, err := os.Open("database/data/exercise_types.json")
	defer jsonFile.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result []map[string]interface{}
	json.Unmarshal(byteValue, &result)
	for _, item := range result {
		res, err := db.Exec("INSERT INTO exercise_types (title) VALUES ($1)", item["title"])
		if err != nil {
			fmt.Print(res)
			fmt.Print(err)
		}
	}
	fmt.Println("Seeded Exercise Types")

	jsonFile, err = os.Open("database/data/exercises.json")
	defer jsonFile.Close()
	byteValue, _ = ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &result)
	for _, item := range result {
		ID := mapTypeToID(item["exercise_type"])
		db.Exec("INSERT INTO exercises (title, exercise_type) VALUES ($1, $2)", item["title"], ID)
	}
	fmt.Println("Seeded Exercises")
}
