package repository_test

import (
	"github.com/stretchr/testify/assert"
	// "simplefitnessApi/model"
	"simplefitnessApi/repository"
	"testing"
)

func TestGetByID(t *testing.T) {
	db := prepTestDB()
	repo := repository.NewWorkoutExerciseSetRepo(db)

	set, err := repo.GetByID(1)
	assert.NoError(t, err, "Fetching should not error")
	assert.NotNil(t, set, "Should fetch set")
	assert.Equal(t, set.ID, 1, "ID should be 1")
}

func TestGetByWorkoutExercise(t *testing.T) {
	db := prepTestDB()
	repo := repository.NewWorkoutExerciseSetRepo(db)

	sets, err := repo.GetByWorkoutExercise(4)
	assert.NoError(t, err, "Fetching should not error")
	assert.NotNil(t, sets, "Should fetch sets")
	assert.Len(t, sets, 3, "WorkoutExercise 4 should have 4 sets")
	for _, set := range sets {
		assert.Equal(t, set.WorkoutExerciseID, 4, "Set should belong to WorkoutExercise 4")
	}
}
