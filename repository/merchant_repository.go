package repository

import (
	"bank-api/model"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type MerchantRepository interface {
	GetMerchantByName(string) (*model.MerchantModel, error)
	UpdateMerchant(*model.MerchantModel) error
	DeleteMerchant(*model.MerchantModel) error
	AddMerchant(*model.MerchantModel) error
}

type merchantRepositoryImpl struct {
	merchants []model.MerchantModel
}

func (m *merchantRepositoryImpl) GetMerchantByName(name string) (*model.MerchantModel,error) {
	for _, mct := range m.merchants {
		if mct.Name == name {
			return &mct, nil
		}
	}
	return nil, errors.New("merchant not found")
}

func (m *merchantRepositoryImpl) UpdateMerchant(merchant *model.MerchantModel) error {
	// Check if the user exists in the slice
	found := false
	for i, mct := range m.merchants {
		if mct.Id == merchant.Id {
			m.merchants[i] = *merchant
			found = true
			break
		}
	}

	// If the merchant is not found, return an error
	if !found {
		return fmt.Errorf("merchant with ID %s not found", merchant.Id)
	}

	// Open the JSON file
	file, err := os.OpenFile("data/merchants.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(m.merchants)
	if err != nil {
		return err
	}

	return nil
}

func (m *merchantRepositoryImpl) AddMerchant(merchant *model.MerchantModel) error {
	// Check if the user exists in the slice
	found := false
	for i, mct := range m.merchants {
		if mct.Id == merchant.Id {
			m.merchants[i] = *merchant
			found = true
			break
		}
	}
	if !found {
		m.merchants = append(m.merchants, *merchant)
	}

	// Open the JSON file
	file, err := os.OpenFile("data/merchants.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(m.merchants)
	if err != nil {
		return err
	}

	return nil
}

func (m *merchantRepositoryImpl) DeleteMerchant(merchant *model.MerchantModel) error {
	// Remove the user from the slice
	for i, mct := range m.merchants {
		if mct.Id == merchant.Id {
			m.merchants = append(m.merchants[:i], m.merchants[i+1:]...)
			break
		}
	}

	// Open the JSON file
	file, err := os.OpenFile("data/merchants.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	// Encode the users slice back into the file
	err = json.NewEncoder(file).Encode(m.merchants)
	if err != nil {
		return err
	}

	return nil
}

func NewMerchantRepository() MerchantRepository {
	repo := &merchantRepositoryImpl{}

	// Open the JSON file
	file, err := os.Open("data/merchants.json")
	if err != nil {
		return nil
	}
	defer file.Close()

	// Decode the file into the users slice
	err = json.NewDecoder(file).Decode(&repo.merchants)
	if err != nil {
		return nil
	}

	return repo
}

