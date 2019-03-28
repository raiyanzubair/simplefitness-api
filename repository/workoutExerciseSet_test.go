package repository_test

import (
	"github.com/stretchr/testify/assert"
	"simplefitnessApi/model"

	// "simplefitnessApi/model"
	"simplefitnessApi/repository"
	"testing"
)

func TestWorkoutExerciseSet_GetByID(t *testing.T) {
	db := repository.PrepTestDB()
	repo := repository.NewWorkoutExerciseSetRepo(db)

	set, err := repo.GetByID(1)
	assert.NoError(t, err, "Fetching should not error")
	assert.NotNil(t, set, "Should fetch set")
	assert.Equal(t, set.ID, 1, "ID should be 1")
}

func TestWorkoutExerciseSet_GetByWorkoutExercise(t *testing.T) {
	db := repository.PrepTestDB()
	repo := repository.NewWorkoutExerciseSetRepo(db)

	sets, err := repo.GetByWorkoutExercise(4)
	assert.NoError(t, err, "Fetching should not error")
	assert.NotNil(t, sets, "Should fetch sets")
	assert.Len(t, sets, 3, "WorkoutExercise 4 should have 4 sets")
	for _, set := range sets {
		assert.Equal(t, set.WorkoutExerciseID, 4, "Set should belong to WorkoutExercise 4")
	}
}

func TestWorkoutExerciseSet_Create(t *testing.T) {
	db := repository.PrepTestDB()
	repo := repository.NewWorkoutExerciseSetRepo(db)

	toCreate := model.WorkoutExerciseSet{
		WorkoutExerciseID:1,
		Reps: 12,
		Resistance: 2,
		MeasurementUnit: model.MeasurementUnit{
			ID: 1,
			Title: "kg",
		},
		Duration: 0,
	}
	created,err := repo.Create(&toCreate)
	assert.NoError(t, err, "Creating should not error")
	assert.NotNil(t, created, "Creating should not return nil")
	assert.Equal(t,toCreate.WorkoutExerciseID, created.WorkoutExerciseID, "Should match inputted struct")
	assert.Equal(t,toCreate.Reps, created.Reps, "Should match inputted struct")
	assert.Equal(t,toCreate.Resistance, created.Resistance, "Should match inputted struct")
	assert.Equal(t,toCreate.MeasurementUnit, created.MeasurementUnit, "Should match inputted struct")
	assert.Equal(t,toCreate.Duration, created.Duration, "Should match inputted struct")
}

func TestWorkoutExerciseSet_Delete(t *testing.T) {
	db := repository.PrepTestDB()
	repo := repository.NewWorkoutExerciseSetRepo(db)

	toCreate := model.WorkoutExerciseSet{
		WorkoutExerciseID:1,
		Reps: 12,
		Resistance: 2,
		MeasurementUnit: model.MeasurementUnit{
			ID: 1,
			Title: "kg",
		},
		Duration: 0,
	}
	created,err := repo.Create(&toCreate)
	assert.NoError(t, err, "Creating should not error")
	assert.NotNil(t, created, "Creating should not return nil")
	assert.Equal(t,toCreate.WorkoutExerciseID, created.WorkoutExerciseID, "Should match inputted struct")
	assert.Equal(t,toCreate.Reps, created.Reps, "Should match inputted struct")
	assert.Equal(t,toCreate.Resistance, created.Resistance, "Should match inputted struct")
	assert.Equal(t,toCreate.MeasurementUnit, created.MeasurementUnit, "Should match inputted struct")
	assert.Equal(t,toCreate.Duration, created.Duration, "Should match inputted struct")

	// Now delete it and check it doesnt exist
	err = repo.Delete(created.ID)
	assert.NoError(t, err, "Creating should not error")

	check, err := repo.GetByID(created.ID)
	assert.Error(t, err, "Should return an error as that ID is gone")
	assert.Nil(t, check, "Should return nil")
}
