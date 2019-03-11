package repository

import (
	"github.com/jmoiron/sqlx"
	// "os"
	"simplefitnessApi/model"
)

//Exercise repo interface
type Exercise struct {
	db *sqlx.DB
}

// NewExerciseRepo creates a new repo interface for exercises
func NewExerciseRepo(db *sqlx.DB) *Exercise {
	return &Exercise{db}
}

// GetAll returns all exercises
func (repo *Exercise) GetAll() ([]*model.Exercise, error) {
	arr := []*model.Exercise{}
	err := repo.db.Select(&arr, "SELECT id, title, exercise_type FROM exercises")
	if err != nil {
		return nil, err
	}
	return arr, nil
}

// GetByID returns a specific Exercise
func (repo *Exercise) GetByID(id int) (*model.Exercise, error) {
	item := model.Exercise{}
	err := repo.db.Get(&item, "SELECT id, title, exercise_type FROM exercises WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	return &item, nil
}
