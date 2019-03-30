package usecase_test

import (
	"github.com/stretchr/testify/assert"
	"simplefitnessApi/model"
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

func TestWorkout_GetAll(t *testing.T) {
	uc := prepWorkoutUsecase()
	workouts, err := uc.GetAll()
	assert.NoError(t, err, "Fetching should not error")
	assert.Len(t, workouts, 2, "Should be 2 workouts")
}

func TestWorkout_GetByID(t *testing.T) {
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

func TestWorkout_Create(t *testing.T) {
	uc := prepWorkoutUsecase()

	mock := model.Workout{
		Title: "Monday Memes",
		Day: 0,
		CreatorID: 1,
		Exercises: []*model.WorkoutExercise{
			{
				Exercise: model.Exercise{
					ID: 1,
				},
				Sets: []*model.WorkoutExerciseSet{
					{
						Reps:       12,
						Resistance: 80,
						Duration:   0,
						MeasurementUnit: model.MeasurementUnit{
							ID: 1,
						},
					},
				},
			},
		},
	}

	created, err := uc.Create(&mock)

	assert.NoError(t, err, "Creating should not error")
	assert.NotNil(t, created, "Creating should not return nil")
	assert.Equal(t, mock.Title, created.Title, "Should match inputted struct")
	assert.Equal(t, mock.Day, created.Day, "Should match inputted struct")
	assert.Equal(t, mock.CreatorID, created.CreatorID, "Should match inputted struct")
}

func TestWorkout_Delete(t *testing.T) {
	uc := prepWorkoutUsecase()
	mock := model.Workout{
		Title: "Monday Memes",
		Day: 0,
		CreatorID: 1,
		Exercises: []*model.WorkoutExercise{
			{
				Exercise: model.Exercise{
					ID: 1,
				},
				Sets: []*model.WorkoutExerciseSet{
					{
						Reps:       12,
						Resistance: 80,
						Duration:   0,
						MeasurementUnit: model.MeasurementUnit{
							ID: 1,
						},
					},
				},
			},
		},
	}
	created, err := uc.Create(&mock)
	assert.NoError(t, err, "Creating should not error")
	assert.NotNil(t, created, "Creating should not return nil")
	assert.Equal(t, 10001, created.ID, "Should match inputted struct")

	err = uc.Delete(created.ID)
	assert.NoError(t, err, "Deleting should not error")

	check, err := uc.GetByID(created.ID)
	assert.Error(t, err, "Should return an error as that ID is gone")
	assert.Nil(t, check, "Should return nil")
}
