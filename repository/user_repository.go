package repository

import (
	"bank-api/model"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type UserRepository interface {
	GetUserByUsername(username string) (*model.UserModel, error)
	UpdateUser(*model.UserModel) error
	DeleteUser(*model.UserModel) error
	AddUser(*model.UserModel) error
}

type userRepositoryImpl struct {
	users []model.UserModel
}

func (r *userRepositoryImpl) GetUserByUsername(username string) (*model.UserModel,error) {
	for _, usr := range r.users {
		if usr.UserName == username {
			return &usr,nil
		}
	}
	return nil, errors.New("user not found")
}

func (r *userRepositoryImpl) UpdateUser(user *model.UserModel) error {
	// Check if the user exists in the slice
	found := false
	for i, mct := range r.users {
		if mct.Id == user.Id {
			r.users[i] = *user
			found = true
			break
		}
	}

	// If the user is not found, return an error
	if !found {
		return fmt.Errorf("user with ID %s not found", user.Id)
	}

	// Open the JSON file
	file, err := os.OpenFile("data/users.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(r.users)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepositoryImpl) DeleteUser(user *model.UserModel) error {
	// Remove the user from the slice
	for i, u := range r.users {
		if u.Id == user.Id {
			r.users = append(r.users[:i], r.users[i+1:]...)
			break
		}
	}

	// Open the JSON file
	file, err := os.OpenFile("data/users.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	// Encode the users slice back into the file
	err = json.NewEncoder(file).Encode(r.users)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepositoryImpl) AddUser(user *model.UserModel) error {
	// Check if the user exists in the slice
	found := false
	for i, u := range r.users {
		if u.Id == user.Id {
			r.users[i] = *user
			found = true
			break
		}
	}
	if !found {
		r.users = append(r.users, *user)
	}

	// Open the JSON file
	file, err := os.OpenFile("data/users.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(r.users)
	if err != nil {
		return err
	}

	return nil
}

func NewUserRepository() UserRepository {
	repo := &userRepositoryImpl{}

	// Open the JSON file
	file, err := os.Open("data/users.json")
	if err != nil {
		return nil
	}
	defer file.Close()

	// Decode the file into the users slice
	err = json.NewDecoder(file).Decode(&repo.users)
	if err != nil {
		return nil
	}

	return repo
}