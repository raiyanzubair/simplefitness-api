package handler

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
	"simplefitnessApi/model"
	"simplefitnessApi/usecase"
	"strconv"
)

//Exercise handler interface
type Exercise struct {
	uc *usecase.Exercise
}

// NewExerciseHandler returns a new handler for exercise requests
func NewExerciseHandler(uc *usecase.Exercise) *Exercise {
	return &Exercise{uc}
}

func (e *Exercise) sliceToMap(slice []*model.Exercise) map[string]*model.Exercise {
	dict := make(map[string]*model.Exercise)
	for _, v := range slice {
		dict[strconv.Itoa(v.ID)] = v
	}
	return dict
}

// GetAll handles the route for getting all exercises
func (e *Exercise) GetAll(w http.ResponseWriter, r *http.Request) {
	slice, err := e.uc.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	dict := e.sliceToMap(slice)

	json.NewEncoder(w).Encode(dict)
}

// GetByID handles the route for getting a specific exercise given an ID
func (e *Exercise) GetByID(w http.ResponseWriter, r *http.Request) {
	exerciseID, err := strconv.Atoi(chi.URLParam(r, "exerciseID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	exercise, err := e.uc.GetByID(exerciseID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(exercise)
}
