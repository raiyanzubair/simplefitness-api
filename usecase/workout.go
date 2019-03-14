package usecase

import (
	"simplefitnessApi/model"
	"simplefitnessApi/repository"
)

//Workout usecase interface
type Workout struct {
	wRepo  *repository.Workout
	weRepo *repository.WorkoutExercise
	eRepo  *repository.Exercise
}

// NewWorkoutUsecase creates a new usecase interface for workouts
func NewWorkoutUsecase(wRepo *repository.Workout, weRepo *repository.WorkoutExercise, eRepo *repository.Exercise) *Workout {
	return &Workout{wRepo, weRepo, eRepo}
}

// Get exercise details for a workout based of a
func (uc *Workout) getWorkoutExercises(workoutID int) ([]model.WorkoutExercise, error) {
	// Get workout exercises for the workout
	result := []model.WorkoutExercise{}
	workoutExercises, err := uc.weRepo.GetByWorkout(workoutID)

	if err != nil {
		return nil, err
	}
	// for each workout exercise append its exercise object
	for _, we := range workoutExercises {
		e, _ := uc.eRepo.GetByID(we.Exercise.ID)
		we.Exercise = *e
		result = append(result, *we)
	}
	return result, nil
}

// GetAll returns all workouts and joins their associated workout exercises
func (uc *Workout) GetAll() ([]*model.Workout, error) {
	workouts, err := uc.wRepo.GetAll()
	if err != nil {
		return nil, err
	}

	for i, w := range workouts {
		result, _ := uc.getWorkoutExercises(w.ID)

		workouts[i].Exercises = []model.WorkoutExercise{}
		if result != nil {
			workouts[i].Exercises = result
		}
	}
	return workouts, nil
}

// GetByID returns a specific Workout and joins its associated workout exercises
func (uc *Workout) GetByID(workoutID int) (*model.Workout, error) {
	workout, err := uc.wRepo.GetByID(workoutID)
	if err != nil {
		return nil, err
	}
	workout.Exercises, err = uc.getWorkoutExercises(workoutID)
	if err != nil {
		return nil, err
	}
	return workout, nil
}

// GetByCreator returns all workouts for a specific user
func (uc *Workout) GetByCreator(creatorID int) ([]*model.Workout, error) {
	workouts, err := uc.wRepo.GetByCreator(creatorID)
	if err != nil {
		return nil, err
	}

	for i, w := range workouts {
		result, _ := uc.getWorkoutExercises(w.ID)

		workouts[i].Exercises = []model.WorkoutExercise{}
		if result != nil {
			workouts[i].Exercises = result
		}
	}
	return workouts, nil
}

// Create creates a new workout
func (uc *Workout) Create(newWorkout *model.Workout) (*model.Workout, error) {
	created, err := uc.wRepo.Create(newWorkout)
	if err != nil {
		return nil, err
	}
	return created, nil
}

// Delete deletes a workout
func (uc *Workout) Delete(id int) error {
	err := uc.wRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
