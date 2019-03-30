package usecase

import (
	"simplefitnessApi/model"
	"simplefitnessApi/repository"
)

//Workout usecase interface
type Workout struct {
	wRepo   *repository.Workout
	weRepo  *repository.WorkoutExercise
	wesRepo *repository.WorkoutExerciseSet
}

// NewWorkoutUsecase creates a new usecase interface for workouts
func NewWorkoutUsecase(wRepo *repository.Workout, weRepo *repository.WorkoutExercise, wesRepo *repository.WorkoutExerciseSet) *Workout {
	return &Workout{wRepo, weRepo, wesRepo}
}

// Get workout exercises for a workout. And for each workout exercise get their sets.
func (uc *Workout) getWorkoutExercises(workoutID int) ([]*model.WorkoutExercise, error) {
	// Get workout exercises for the workout
	workoutExercises, err := uc.weRepo.GetByWorkout(workoutID)
	if err != nil {
		return nil, err
	}
	// for each workout exercise get the sets
	for i, we := range workoutExercises {
		sets, err := uc.wesRepo.GetByWorkoutExercise(we.ID)
		if err != nil {
			return nil, err
		}
		workoutExercises[i].Sets = sets
	}
	return workoutExercises, nil
}

// GetAll returns all workouts and joins their associated workout exercises
func (uc *Workout) GetAll() ([]*model.Workout, error) {
	workouts, err := uc.wRepo.GetAll()
	if err != nil {
		return nil, err
	}

	for i, w := range workouts {
		result, _ := uc.getWorkoutExercises(w.ID)

		workouts[i].Exercises = []*model.WorkoutExercise{}
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

	for _, w := range workouts {
		result, err := uc.getWorkoutExercises(w.ID)
		if err != nil {
			return nil, err
		}
		w.Exercises = result
		//workouts[i].Exercises = []*model.WorkoutExercise{}

	}
	return workouts, nil
}

/*
Create will create a new workout
Attach sets and exercises to it
*/
func (uc *Workout) Create(w *model.Workout) (*model.Workout, error) {
	// create base workout
	newWorkout, err := uc.wRepo.Create(w)
	if err != nil {
		return nil, err
	}

	// create workout exercises
	for _, we := range w.Exercises {
		we.WorkoutID = newWorkout.ID
		newWorkoutExercise, err := uc.weRepo.Create(we)
		if err != nil {
			return nil, err
		}

		// create workout exercise sets
		for _, wes := range we.Sets{
			wes.WorkoutExerciseID = newWorkoutExercise.ID
			newSet,err := uc.wesRepo.Create(wes)

			if err != nil {
				return nil, err
			}
			newWorkoutExercise.Sets = append(newWorkoutExercise.Sets, newSet)
		}
		newWorkout.Exercises = append(newWorkout.Exercises, newWorkoutExercise)
	}
	return newWorkout, nil
}

// Delete deletes a workout
func (uc *Workout) Delete(id int) error {
	err := uc.wRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
