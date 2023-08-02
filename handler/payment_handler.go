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

type PaymentHandler struct {
	pyUseCase usecase.PaymentUseCase
}

func (pyHandler *PaymentHandler) GetPaymentByUserId(ctx *gin.Context) {
	id := ctx.Param("id")
	py, err := pyHandler.pyUseCase.GetPaymentByUserId(id)
	if err != nil {
		appError := &utils.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("PaymentHandler.GetPaymentByName() 1: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("PaymentHandler.GetPaymentByName() 2: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "An error occurred while fetching Payment data",
			})
			return
		}
		return
	}
	if py == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success":      false,
			"errorMessage": "Data not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    py,
	})
}

func (pyHandler *PaymentHandler) AddPayment(ctx *gin.Context) {
	py := &model.PaymentModel{}
	err := ctx.ShouldBindJSON(&py)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}

	err = pyHandler.pyUseCase.AddPayment(py)
	if err != nil {
		appError := &utils.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("PaymentHandler.InsertPayment() 1: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("PaymentHandler.InsertPayment() 2: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "An error occurred while saving Payment data",
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func NewPaymentHandler(srv *gin.Engine, pyUseCase usecase.PaymentUseCase) *PaymentHandler {
	pyHandler := &PaymentHandler{
		pyUseCase: pyUseCase,
	}

	// route
	srv.POST("/payment", middleware.RequireToken(), pyHandler.AddPayment)
	srv.GET("/payment/:id", middleware.RequireToken(), pyHandler.GetPaymentByUserId)
	return pyHandler
}