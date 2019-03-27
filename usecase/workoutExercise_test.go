package usecase_test

import (
	"github.com/stretchr/testify/assert"
	"simplefitnessApi/usecase"

	// "simplefitnessApi/model"
	"simplefitnessApi/repository"
	"testing"
)

func prepWorkoutExerciseUsecase() *usecase.WorkoutExercise {
	db := repository.PrepTestDB()
	weRepo := repository.NewWorkoutExerciseRepo(db)
	wesRepo := repository.NewWorkoutExerciseSetRepo(db)
	return usecase.NewWorkoutExerciseUsecase(weRepo, wesRepo)
}

func TestGetAllWorkoutExercises(t *testing.T) {
	uc := prepWorkoutExerciseUsecase()
	workoutExercises, err := uc.GetAll()
	assert.NoError(t, err, "Fetching should not error")
	assert.Len(t, workoutExercises, 4, "Should be 4 workout exercises")
}
func TestGetWorkoutExerciseById(t *testing.T) {
	uc := prepWorkoutExerciseUsecase()
	workoutExercise, err := uc.GetByID(1)
	assert.NoError(t, err, "Fetching should not error")
	assert.NotNil(t, workoutExercise, "Should fetch a single workoutExercise")
	assert.Equal(t, workoutExercise.WorkoutID, 1, "WorkoutExercise 1 should belong to Workout 1")
	assert.Len(t, workoutExercise.Sets, 2, "WorkoutExercise 1 should have 2 Sets")
}
