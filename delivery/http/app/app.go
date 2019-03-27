package app

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jmoiron/sqlx"
	"net/http"
	"os"
	// "simplefitnessApi/delivery/http/auth"
	"simplefitnessApi/delivery/http/handler"
	"simplefitnessApi/repository"
	"simplefitnessApi/usecase"
)

// App is the simplefitness api
type App struct {
	Router *chi.Mux
}

// New creates a instance of our app and sets up our routes/middleware
func New() (*App, error) {
	fmt.Println("Creating new server")

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	// Initialise our postgres db
	connStr := os.Getenv("DATABASE_URL")
	env := os.Getenv("GO_ENV")
	if connStr == "" || env == "development" {
		fmt.Println("We in dev")
		connStr = "user=raiyanzubair dbname=simplefitness sslmode=disable"
	}

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Init Repository Interfaces
	userRepo := repository.NewUserRepo(db)
	exerciseRepo := repository.NewExerciseRepo(db)
	exerciseTypeRepo := repository.NewExerciseTypeRepo(db)
	workoutRepo := repository.NewWorkoutRepo(db)
	workoutExerciseRepo := repository.NewWorkoutExerciseRepo(db)
	workoutExerciseSetRepo := repository.NewWorkoutExerciseSetRepo(db)

	// Init Usecase Interfaces
	userUsecase := usecase.NewUserUsecase(userRepo)
	exerciseUsecase := usecase.NewExerciseUsecase(exerciseRepo)
	exerciseTypeUsecase := usecase.NewExerciseTypeUsecase(exerciseTypeRepo)
	workoutUsecase := usecase.NewWorkoutUsecase(workoutRepo, workoutExerciseRepo, workoutExerciseSetRepo)
	workoutExerciseUsecase := usecase.NewWorkoutExerciseUsecase(workoutExerciseRepo, workoutExerciseSetRepo)

	// Init Handler Interfaces
	authHandler := handler.NewAuthHandler(userUsecase)
	exerciseHandler := handler.NewExerciseHandler(exerciseUsecase)
	exerciseTypeHandler := handler.NewExerciseTypeHandler(exerciseTypeUsecase)
	workoutHandler := handler.NewWorkoutHandler(workoutUsecase)
	workoutExerciseHandler := handler.NewWorkoutExerciseHandler(workoutExerciseUsecase)

	r.Route("/user", func(r chi.Router) {
		r.Post("/signin", authHandler.HandleSignIn)
		r.Post("/signup", authHandler.HandleSignUp)

		r.Route("/{userID}", func(r chi.Router) {
			r.Get("/", authHandler.GetUser)
			r.Get("/workout", workoutHandler.GetUserWorkouts)
		})
	})

	r.Route("/exercise", func(r chi.Router) {
		r.Get("/", exerciseHandler.GetAll)

		r.Route("/{exerciseID}", func(r chi.Router) {
			r.Get("/", exerciseHandler.GetByID)
		})
	})

	r.Route("/exercise_type", func(r chi.Router) {
		r.Get("/", exerciseTypeHandler.GetAll)

		r.Route("/{typeID}", func(r chi.Router) {
			r.Get("/", exerciseTypeHandler.GetByID)
		})
	})

	r.Route("/workout", func(r chi.Router) {
		// r.Use(auth.ValidateJWTMiddleware)

		r.Get("/", workoutHandler.GetAll)
		r.Post("/", workoutHandler.Create)

		r.Route("/{workoutID}", func(r chi.Router) {
			r.Get("/", workoutHandler.GetByID)
			r.Delete("/", workoutHandler.Delete)
		})
	})

	r.Route("/workout_exercise", func(r chi.Router) {
		r.Get("/", workoutExerciseHandler.GetAll)

		r.Route("/{workoutExerciseID}", func(r chi.Router) {
			r.Get("/", workoutExerciseHandler.GetByID)
		})
	})

	app := App{r}
	return &app, nil
}
