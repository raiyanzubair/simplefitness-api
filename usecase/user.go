package usecase

import (
	"simplefitnessApi/model"
	"simplefitnessApi/repository"
)

//User usecase interface
type User struct {
	repo *repository.User
}

// NewUserUsecase creates a new usecase interface for user
func NewUserUsecase(repo *repository.User) *User {
	return &User{repo}
}

// GetAll returns all user
func (uc *User) GetAll() ([]*model.User, error) {
	slice, err := uc.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return slice, nil
}

// GetByID returns a specific User
func (uc *User) GetByID(id int) (*model.User, error) {
	item, err := uc.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return item, nil
}

// GetByEmail returns a specific User given an email
func (uc *User) GetByEmail(email string) (*model.User, error) {
	item, err := uc.repo.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	return item, nil
}

// Create creates a new user
func (uc *User) Create(newUser *model.User) (*model.User, error) {
	created, err := uc.repo.Create(newUser)
	if err != nil {
		return nil, err
	}
	return created, nil
}

// Delete deletes a user
func (uc *User) Delete(id int) error {
	err := uc.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
