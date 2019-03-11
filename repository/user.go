package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"simplefitnessApi/model"
)

//User repo interface
type User struct {
	db *sqlx.DB
}

// NewUserRepo creates a new repo interface for users
func NewUserRepo(db *sqlx.DB) *User {
	return &User{db}
}

// GetAll returns all users
func (repo *User) GetAll() ([]*model.User, error) {
	arr := []*model.User{}
	err := repo.db.Select(&arr, "SELECT id, email, hashed_password FROM users")
	if err != nil {
		return nil, err
	}
	return arr, nil
}

// GetByID returns a specific User
func (repo *User) GetByID(id int) (*model.User, error) {
	item := model.User{}
	err := repo.db.Get(&item, "SELECT id, email, hashed_password FROM users WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

// GetByEmail returns a specific User given an email
func (repo *User) GetByEmail(email string) (*model.User, error) {
	item := model.User{}
	err := repo.db.Get(&item, "SELECT id, email, hashed_password FROM users WHERE email=$1", email)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

// Create will create insert a user into the db
func (repo *User) Create(newUser *model.User) (*model.User, error) {
	result := repo.db.QueryRow("INSERT INTO users (email, hashed_password, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id",
		newUser.Email,
		newUser.HashedPassword,
		newUser.CreatedAt,
		newUser.UpdatedAt,
	)
	var id int
	result.Scan(&id)

	user, err := repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Delete removes a user from the db
func (repo *User) Delete(id int) error {
	result, err := repo.db.Exec("DELETE FROM \"user\" WHERE id = $1", id)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}
