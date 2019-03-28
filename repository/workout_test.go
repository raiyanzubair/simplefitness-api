package repository_test

import (
	"github.com/stretchr/testify/assert"
	// "simplefitnessApi/model"
	"simplefitnessApi/repository"
	"testing"
)

func TestWorkout_GetAll(t *testing.T) {
	db := repository.PrepTestDB()
	repo := repository.NewWorkoutRepo(db)

	workouts, err := repo.GetAll()
	assert.NoError(t, err, "Fetching should not error")
	assert.NotNil(t, workouts, "Should fetch workouts")
}

func TestWorkout_GetByID(t *testing.T) {
	db := repository.PrepTestDB()
	repo := repository.NewWorkoutRepo(db)

	exercise, err := repo.GetByID(1)
	assert.NoError(t, err, "Fetching should not error")
	assert.NotNil(t, exercise, "Should fetch workouts")
	assert.Equal(t, exercise.ID, 1, "ID should be 1")
}

func TestWorkout_GetByCreator(t *testing.T) {
	db := repository.PrepTestDB()
	repo := repository.NewWorkoutRepo(db)

	workouts, err := repo.GetByCreator(1)
	assert.NoError(t, err, "Fetching should not error")
	assert.NotNil(t, workouts, "Should fetch set")
	assert.Len(t, workouts, 1, "User 1 should have 1 workouts")

	workouts, err = repo.GetByCreator(3)
	assert.NoError(t, err, "Fetching should not error")
	assert.NotNil(t, workouts, "Should fetch set")
	assert.Len(t, workouts, 0, "User 3 should have 0 workouts")
}
