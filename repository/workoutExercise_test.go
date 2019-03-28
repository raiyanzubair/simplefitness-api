package repository_test

import (
	"github.com/stretchr/testify/assert"
	"simplefitnessApi/model"

	// "simplefitnessApi/model"
	"simplefitnessApi/repository"
	"testing"
)

func TestWorkoutExercise_GetAll(t *testing.T) {
	db := repository.PrepTestDB()
	repo := repository.NewWorkoutExerciseRepo(db)

	exercises, err := repo.GetAll()
	assert.NoError(t, err, "Fetching should not error")
	assert.NotNil(t, exercises, "Should fetch workoutExercise")
}

func TestWorkoutExercise_GetByID(t *testing.T) {
	db := repository.PrepTestDB()
	repo := repository.NewWorkoutExerciseRepo(db)

	exercise, err := repo.GetByID(1)
	assert.NoError(t, err, "Fetching should not error")
	assert.NotNil(t, exercise, "Should fetch workoutExercise")
	assert.Equal(t, exercise.ID, 1, "ID should be 1")
}

func TestWorkoutExercise_GetByWorkout(t *testing.T) {
	db := repository.PrepTestDB()
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

func TestWorkoutExercise_Create(t *testing.T) {
	db := repository.PrepTestDB()
	repo := repository.NewWorkoutExerciseRepo(db)

	toCreate := model.WorkoutExercise{
		WorkoutID: 1,
		Exercise: model.Exercise{
			ID: 1,
		},
	}
	created, err := repo.Create(&toCreate)
	assert.NoError(t, err, "Creating should not error")
	assert.NotNil(t, created, "Creating should not return nil")
	assert.Equal(t, toCreate.WorkoutID, created.WorkoutID, "Should match inputted struct")
	assert.Equal(t, toCreate.Exercise.ID, created.Exercise.ID, "Should match inputted struct")
}

func TestWorkoutExercise_Delete(t *testing.T) {
	db := repository.PrepTestDB()
	repo := repository.NewWorkoutExerciseRepo(db)

	toCreate := model.WorkoutExercise{
		WorkoutID: 1,
		Exercise: model.Exercise{
			ID: 1,
		},
	}
	created, err := repo.Create(&toCreate)
	assert.NoError(t, err, "Creating should not error")
	assert.NotNil(t, created, "Creating should not return nil")
	assert.Equal(t, toCreate.WorkoutID, created.WorkoutID, "Should match inputted struct")
	assert.Equal(t, toCreate.Exercise.ID, created.Exercise.ID, "Should match inputted struct")

	// Now delete it and check it doesnt exist
	err = repo.Delete(created.ID)
	assert.NoError(t, err, "Creating should not error")

	check, err := repo.GetByID(created.ID)
	assert.Error(t, err, "Should return an error as that ID is gone")
	assert.Nil(t, check, "Should return nil")
}
