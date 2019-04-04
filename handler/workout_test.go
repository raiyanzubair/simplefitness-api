package handler_test

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"simplefitnessApi/handler"
	"simplefitnessApi/model"

	//"simplefitnessApi/model"
	"simplefitnessApi/repository"
	"simplefitnessApi/usecase"
	"testing"
)

func initTestWorkoutRouter() *chi.Mux {
	db := repository.PrepTestDB()
	wRepo := repository.NewWorkoutRepo(db)
	weRepo := repository.NewWorkoutExerciseRepo(db)
	wesRepo := repository.NewWorkoutExerciseSetRepo(db)

	workout := usecase.NewWorkoutUsecase(wRepo, weRepo, wesRepo)

	h := handler.NewWorkoutHandler(workout)

	router := chi.NewRouter()
	router.Get("/workout", h.GetAll)
	router.Get("/workout/{workoutID}", h.GetByID)
	router.Get("/user/{userID}/workout", h.GetUserWorkouts)
	router.Post("/workout}", h.Create)
	router.Delete("/workout/{workoutID}", h.Delete)

	return router
}

func TestWorkout_GetAll(t *testing.T) {
	req :=  httptest.NewRequest("GET", "/workout", nil)
	rr := httptest.NewRecorder()

	router := initTestWorkoutRouter()
	router.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code, "Should return 200")
}

func TestWorkout_GetByID(t *testing.T) {
	tests := []struct{
		id int
		status int
	}{
		{id: 1, status: 200},
		{id: 2, status: 200},
		{id: 3, status: 500},
		{id: 4, status: 500},
	}
	router := initTestWorkoutRouter()
	for _, test := range tests {
		req :=  httptest.NewRequest("GET", fmt.Sprintf("/workout/%d", test.id), nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
		assert.Equal(t, test.status, rr.Code, "Expected %d, got %d", test.status, rr.Code)
	}
}

func TestWorkout_GetUserWorkouts(t *testing.T) {
	tests := []struct{
		user int
		length int
	}{
		{user: 1, length: 1},
		{user: 2, length: 1},
		{user: 3, length: 0},
	}
	router := initTestWorkoutRouter()
	for _, test := range tests {
		req :=  httptest.NewRequest("GET", fmt.Sprintf("/user/%d/workout", test.user), nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		var workouts []model.Workout
		err := json.NewDecoder(rr.Body).Decode(&workouts)

		assert.NoError(t, err, "Expected no error in unmarshalling json")
		assert.Len(t,workouts, test.length, "User %d should have %d workouts, got %d", test.user, test.length, len(workouts))
	}
}

func TestWorkout_Create(t *testing.T) {
}

func TestWorkout_Delete(t *testing.T) {
}

