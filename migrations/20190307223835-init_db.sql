-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  email VARCHAR(255) UNIQUE NOT NULL,
  hashed_password VARCHAR(255) NOT NULL,
  created_at TIMESTAMPTZ,
  updated_at TIMESTAMPTZ
);
CREATE TABLE IF NOT EXISTS workouts (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255) UNIQUE NOT NULL,
  day INTEGER CONSTRAINT day_enum CHECK (day >= 0 AND day < 7),
  creator INTEGER NOT NULL REFERENCES users(id)
);
CREATE TABLE IF NOT EXISTS exercise_types (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL
);
CREATE TABLE IF NOT EXISTS measurement_unit (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL
);
CREATE TABLE IF NOT EXISTS exercises (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  exercise_type INTEGER NOT NULL REFERENCES exercise_types(id)
);
CREATE TABLE IF NOT EXISTS workout_exercises (
  id SERIAL PRIMARY KEY,
  workout INTEGER NOT NULL REFERENCES workouts(id),
  exercise INTEGER NOT NULL REFERENCES exercises(id)
);
CREATE TABLE IF NOT EXISTS workout_exercise_sets (
  id SERIAL PRIMARY KEY,
  workout_exercise INTEGER NOT NULL REFERENCES workout_exercises(id),
  measurement_unit INTEGER NOT NULL REFERENCES measurement_unit(id),
  reps INTEGER,
  resistance FLOAT,
  duration INTEGER
);

-- +migrate Down
DROP TABLE IF EXISTS "users" CASCADE;
DROP TABLE IF EXISTS "workouts" CASCADE;
DROP TABLE IF EXISTS "exercise_types" CASCADE;
DROP TABLE IF EXISTS "measurement_unit" CASCADE;
DROP TABLE IF EXISTS "exercises" CASCADE;
DROP TABLE IF EXISTS "workout_exercises" CASCADE;
DROP TABLE IF EXISTS "workout_exercise_sets" CASCADE;
