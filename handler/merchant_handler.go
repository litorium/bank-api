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

type MerchantHandler struct {
	mctUseCase usecase.MerchantUseCase
}

func (mctHandler *MerchantHandler) GetMerchantByMerchantname(ctx *gin.Context) {
	name := ctx.Param("name")
	mct, err := mctHandler.mctUseCase.GetMerchantByName(name)
	if err != nil {
		appError := &utils.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("MerchantHandler.GetMerchantByName() 1: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("MerchantHandler.GetMerchantByName() 2: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "An error occurred while fetching Merchant data",
			})
			return
		}
		return
	}
	if mct == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success":      false,
			"errorMessage": "Data not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    mct,
	})
}

func (mctHandler *MerchantHandler) AddMerchant(ctx *gin.Context) {
	mct := &model.MerchantModel{}
	err := ctx.ShouldBindJSON(&mct)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}

	err = mctHandler.mctUseCase.AddMerchant(mct)
	if err != nil {
		appError := &utils.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("MerchantHandler.InsertMerchant() 1: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("MerchantHandler.InsertMerchant() 2: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "An error occurred while saving Merchant data",
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Successfully added Merchant",
	})
}

func (mctHandler *MerchantHandler) UpdateMerchant(ctx *gin.Context) {
	mct := &model.MerchantModel{}
	err := ctx.ShouldBindJSON(&mct)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}

	err = mctHandler.mctUseCase.UpdateMerchant(mct)
	if err != nil {
		appError := &utils.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("MerchantHandler.EditMerchant() 1: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("MerchantHandler.EditMerchant() 2: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "An error occurred while saving Merchant data",
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Successfully updated Merchant",
	})
}

func (mctHandler *MerchantHandler) DeleteMerchant(ctx *gin.Context) {
	name := ctx.Param("name")
	
	if err := mctHandler.mctUseCase.DeleteMerchant(name); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Merchant"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Merchant deleted successfully",
	})
}

func NewMerchantHandler(srv *gin.Engine, mctUseCase usecase.MerchantUseCase) *MerchantHandler {
	mctHandler := &MerchantHandler{
		mctUseCase: mctUseCase,
	}

	// route
	srv.POST("/merchant", mctHandler.AddMerchant)
	srv.PUT("/merchant", mctHandler.UpdateMerchant)
	srv.GET("/merchant/:name", mctHandler.GetMerchantByMerchantname)
	srv.DELETE("/merchant/:name", mctHandler.DeleteMerchant)
	return mctHandler
}
