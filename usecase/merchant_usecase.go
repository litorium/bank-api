package usecase

import (
	"bank-api/model"
	"bank-api/repository"
	"bank-api/utils"
	"fmt"

	"github.com/google/uuid"
)

type MerchantUseCase interface {
	GetMerchantByName(string) (*model.MerchantModel, error)
	AddMerchant(*model.MerchantModel) error
	UpdateMerchant(*model.MerchantModel) error
	DeleteMerchant(string) error
}

type merchantUseCaseImpl struct {
	mctRepo repository.MerchantRepository
}

func (mctUseCase *merchantUseCaseImpl) GetMerchantByName(name string) (*model.MerchantModel, error) {
	return mctUseCase.mctRepo.GetMerchantByName(name)
}

func (mctUseCase *merchantUseCaseImpl) AddMerchant(mct *model.MerchantModel) error {
	merchant, _ := mctUseCase.mctRepo.GetMerchantByName(mct.Name)
	if merchant != nil {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: fmt.Sprintf("User data with the name %v already exists", mct.Name),
		}
	}
	mct.Id = uuid.New().String()
	return mctUseCase.mctRepo.AddMerchant(mct)
}

func (mctUseCase *merchantUseCaseImpl) UpdateMerchant(mct *model.MerchantModel) error {
	merchant, _ := mctUseCase.mctRepo.GetMerchantByName(mct.Name)
	if merchant != nil {
		return &utils.AppError{
			ErrorCode:    1,
			ErrorMessage: fmt.Sprintf("User data with the name %v already exists", mct.Name),
		}
	}
	return mctUseCase.mctRepo.UpdateMerchant(mct)
}

func (mctUseCase *merchantUseCaseImpl) DeleteMerchant(name string) error {
	user , err := mctUseCase.mctRepo.GetMerchantByName(name)
	if user == nil {
		return err
	}
	return mctUseCase.mctRepo.DeleteMerchant(user)
}

func NewMerchantUseCase(mctRepo repository.MerchantRepository) MerchantUseCase {
	return &merchantUseCaseImpl{
		mctRepo: mctRepo,
	}
}