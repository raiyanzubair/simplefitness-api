package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
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

	db, err := sql.Open("postgres", "user=raiyanzubair dbname=simplefitness sslmode=disable")
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()

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
		db.Exec("INSERT INTO exercise_types (title) VALUES ($1)", item["title"])
	}

	jsonFile, err = os.Open("database/data/exercises.json")
	defer jsonFile.Close()
	byteValue, _ = ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &result)
	for _, item := range result {
		ID := mapTypeToID(item["exercise_type"])
		db.Exec("INSERT INTO exercises (title, exercise_type) VALUES ($1, $2)", item["title"], ID)
	}
}
