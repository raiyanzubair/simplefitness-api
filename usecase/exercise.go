package usecase

import (
	"simplefitnessApi/model"
	"simplefitnessApi/repository"
)

//Exercise usecase interface
type Exercise struct {
	repo *repository.Exercise
}

// NewExerciseUsecase creates a new usecase interface for exercises
func NewExerciseUsecase(repo *repository.Exercise) *Exercise {
	return &Exercise{repo}
}

// GetAll returns all exercises
func (uc *Exercise) GetAll() ([]*model.Exercise, error) {
	arr, err := uc.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return arr, nil
}

// GetByID returns a specific Exercise
func (uc *Exercise) GetByID(id int) (*model.Exercise, error) {
	item, err := uc.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

// GetByName returns a specific Exercise
func (uc *Exercise) GetByName(name string) ([]*model.Exercise, error) {
	arr, err := uc.repo.GetByName(name)
	if err != nil {
		return nil, err
	}
	return arr, nil
}
