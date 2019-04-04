package handler

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
	"simplefitnessApi/model"
	"simplefitnessApi/usecase"
	"strconv"
)

//ExerciseType handler interface
type ExerciseType struct {
	uc *usecase.ExerciseType
}

// NewExerciseTypeHandler returns a new handler for exercise type requests
func NewExerciseTypeHandler(uc *usecase.ExerciseType) *ExerciseType {
	return &ExerciseType{uc}
}

func (et *ExerciseType) sliceToMap(slice []*model.ExerciseType) map[string]*model.ExerciseType {
	dict := make(map[string]*model.ExerciseType)
	for _, v := range slice {
		dict[strconv.Itoa(v.ID)] = v
	}
	return dict
}

// GetAll handles the route for getting all exercise types
func (et *ExerciseType) GetAll(w http.ResponseWriter, r *http.Request) {
	slice, _ := et.uc.GetAll()
	dict := et.sliceToMap(slice)
	json.NewEncoder(w).Encode(dict)
}

// GetByID handles the route for getting a specific exercise type given an ID
func (et *ExerciseType) GetByID(w http.ResponseWriter, r *http.Request) {
	typeID, err := strconv.Atoi(chi.URLParam(r, "typeID"))
	if err != nil {
		http.Error(w, "Sorry something went wrong!", http.StatusInternalServerError)
		return
	}
	eType, err := et.uc.GetByID(typeID)
	if err != nil {
		http.Error(w, "Sorry something went wrong!", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(eType)
}
