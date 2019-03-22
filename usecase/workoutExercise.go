package usecase

import (
	"simplefitnessApi/model"
	"simplefitnessApi/repository"
)

//WorkoutExercise usecase interface
type WorkoutExercise struct {
	repo *repository.WorkoutExercise
}

// NewWorkoutExerciseUsecase creates a new usecase interface for exercises
func NewWorkoutExerciseUsecase(repo *repository.WorkoutExercise) *WorkoutExercise {
	return &WorkoutExercise{repo}
}

// GetAll returns all exercises
func (uc *WorkoutExercise) GetAll() ([]*model.WorkoutExercise, error) {
	arr, err := uc.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return arr, nil
}

// GetByID returns a specific WorkoutExercise
func (uc *WorkoutExercise) GetByID(id int) (*model.WorkoutExercise, error) {
	item, err := uc.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

// // Create creates a new workout
// func (uc *WorkoutExercise) Create(newWorkoutExercise *model.WorkoutExercise) (*model.WorkoutExercise, error) {
// 	created, err := uc.repo.Create(newWorkoutExercise)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return created, nil
// }

// // Delete deletes a workout
// func (uc *WorkoutExercise) Delete(id int) error {
// 	err := uc.repo.Delete(id)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
