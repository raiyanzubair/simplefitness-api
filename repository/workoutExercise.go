package repository

import (
	"github.com/jmoiron/sqlx"
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
	workoutExercises := []*model.WorkoutExercise{}
	query := `
		SELECT 
			we.id, we.workout,
			e.id as "exercise.id", e.title as "exercise.title",
			et.id as "exercise.exercise_type.id", et.title as "exercise.exercise_type.title"
		FROM workout_exercises as we
		JOIN exercises as e ON we.exercise = e.id
		JOIN exercise_types as et ON e.exercise_type = et.id
	`
	err := repo.db.Select(&workoutExercises, query)
	if err != nil {
		return nil, err
	}
	return workoutExercises, nil
}

// GetByID returns a specific WorkoutExercise
func (repo *WorkoutExercise) GetByID(id int) (*model.WorkoutExercise, error) {
	workoutExercise := model.WorkoutExercise{}
	query := `
		SELECT 
			we.id, we.workout,
			e.id as "exercise.id", e.title as "exercise.title",
			et.id as "exercise.exercise_type.id", et.title as "exercise.exercise_type.title"
		FROM workout_exercises as we
		JOIN exercises as e ON we.exercise = e.id
		JOIN exercise_types as et ON e.exercise_type = et.id
		WHERE we.id = $1
	`
	err := repo.db.Get(&workoutExercise, query, id)
	if err != nil {
		return nil, err
	}
	return &workoutExercise, nil
}

// GetByWorkout returns WorkoutExercises for a specific workout
func (repo *WorkoutExercise) GetByWorkout(workoutID int) ([]*model.WorkoutExercise, error) {
	workoutExercises := []*model.WorkoutExercise{}
	query := `
		SELECT 
			we.id, we.workout,
			e.id as "exercise.id", e.title as "exercise.title",
			et.id as "exercise.exercise_type.id", et.title as "exercise.exercise_type.title"
		FROM workout_exercises as we
		JOIN exercises as e ON we.exercise = e.id
		JOIN exercise_types as et ON e.exercise_type = et.id
		WHERE we.workout = $1
	`
	err := repo.db.Select(&workoutExercises, query, workoutID)
	if err != nil {
		return nil, err
	}

	return workoutExercises, nil
}

// // Create will insert a new workout exercise into the db
// func (repo *WorkoutExercise) Create(newWorkoutExercise *model.WorkoutExercise) (*model.WorkoutExercise, error) {
// 	query := `
// 		INSERT INTO workout_exercises (sets, reps, workout, exercise)
// 		VALUES ($1, $2, $3, $4)
// 		RETURNING id
// 	`
// 	result := repo.db.QueryRow(query,
// 		newWorkoutExercise.Sets,
// 		newWorkoutExercise.Reps,
// 		newWorkoutExercise.WorkoutID,
// 	)
// 	var id int
// 	result.Scan(&id)
// 	workoutExercise, _ := repo.GetByID(id)
// 	return workoutExercise, nil
// }

// // Delete removes a workout from the db
// func (repo *WorkoutExercise) Delete(id int) error {
// 	query := `
// 		DELETE FROM workout_exercises
// 		WHERE id = $1
// 	`
// 	result, err := repo.db.Exec(query, id)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
