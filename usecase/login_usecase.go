package usecase

import (
	"bank-api/model"
	"bank-api/repository"
	"bank-api/utils"
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginUseCase interface {
	Login(usr *model.LoginModel, ctx *gin.Context) (*model.UserModel, error)
	Logout(ctx *gin.Context)
}

type loginUsecase struct {
	usrRepo repository.UserRepository
}

func (loginUsecase *loginUsecase) Login(usr *model.LoginModel, ctx *gin.Context) (*model.UserModel, error) {
	// Login session
	session := sessions.Default(ctx)
	existSession := session.Get("Username")
	if existSession != nil {
		return nil, &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: fmt.Sprintf("You are already logged in as %v", existSession),
		}
	}

	existData, err := loginUsecase.usrRepo.GetUserByUsername(usr.Username)
	if err != nil {
		return nil, fmt.Errorf("loginUsecase.GetUserByName(): %w", err)
	}
	if existData == nil {
		return nil, &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: "Username is not registered",
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(existData.Password), []byte(usr.Password))
	if err != nil {
		return nil, &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: "Password does not match",
		}
	}


	// Login session
	session.Set("Username", existData.UserName)
	session.Save()

	existData.Password = ""
	return existData, nil
}

func (loginUsecase *loginUsecase) Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	session.Save()
}

func NewLoginUseCase(usrRepo repository.UserRepository) LoginUseCase {
	return &loginUsecase{
		usrRepo: usrRepo,
	}
}

