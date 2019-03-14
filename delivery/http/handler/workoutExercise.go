package handler

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
	"simplefitnessApi/model"
	"simplefitnessApi/usecase"
	"strconv"
)

//WorkoutExercise handler interface
type WorkoutExercise struct {
	uc *usecase.WorkoutExercise
}

// NewWorkoutExerciseHandler returns a new handler for exercise requests
func NewWorkoutExerciseHandler(uc *usecase.WorkoutExercise) *WorkoutExercise {
	return &WorkoutExercise{uc}
}

func (we *WorkoutExercise) sliceToMap(arr []*model.WorkoutExercise) map[string]*model.WorkoutExercise {
	dict := make(map[string]*model.WorkoutExercise)
	for _, v := range arr {
		dict[strconv.Itoa(v.ID)] = v
	}
	return dict
}

// GetAll handles the route for getting all workout exercises
func (we *WorkoutExercise) GetAll(w http.ResponseWriter, r *http.Request) {
	arr, err := we.uc.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	dict := we.sliceToMap(arr)
	json.NewEncoder(w).Encode(dict)
}

// GetByID handles the route for getting a specific workout exercise given an ID
func (we *WorkoutExercise) GetByID(w http.ResponseWriter, r *http.Request) {
	exerciseID, err := strconv.Atoi(chi.URLParam(r, "exerciseID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	exercise, err := we.uc.GetByID(exerciseID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(exercise)
}
