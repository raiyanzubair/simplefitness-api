package repository

import (
	"github.com/jmoiron/sqlx"
	"log"
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
	workoutExercise := model.WorkoutExerciseSet{
		MeasurementUnit: model.MeasurementUnit{},
	}
	query := `
		SELECT 
			wes.id, wes.workout_exercise, wes.reps, wes.resistance, wes.duration,
			mu.id as "measurement_unit.id", mu.title as "measurement_unit.title"
		FROM workout_exercise_sets as wes
		JOIN measurement_units as mu ON wes.measurement_unit = mu.id
		WHERE wes.id = $1
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
			mu.id as "measurement_unit.id", mu.title as "measurement_unit.title"
		FROM workout_exercise_sets as wes
		JOIN measurement_units as mu ON wes.measurement_unit = mu.id
		WHERE wes.workout_exercise = $1
	`
	err := repo.db.Select(&workoutExerciseSets, query, workoutExerciseID)
	if err != nil {
		return nil, err
	}
	return workoutExerciseSets, nil
}

// Create will insert a new set into the db and return it
func (repo *WorkoutExerciseSet) Create(newSet *model.WorkoutExerciseSet) (*model.WorkoutExerciseSet, error) {
	query := `
		INSERT INTO workout_exercise_sets (workout_exercise, reps, resistance, measurement_unit, duration)
		VALUES ($1, $2, $3, $4, $5) RETURNING workout_exercise_sets.id
	`
	result := repo.db.QueryRow(query,
		newSet.WorkoutExerciseID,
		newSet.Reps,
		newSet.Resistance,
		newSet.MeasurementUnit.ID,
		newSet.Duration,
	)

	var id int
	err := result.Scan(&id)
	if err != nil {
		log.Print(newSet)
		log.Print(id)
		return nil, err
	}
	set, err := repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return set, nil
}

// Delete removes a workout from the db
func (repo *WorkoutExerciseSet) Delete(id int) error {
	query := `
		DELETE FROM workout_exercise_sets
		WHERE id = $1
	`
	_, err := repo.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}
