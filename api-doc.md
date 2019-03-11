# API Documentation (Work in Progress)

## Resources
- User
- Exercise
- Exercise Type
- Workout
- Workout Exercise

## Schema

```
User: {
  id: PK
  email: string
  hashed_pass: string
  birthday: string
}
```
```
Exercise: {
  id: PK
  title: string
  exercise_type: FK ExerciseType
}
```
```
ExerciseType: {
  id: PK
  title: string
}
```
```
Workout: {
  id: PK
  title: 
  day: string/enum/??
  creator: FK User
}
```
```
WorkoutExercise: {
  id: PK
  sets: int
  reps: int
  workout: FK Workout
  exercise: FK Exercise
}
```

## User Stories

- User can create a workout
  - User can add an exercise to a workout
  - User can adjust sets/reps for an exercise in a workout
- User can delete a workout
- User can list all exercises


## Routes
`/api/v1`

### Auth
`/auth/signin`
- POST login user

`/auth/signup`
- POST create user

### User
`/user/{id}` 
- GET specific user
v
`/user/{id}/workout`
- GET all workouts from a specific user

### Exercise
`/exercise` 
- GET all exercises  

`/exercise/{id}` 
- GET specific exercise

### Exercise Type
`/exercise_type` 
- GET all exercises types

### Workout
`/workout`
- GET all workouts
- POST create a workout 
```
data: [
  {
    exercise: (int)
    user: (int)
    sets: (int)
    reps: (int)
  },
  ...
]
```

`/workout/workout_exercise`
- GET all workouts and their exercises

`/workout/{id}`
- GET specific workout

`/workout/{id}/workout_exercise`
- GET specific workout and its exercises

