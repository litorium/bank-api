package usecase

import (
	"bank-api/model"
	"bank-api/repository"
	"bank-api/utils"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	AddUser(*model.UserModel) error
	GetUserByUsername(string) (*model.UserModel, error)
	UpdateUser(*model.UserModel) error
	DeleteUser(string) error
}

type userUseCaseImpl struct {
	usrRepo repository.UserRepository
}

func (usrUseCase *userUseCaseImpl) GetUserByUsername(username string) (*model.UserModel, error) {
	return usrUseCase.usrRepo.GetUserByUsername(username)
}

func (usrUseCase *userUseCaseImpl) AddUser(usr *model.UserModel) error {
	user,_ := usrUseCase.usrRepo.GetUserByUsername(usr.UserName)
	if user != nil {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: fmt.Sprintf("User data with the name %v already exists", usr.UserName),
		}
	}
	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
    }
	usr.Id = uuid.New().String()
	usr.Password = string(hashedPassword)
   return usrUseCase.usrRepo.AddUser(usr)
}

func (usrUseCase *userUseCaseImpl) UpdateUser(usr *model.UserModel) error {
	if usr.UserName == "" {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: "Name cannot be empty",
		}
	}
	if usr.Password == "" {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: "Password cannot be empty",
		}
	}

	existDataUsr, _ := usrUseCase.usrRepo.GetUserByUsername(usr.UserName)
	if existDataUsr != nil {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: fmt.Sprintf("User data with the username %v already exists", usr.UserName),
		}
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("userUsecaseImpl.GenerateFromPassword(): %w", err)
	}
	usr.Password = string(passHash)

	return usrUseCase.usrRepo.UpdateUser(usr)
}

func (usrUseCase *userUseCaseImpl) DeleteUser(username string) error {
	user , err := usrUseCase.usrRepo.GetUserByUsername(username)
	if user == nil {
		return err
	}
	return usrUseCase.usrRepo.DeleteUser(user)
}

func NewUserUseCase(usrRepo repository.UserRepository) UserUseCase {
	return &userUseCaseImpl{
		usrRepo: usrRepo,
	}
}


