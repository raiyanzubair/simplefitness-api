package usecase

import (
	"simplefitnessApi/model"
	"simplefitnessApi/repository"
)

//ExerciseType usecase interface
type ExerciseType struct {
	repo *repository.ExerciseType
}

// NewExerciseTypeUsecase creates a new usecase interface for exercise_types
func NewExerciseTypeUsecase(repo *repository.ExerciseType) *ExerciseType {
	return &ExerciseType{repo}
}

// GetAll returns all exercise_types
func (uc *ExerciseType) GetAll() ([]*model.ExerciseType, error) {
	slice, err := uc.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return slice, nil
}

// GetByID returns a specific ExerciseType
func (uc *ExerciseType) GetByID(id int) (*model.ExerciseType, error) {
	item, err := uc.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}
