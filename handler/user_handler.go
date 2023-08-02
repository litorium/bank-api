package handler

import (
	"bank-api/middleware"
	"bank-api/model"
	"bank-api/usecase"
	"bank-api/utils"
	"errors"
	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	usrUseCase usecase.UserUseCase
}

func (usrHandler UserHandler) GetUserByUsername(ctx *gin.Context) {
	username := ctx.Param("username")
	usr, err := usrHandler.usrUseCase.GetUserByUsername(username)
	if err != nil {
		appError := &utils.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("UserHandler.GetUserByName() 1: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("UserHandler.GetUserByName() 2: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "An error occurred while fetching user data",
			})
			return
		}
		return
	}
	if usr == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success":      false,
			"errorMessage": "Data not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    usr,
	})
}

func (usrHandler UserHandler) AddUser(ctx *gin.Context) {
	usr := &model.UserModel{}
	err := ctx.ShouldBindJSON(&usr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}

	err = usrHandler.usrUseCase.AddUser(usr)
	if err != nil {
		appError := &utils.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("UserHandler.InsertUser() 1: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("UserHandler.InsertUser() 2: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "An error occurred while saving user data",
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (usrHandler UserHandler) UpdateUser(ctx *gin.Context) {
	usr := &model.UserModel{}
	err := ctx.ShouldBindJSON(&usr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}

	err = usrHandler.usrUseCase.UpdateUser(usr)
	if err != nil {
		appError := &utils.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("UserHandler.EditUser() 1: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("UserHandler.EditUser() 2: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "An error occurred while saving user data",
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (usrHandler *UserHandler) DeleteUser(ctx *gin.Context) {
	username := ctx.Param("username")
	
	if err := usrHandler.usrUseCase.DeleteUser(username); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User deleted successfully",
	})
}

func NewUserHandler(srv *gin.Engine, usrUseCase usecase.UserUseCase) *UserHandler {
	usrHandler := &UserHandler{
		usrUseCase: usrUseCase,
	}

	// route
	srv.POST("/user", usrHandler.AddUser)
	srv.PUT("/user", middleware.RequireToken(), usrHandler.UpdateUser)
	srv.GET("/user/:username", middleware.RequireToken(), usrHandler.GetUserByUsername)
	srv.DELETE("/users/:username", middleware.RequireToken(), usrHandler.DeleteUser)
	return usrHandler
}
