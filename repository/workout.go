package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	// "os"
	"simplefitnessApi/model"
)

//Workout repo interface
type Workout struct {
	db *sqlx.DB
}

// NewWorkoutRepo creates a new repo interface for workouts
func NewWorkoutRepo(db *sqlx.DB) *Workout {
	return &Workout{db}
}

// GetAll returns all Workouts
func (repo *Workout) GetAll() ([]*model.Workout, error) {
	workouts := []*model.Workout{}
	query := `
		SELECT * FROM workouts
	`
	err := repo.db.Select(&workouts, query)
	if err != nil {
		return nil, err
	}
	return workouts, nil
}

// GetByID returns a specific Workout and its associated exercises
func (repo *Workout) GetByID(id int) (*model.Workout, error) {
	//Get workouts
	workout := model.Workout{}
	query := `
		SELECT * FROM workouts
		WHERE id = $1
	`
	err := repo.db.Get(&workout, query, id)
	if err != nil {
		return nil, err
	}
	return &workout, nil
}

// GetByCreator returns Workouts for a specific user
func (repo *Workout) GetByCreator(creatorID int) ([]*model.Workout, error) {
	arr := []*model.Workout{}
	query := `
		SELECT * FROM workouts 
		WHERE creator = $1
	`
	err := repo.db.Select(&arr, query, creatorID)
	if err != nil {
		fmt.Printf(err.Error())
		return nil, err
	}
	return arr, nil
}

// Create will create insert a new workout into the db and return it
func (repo *Workout) Create(newWorkout *model.Workout) (*model.Workout, error) {
	query := `
		INSERT INTO workouts (title, day, creator) 
		VALUES ($1, $2, $3) RETURNING id
	`
	result := repo.db.QueryRow(query,
		newWorkout.Title,
		newWorkout.Day,
		newWorkout.CreatorID,
	)
	var id int
	err := result.Scan(&id)
	if err != nil {
		return nil, err
	}
	workout, err := repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return workout, nil
}

// Delete removes a workout from the db
func (repo *Workout) Delete(id int) error {
	query := `
		DELETE FROM workouts
		WHERE id = $1
	`
	_, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
