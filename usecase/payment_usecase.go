package usecase

import (
	"bank-api/model"
	"bank-api/repository"

	"github.com/gin-gonic/gin"
)

type PaymentUseCase interface {
	AddPayment(*model.PaymentModel, *gin.Context) error
	GetPaymentByUserId(string) ([]model.PaymentModel, error)
}

type paymentUseCaseImpl struct {
	pyRepo repository.PaymentRepository
}

func (pyUseCase *paymentUseCaseImpl) AddPayment(py *model.PaymentModel, ctx *gin.Context) error {
	return pyUseCase.pyRepo.AddPayment(py, ctx)
}

func (pyUseCase *paymentUseCaseImpl) GetPaymentByUserId(userId string) ([]model.PaymentModel, error) {
	return pyUseCase.pyRepo.GetPaymentByUserId(userId)
}

func NewPaymentUseCase(pyRepo repository.PaymentRepository) PaymentUseCase {
	return &paymentUseCaseImpl{
		pyRepo: pyRepo,
	}
}
