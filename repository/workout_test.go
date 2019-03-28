package repository_test

import (
	"github.com/stretchr/testify/assert"
	"simplefitnessApi/model"

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

func TestWorkout_Create	(t *testing.T) {
	db := repository.PrepTestDB()
	repo := repository.NewWorkoutRepo(db)

	toCreate := model.Workout{
		Title: "Monday Memes",
		Day: 0,
		CreatorID: 1,
	}
	created, err := repo.Create(&toCreate)
	assert.NoError(t, err, "Creating should not error")
	assert.NotNil(t, created, "Creating should not return nil")
	assert.Equal(t, toCreate.Title, created.Title, "Should match inputted struct")
	assert.Equal(t, toCreate.Day, created.Day, "Should match inputted struct")
	assert.Equal(t, toCreate.CreatorID, created.CreatorID, "Should match inputted struct")
}

func TestWorkout_Delete(t *testing.T) {
	db := repository.PrepTestDB()
	repo := repository.NewWorkoutRepo(db)

	toCreate := model.Workout{
		Title: "Monday Memes",
		Day: 0,
		CreatorID: 1,
	}
	created, err := repo.Create(&toCreate)
	assert.NoError(t, err, "Creating should not error")
	assert.NotNil(t, created, "Creating should not return nil")
	assert.Equal(t, 10001, created.ID, "Should match inputted struct")
	assert.Equal(t, toCreate.Title, created.Title, "Should match inputted struct")
	assert.Equal(t, toCreate.Day, created.Day, "Should match inputted struct")
	assert.Equal(t, toCreate.CreatorID, created.CreatorID, "Should match inputted struct")

	// Now delete it and check it doesnt exist
	err = repo.Delete(created.ID)
	assert.NoError(t, err, "Creating should not error")

	check, err := repo.GetByID(created.ID)
	assert.Error(t, err, "Should return an error as that ID is gone")
	assert.Nil(t, check, "Should return nil")
}
