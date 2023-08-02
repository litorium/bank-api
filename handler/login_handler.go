package handler

import (
	"bank-api/model"
	"bank-api/usecase"
	"bank-api/utils"
	"errors"
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	lgUsecase usecase.LoginUseCase
}

func (lgHandler LoginHandler) Login(ctx *gin.Context) {
	loginReq := &model.LoginModel{}
	err := ctx.ShouldBindJSON(&loginReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	if loginReq.Username == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Name cannot be empty",
		})
		return
	}
	if loginReq.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Password cannot be empty",
		})
		return
	}

	usr, err := lgHandler.lgUsecase.Login(loginReq, ctx)

	if err != nil {
		appError := &utils.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("LoginHandler.GetUserByName() 1: %v", err.Error())
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("LoginHandler.GetUserByName() 2: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "An error occurred during login",
				"errorCode": err.Error(),
			})
		}
		return
	}
	if usr == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Name is not registered",
		})
		return
	}

	tokenJwt, err := utils.GenerateToken(loginReq.Username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "Invalid Token",
		})
		return
	}

	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    tokenJwt,
		HttpOnly: true,
	})
	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": tokenJwt,
	})
}

func (lgHandler LoginHandler) Logout(ctx *gin.Context) {
	lgHandler.lgUsecase.Logout(ctx)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Logout",
	})
}

func NewLoginHandler(srv *gin.Engine, lgUsecase usecase.LoginUseCase) *LoginHandler {
	lgHandler := &LoginHandler{
		lgUsecase: lgUsecase,
	}

	srv.POST("/login", lgHandler.Login)
	srv.POST("/logout", lgHandler.Logout)

	return lgHandler
}