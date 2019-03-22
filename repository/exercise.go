package repository

import (
	"github.com/jmoiron/sqlx"
	"log"
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
	query := `
		SELECT e.id, e.title, et.id as "exercise_type.id", et.title as "exercise_type.title"
		FROM exercises as e 
		JOIN exercise_types as et 
		ON e.exercise_type = et.id	
	`
	err := repo.db.Select(&arr, query)
	if err != nil {
		return nil, err
	}
	return arr, nil
}

// GetByID returns a specific Exercise
func (repo *Exercise) GetByID(id int) (*model.Exercise, error) {
	item := model.Exercise{}
	query := `
		SELECT e.id, e.title, et.id as "exercise_type.id", et.title as "exercise_type.title"
		FROM exercises as e 
		JOIN exercise_types as et 
		ON e.exercise_type = et.id
		WHERE e.id = $1
	`
	err := repo.db.Get(&item, query, id)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

// GetByName returns a specific Exercise
func (repo *Exercise) GetByName(name string) ([]*model.Exercise, error) {
	arr := []*model.Exercise{}
	query := `
		SELECT e.id, e.title, et.id as "exercise_type.id", et.title as "exercise_type.title"
		FROM exercises as e 
		JOIN exercise_types as et 
		ON e.exercise_type = et.id
		WHERE e.title LIKE '%' || $1 || '%'
	`
	log.Println("hello")
	log.Println(name)
	err := repo.db.Select(&arr, query, name)
	if err != nil {
		return nil, err
	}
	return arr, nil
}
