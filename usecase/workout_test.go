package usecase_test

import (
	"github.com/stretchr/testify/assert"
	"simplefitnessApi/usecase"

	// "simplefitnessApi/model"
	"simplefitnessApi/repository"
	"testing"
)

func prepWorkoutUsecase() *usecase.Workout {
	db := repository.PrepTestDB()
	wRepo := repository.NewWorkoutRepo(db)
	weRepo := repository.NewWorkoutExerciseRepo(db)
	wesRepo := repository.NewWorkoutExerciseSetRepo(db)
	return usecase.NewWorkoutUsecase(wRepo, weRepo, wesRepo)
}

func TestGetAllWorkouts(t *testing.T) {
	uc := prepWorkoutUsecase()
	workouts, err := uc.GetAll()
	assert.NoError(t, err, "Fetching should not error")
	assert.Len(t, workouts, 2, "Should be 2 workouts")
}
func TestGetWorkoutById(t *testing.T) {
	uc := prepWorkoutUsecase()
	workout, err := uc.GetByID(1)

	assert.NoError(t, err, "Fetching should not error")
	assert.NotNil(t, workout, "Should fetch a single workout")
	assert.Equal(t, workout.CreatorID, 1, "Workout should have Creator 1")
	assert.Len(t, workout.Exercises, 2, "Workout 1 should have 2 Exercises")
	for _, e := range workout.Exercises {
		assert.Equal(t, e.ExerciseType.Title, "Chest", "All Exercises for workout 1 should be Chest")
	}
}
