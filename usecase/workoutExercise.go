package usecase

import (
	"simplefitnessApi/model"
	"simplefitnessApi/repository"
)

//WorkoutExercise usecase interface
type WorkoutExercise struct {
	wRepo *repository.WorkoutExercise
	wesRepo *repository.WorkoutExerciseSet
}

// NewWorkoutExerciseUsecase creates a new usecase interface for exercises
func NewWorkoutExerciseUsecase(wRepo *repository.WorkoutExercise, wesRepo *repository.WorkoutExerciseSet) *WorkoutExercise {
	return &WorkoutExercise{wRepo, wesRepo}
}

// GetAll returns all exercises
func (uc *WorkoutExercise) GetAll() ([]*model.WorkoutExercise, error) {
	slice, err := uc.wRepo.GetAll()
	if err != nil {
		return nil, err
	}
	for i, we := range slice{
		sets, err := uc.wesRepo.GetByWorkoutExercise(we.ID)
		if err != nil {
			return nil, err
		}
		slice[i].Sets = sets
	}
	return slice, nil
}

// GetByID returns a specific WorkoutExercise
func (uc *WorkoutExercise) GetByID(id int) (*model.WorkoutExercise, error) {
	item, err := uc.wRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	sets, err := uc.wesRepo.GetByWorkoutExercise(item.ID)
	if err != nil {
		return nil, err
	}
	item.Sets = sets
	return item, nil
}

// // Create creates a new workout
// func (uc *WorkoutExercise) Create(newWorkoutExercise *model.WorkoutExercise) (*model.WorkoutExercise, error) {
// 	created, err := uc.wRepo.Create(newWorkoutExercise)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return created, nil
// }

// // Delete deletes a workout
// func (uc *WorkoutExercise) Delete(id int) error {
// 	err := uc.wRepo.Delete(id)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
