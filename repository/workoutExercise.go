package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	// "os"
	"simplefitnessApi/model"
)

//WorkoutExercise repo interface
type WorkoutExercise struct {
	db *sqlx.DB
}

// NewWorkoutExerciseRepo creates a new repo interface for workout exercises
func NewWorkoutExerciseRepo(db *sqlx.DB) *WorkoutExercise {
	return &WorkoutExercise{db}
}

// GetAll returns all exercises for a workout
func (repo *WorkoutExercise) GetAll() ([]*model.WorkoutExercise, error) {
	arr := []*model.WorkoutExercise{}
	err := repo.db.Select(&arr, "SELECT id, title, sets, reps, workout, exercise FROM workout_exercises")
	if err != nil {
		return nil, err
	}
	return arr, nil
}

// GetByID returns a specific WorkoutExercise
func (repo *WorkoutExercise) GetByID(id int) (*model.WorkoutExercise, error) {
	item := model.WorkoutExercise{}
	err := repo.db.Get(&item, "SELECT id, title, sets, reps, workout, exercise FROM workout_exercises WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

// Create will insert a new workout exercise into the db
func (repo *WorkoutExercise) Create(newWorkoutExercise *model.WorkoutExercise) (*model.WorkoutExercise, error) {
	result := repo.db.QueryRow("INSERT INTO workout_exercises (sets, reps, workout, exercise) VALUES ($1, $2, $3, $4) RETURNING id",
		newWorkoutExercise.Sets,
		newWorkoutExercise.Reps,
		newWorkoutExercise.WorkoutID,
		newWorkoutExercise.ExerciseID,
	)
	var id int
	result.Scan(&id)
	workoutExercise, _ := repo.GetByID(id)
	return workoutExercise, nil
}

// Delete removes a workout from the db
func (repo *WorkoutExercise) Delete(id int) error {
	result, err := repo.db.Exec("DELETE FROM workout_exercises WHERE id = $1", id)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}
