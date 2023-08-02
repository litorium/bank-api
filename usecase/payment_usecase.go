package usecase

import (
	"bank-api/model"
	"bank-api/repository"
)

type PaymentUseCase interface {
	AddPayment(*model.PaymentModel) error
	GetPaymentByUserId(string) ([]model.PaymentModel, error)
}

type paymentUseCaseImpl struct {
	pyRepo repository.PaymentRepository
}

func (pyUseCase *paymentUseCaseImpl) AddPayment(py *model.PaymentModel) error {
	return pyUseCase.pyRepo.AddPayment(py)
}

func (pyUseCase *paymentUseCaseImpl) GetPaymentByUserId(userId string) ([]model.PaymentModel, error) {
	return pyUseCase.pyRepo.GetPaymentByUserId(userId)
}

func NewPaymentUseCase(pyRepo repository.PaymentRepository) PaymentUseCase {
	return &paymentUseCaseImpl{
		pyRepo: pyRepo,
	}
}
