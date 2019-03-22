package repository

import (
	"github.com/jmoiron/sqlx"
	"simplefitnessApi/model"
)

//WorkoutExerciseSet repo interface
type WorkoutExerciseSet struct {
	db *sqlx.DB
}

// NewWorkoutExerciseSetRepo creates a new repo interface for workout exercises
func NewWorkoutExerciseSetRepo(db *sqlx.DB) *WorkoutExerciseSet {
	return &WorkoutExerciseSet{db}
}

// GetByID returns a specific WorkoutExerciseSet
func (repo *WorkoutExerciseSet) GetByID(id int) (*model.WorkoutExerciseSet, error) {
	workoutExercise := model.WorkoutExerciseSet{}
	query := `
		SELECT 
			wes.id, wes.workout_exercise, wes.reps, wes.resistance, wes.duration,
			mu.id as "workout_exercise_set.measurement_unit.id", mu.title as "workout_exercise_set.measurement_unit.title"
		FROM workout_exercise_sets as wes
		JOIN measurement_units as mu ON wes.measurement_unit = mu.id
		WHERE id = $1
	`
	err := repo.db.Get(&workoutExercise, query, id)
	if err != nil {
		return nil, err
	}
	return &workoutExercise, nil
}

// GetByWorkoutExercise returns all WorkoutExerciseSets for a WorkoutExercise
func (repo *WorkoutExerciseSet) GetByWorkoutExercise(workoutExerciseID int) ([]*model.WorkoutExerciseSet, error) {
	workoutExerciseSets := []*model.WorkoutExerciseSet{}
	query := `
		SELECT 
			wes.id, wes.workout_exercise, wes.reps, wes.resistance, wes.duration,
			mu.id as "workout_exercise_set.measurement_unit.id", mu.title as "workout_exercise_set.measurement_unit.title"
		FROM workout_exercise_sets as wes
		JOIN measurement_units as mu ON wes.measurement_unit = mu.id
		WHERE workout_exercise = $1
	`
	err := repo.db.Select(&workoutExerciseSets, query, workoutExerciseID)
	if err != nil {
		return nil, err
	}
	return workoutExerciseSets, nil
}
