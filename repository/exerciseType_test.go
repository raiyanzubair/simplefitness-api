package repository_test

import (
	"github.com/stretchr/testify/assert"
	"simplefitnessApi/repository"
	"testing"
)

func TestGetAllExerciseTypes(t *testing.T) {
	db := prepTestDB()
	repo := repository.NewExerciseTypeRepo(db)
	eTypes, err := repo.GetAll()

	assert.NoError(t, err, "Should not error")
	assert.NotNil(t, eTypes, "Should fetch exercise types")
	assert.Len(t, eTypes, 12, "There should be 12 exercise types")
}

func TestGetExerciseTypesByID(t *testing.T) {
	db := prepTestDB()
	repo := repository.NewExerciseTypeRepo(db)

	for i := 1; i < 12; i++ {
		exercise, err := repo.GetByID(i)

		assert.NoError(t, err, "Should not error")
		assert.Equal(t, i, exercise.ID, "Should have matching IDs")
		assert.NotNil(t, exercise, "Should fetch exercise")
	}
}
