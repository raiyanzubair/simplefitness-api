package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	// "os"
	"simplefitnessApi/model"
)

//Workout repo interface
type Workout struct {
	db *sqlx.DB
}

// NewWorkoutRepo creates a new repo interface for workouts
func NewWorkoutRepo(db *sqlx.DB) *Workout {
	return &Workout{db}
}

// GetAll returns all Workouts
func (repo *Workout) GetAll() ([]*model.Workout, error) {
	arr := []*model.Workout{}
	err := repo.db.Select(&arr, "SELECT id, title workouts FROM Workout")
	if err != nil {
		return nil, err
	}
	return arr, nil
}

// GetByID returns a specific Workout
func (repo *Workout) GetByID(id int) (*model.Workout, error) {
	item := model.Workout{}
	err := repo.db.Get(&item, "SELECT id, title FROM workouts WHERE id = $1", id)
	if err != nil {
		fmt.Printf(err.Error())
		return nil, err
	}
	return &item, nil
}

// Create will create insert a new workout into the db and return it
func (repo *Workout) Create(newWorkout *model.Workout) (*model.Workout, error) {
	result := repo.db.QueryRow("INSERT INTO workouts (title, day, creator) VALUES ($1, $2, $3) RETURNING id",
		newWorkout.Title,
		newWorkout.Day,
		newWorkout.CreatorID,
	)
	var id int
	result.Scan(&id)
	workout, _ := repo.GetByID(int(id))
	return workout, nil
}

// Delete removes a workout from the db
func (repo *Workout) Delete(id int) error {
	result, err := repo.db.Exec("DELETE FROM workout WHERE id = $1", id)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}
