package repository_test

import (
	"github.com/stretchr/testify/assert"
	// "simplefitnessApi/model"
	"simplefitnessApi/repository"
	"testing"
)

func TestGetAllWorkoutExercises(t *testing.T) {
	db := prepTestDB()
	repo := repository.NewWorkoutExerciseRepo(db)

	exercises, err := repo.GetAll()
	assert.NoError(t, err, "Fetching should not error")
	assert.NotNil(t, exercises, "Should fetch workoutExercise")
}

func TestGetWorkoutExerciseByID(t *testing.T) {
	db := prepTestDB()
	repo := repository.NewWorkoutExerciseRepo(db)

	exercise, err := repo.GetByID(1)
	assert.NoError(t, err, "Fetching should not error")
	assert.NotNil(t, exercise, "Should fetch workoutExercise")
	assert.Equal(t, exercise.ID, 1, "ID should be 1")
}

func TestGetWorkoutExerciseByWorkout(t *testing.T) {
	db := prepTestDB()
	repo := repository.NewWorkoutExerciseRepo(db)

	exercises, err := repo.GetByWorkout(1)
	assert.NoError(t, err, "Fetching should not error")
	assert.NotNil(t, exercises, "Should fetch set")
	assert.Len(t, exercises, 2, "Workout 1 should have 2 exercises")
	for _, e := range exercises {
		assert.Equal(t, e.WorkoutID, 1, "Set should belong to WorkoutExercise 1")
	}

	exercises, err = repo.GetByWorkout(2)
	assert.NoError(t, err, "Fetching should not error")
	assert.NotNil(t, exercises, "Should fetch set")
	assert.Len(t, exercises, 2, "Workout 1 should have 2 exercises")
	for _, e := range exercises {
		assert.Equal(t, e.WorkoutID, 2, "Set should belong to WorkoutExercise 2")
	}
}
