package usecase

import (
	"simplefitnessApi/model"
	"simplefitnessApi/repository"
)

//Workout usecase interface
type Workout struct {
	repo *repository.Workout
}

// NewWorkoutUsecase creates a new usecase interface for exercise_types
func NewWorkoutUsecase(repo *repository.Workout) *Workout {
	return &Workout{repo}
}

// GetAll returns all exercise_types
func (uc *Workout) GetAll() ([]*model.Workout, error) {
	arr, err := uc.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return arr, nil
}

// GetByID returns a specific Workout
func (uc *Workout) GetByID(id int) (*model.Workout, error) {
	item, err := uc.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

// Create creates a new workout
func (uc *Workout) Create(newWorkout *model.Workout) (*model.Workout, error) {
	created, err := uc.repo.Create(newWorkout)
	if err != nil {
		return nil, err
	}
	return created, nil
}

// Delete deletes a workout
func (uc *Workout) Delete(id int) error {
	err := uc.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
