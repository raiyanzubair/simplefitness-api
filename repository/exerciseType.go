package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	// "os"
	"simplefitnessApi/model"
)

//ExerciseType repo interface
type ExerciseType struct {
	db *sqlx.DB
}

// NewExerciseTypeRepo creates a new repo interface for exercise types
func NewExerciseTypeRepo(db *sqlx.DB) *ExerciseType {
	return &ExerciseType{db}
}

// GetAll returns all ExerciseTypes
func (repo *ExerciseType) GetAll() ([]*model.ExerciseType, error) {
	slice := []*model.ExerciseType{}
	err := repo.db.Select(&slice, "SELECT id, title FROM exercise_types")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return slice, nil
}

// GetByID returns a specific ExerciseType
func (repo *ExerciseType) GetByID(id int) (*model.ExerciseType, error) {
	item := model.ExerciseType{}
	err := repo.db.Get(&item, "SELECT id, title FROM exercise_types WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &item, nil
}
