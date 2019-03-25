package repository_test

import (
	"github.com/stretchr/testify/assert"
	"simplefitnessApi/repository"
	"testing"
)

func TestGetAllExercises(t *testing.T) {
	db := repository.PrepTestDB()
	repo := repository.NewExerciseRepo(db)

	exercises, err := repo.GetAll()

	assert.NoError(t, err, "Should not error")
	assert.NotNil(t, exercises, "Should fetch exercises")
	assert.Len(t, exercises, 1299, "There should be 1299 exercises")
}

func TestGetExerciseByID(t *testing.T) {
	db := repository.PrepTestDB()
	repo := repository.NewExerciseRepo(db)

	for i := 1; i < 100; i++ {
		exercise, err := repo.GetByID(i)

		assert.NoError(t, err, "Should not error")
		assert.Equal(t, i, exercise.ID, "Should have matching IDs")
		assert.NotNil(t, exercise, "Should fetch exercise")
	}
}
