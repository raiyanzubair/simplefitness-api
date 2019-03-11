package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
	"simplefitnessApi/model"
	"simplefitnessApi/usecase"
	"strconv"
)

//Workout handler interface
type Workout struct {
	uc *usecase.Workout
}

// NewWorkoutHandler returns a new handler for exercise requests
func NewWorkoutHandler(uc *usecase.Workout) *Workout {
	return &Workout{uc}
}

func (wo *Workout) sliceToMap(arr []*model.Workout) map[string]*model.Workout {
	dict := make(map[string]*model.Workout)
	for _, v := range arr {
		dict[strconv.Itoa(v.ID)] = v
	}
	return dict
}

// GetAll handles the route for getting all workouts
func (wo *Workout) GetAll(w http.ResponseWriter, r *http.Request) {
	arr, _ := wo.uc.GetAll()
	dict := wo.sliceToMap(arr)
	json.NewEncoder(w).Encode(dict)
}

// GetByID handles the route for getting a specific workout given an ID
func (wo *Workout) GetByID(w http.ResponseWriter, r *http.Request) {
	workoutID, err := strconv.Atoi(chi.URLParam(r, "workoutID"))
	if err != nil {
		http.Error(w, "Sorry something went wrong!", http.StatusInternalServerError)
		return
	}
	exercise, err := wo.uc.GetByID(workoutID)
	if err != nil {
		http.Error(w, "Sorry something went wrong!", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(exercise)
}

// Create handles the route for creating a workout
func (wo *Workout) Create(w http.ResponseWriter, r *http.Request) {
	newWorkout := model.Workout{}
	json.NewDecoder(r.Body).Decode(&newWorkout)
	created, err := wo.uc.Create(&newWorkout)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(created)
}

// Delete handles the route for creating a workout
func (wo *Workout) Delete(w http.ResponseWriter, r *http.Request) {
	workoutID, err := strconv.Atoi(chi.URLParam(r, "workoutID"))
	if err != nil {
		http.Error(w, "Sorry something went wrong!", http.StatusInternalServerError)
		return
	}
	err = wo.uc.Delete(workoutID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Workout deleted")
}
