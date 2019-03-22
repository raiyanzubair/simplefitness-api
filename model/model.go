package model

// User Model
type User struct {
	ID             int    `db:"id" json:"id"`
	Email          string `db:"email" json:"email"`
	Password       string `json:"-"`
	HashedPassword string `db:"hashed_password" json:"-"`
	CreatedAt      string `db:"created_at" json:"-"`
	UpdatedAt      string `db:"updated_at" json:"-"`
}

// Workout Model
type Workout struct {
	ID        int               `db:"id" json:"id"`
	Title     string            `db:"title" json:"title"`
	Day       int               `db:"day" json:"day"`
	CreatorID int               `db:"creator" json:"creator"`
	Exercises []WorkoutExercise `json:"exercises"`
}

// Exercise Model
type Exercise struct {
	ID           int    `db:"id" json:"id"`
	Title        string `db:"title" json:"title"`
	ExerciseType `db:"exercise_type" json:"exercise_type"`
}

// ExerciseType Model
type ExerciseType struct {
	ID    int    `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
}

// WorkoutExercise Model
type WorkoutExercise struct {
	ID        int `db:"id" json:"id"`
	WorkoutID int `db:"workout" json:"workout"`
	Exercise  `db:"exercise" json:"exercise"`
	Sets      []WorkoutExerciseSet `json:"sets"`
}

// WorkoutExerciseSet Model
type WorkoutExerciseSet struct {
	ID                int     `db:"id" json:"id"`
	WorkoutExerciseID int     `db:"workout_exercise" json:"workout_exercise"`
	Reps              int     `db:"reps" json:"reps"`
	Resistance        float64 `db:"resistance" json:"resistance"`
	MeasurementUnit   `db:"measurement_unit" json:"measurement_unit"`
	Duration          int `db:"duration" json:"duration"`
}

// MeasurementUnit Model
type MeasurementUnit struct {
	ID    int    `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
}
