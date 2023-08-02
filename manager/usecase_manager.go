package manager

import (
	"bank-api/usecase"
	"sync"
)

type UsecaseManager interface {
	GetUserUsecase() usecase.UserUseCase
	GetLoginUsecase() usecase.LoginUseCase
	GetMerchantUsecase() usecase.MerchantUseCase
	GetPaymentUsecase() usecase.PaymentUseCase
}

type usecaseManager struct {
	repoManager RepoManager

	usrUsecase    usecase.UserUseCase
	lgUsecase     usecase.LoginUseCase
	mctUsecase 	  usecase.MerchantUseCase
	pyUsecase 	  usecase.PaymentUseCase
}

var onceLoadUserUsecase sync.Once
var onceLoadLoginUsecase sync.Once
var onceLoadMerchantUsecase sync.Once
var onceLoadPaymentUsecase sync.Once

func (um *usecaseManager) GetUserUsecase() usecase.UserUseCase {
	onceLoadUserUsecase.Do(func() {
		um.usrUsecase = usecase.NewUserUseCase(um.repoManager.GetUserRepo())
	})
	return um.usrUsecase
}

func (um *usecaseManager) GetLoginUsecase() usecase.LoginUseCase {
	onceLoadLoginUsecase.Do(func() {
		um.lgUsecase = usecase.NewLoginUseCase(um.repoManager.GetUserRepo())
	})
	return um.lgUsecase
}

func (um *usecaseManager) GetMerchantUsecase() usecase.MerchantUseCase {
	onceLoadMerchantUsecase.Do(func() {
		um.mctUsecase = usecase.NewMerchantUseCase(um.repoManager.GetMerchantRepo())
	})
	return um.mctUsecase
}

func (um *usecaseManager) GetPaymentUsecase() usecase.PaymentUseCase {
	onceLoadPaymentUsecase.Do(func() {
		um.pyUsecase = usecase.NewPaymentUseCase(um.repoManager.GetPaymentRepo())
	})
	return um.pyUsecase
}

func NewUsecaseManager(repoManager RepoManager) UsecaseManager {
	return &usecaseManager{
		repoManager: repoManager,
	}
}

