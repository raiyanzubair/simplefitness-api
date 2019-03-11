package model

// User Model
type User struct {
	ID             int    `db:"id" json:"id"`
	Email          string `db:"email" json:"email"`
	Password       string `json:"password"`
	HashedPassword string `db:"hashed_password" json:"hashed_password"`
	CreatedAt      string `db:"created_at" json:"created_at"`
	UpdatedAt      string `db:"updated_at" json:"updated_at"`
}

// Workout Model
type Workout struct {
	ID        int     `db:"id" json:"id"`
	Title     *string `db:"title" json:"title"`
	Day       *string `db:"day" json:"day"`
	CreatorID int     `db:"creator" json:"creator"`
}

// Exercise Model
type Exercise struct {
	ID             int     `db:"id" json:"id"`
	Title          *string `db:"title" json:"title"`
	ExerciseTypeID int     `db:"exercise_type" json:"exercise_type"`
}

// ExerciseType Model
type ExerciseType struct {
	ID    int     `db:"id" json:"id"`
	Title *string `db:"title" json:"title"`
}

// WorkoutExercise Model
type WorkoutExercise struct {
	ID         int `db:"id" json:"id"`
	Sets       int `db:"sets" json:"sets"`
	Reps       int `db:"reps" json:"reps"`
	WorkoutID  int `db:"workout" json:"workout"`
	ExerciseID int `db:"exercise" json:"exercise"`
	// Exercises *Exercise
}
