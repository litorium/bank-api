package repository

import (
	"bank-api/model"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PaymentRepository interface {
	AddPayment(*model.PaymentModel, *gin.Context) error
	GetPaymentByUserId(string) ([]model.PaymentModel, error)
}

type paymentRepositoryImpl struct {
	payments  []model.PaymentModel
	merchants []model.MerchantModel
	users     []model.UserModel
}

func (p *paymentRepositoryImpl) GetPaymentByUserId(userId string) ([]model.PaymentModel, error) {
	// Validate userId
	if userId == "" {
		return nil, errors.New("userId cannot be empty")
	}

	// Check if userId exists in the users slice
	var userExists bool
	for _, user := range p.users {
		if user.Id == userId {
			userExists = true
			break
		}
	}

	if !userExists {
		return nil, fmt.Errorf("user with id %s not found", userId)
	}
	var userPayments []model.PaymentModel
	for _, py := range p.payments {
		if py.UserId == userId {
			userPayments = append(userPayments, py)
		}
	}

	if len(userPayments) == 0 {
		return nil, errors.New("no payments found for the user")
	}

	return userPayments, nil
}

func (p *paymentRepositoryImpl) AddPayment(paymentInput *model.PaymentModel, ctx *gin.Context) error {
	var user model.UserModel
	session := sessions.Default(ctx)
	existSession := session.Get("Id")
	userId, ok := existSession.(string)
	if !ok {
		return errors.New("session user ID not found")
	}

	for _, u := range p.users {
		if u.Id == userId {
			user = u
			break
		}
	}

	if user.Id == "" {
		return fmt.Errorf("user dengan id %s tidak ditemukan", userId)
	}

	var merchant model.MerchantModel
	for _, mct := range p.merchants {
		if mct.NoRek == paymentInput.MerchantNoRek {
			merchant = mct
			break
		}
	}

	if merchant.NoRek == "" {
		return fmt.Errorf("merchant dengan nomor rekening tersebut tidak ditemukan")
	}

	payment := model.PaymentModel{
		Id:           uuid.New().String(),
		UserId:       userId, // Mengisi UserId dengan session user ID
		MerchantNoRek: merchant.NoRek,
		Amount:       paymentInput.Amount,
		CreatedAt:    time.Now().UTC(),
	}

	p.payments = append(p.payments, payment)

	err := writeJSONFile("data/payments.json", p.payments)
	if err != nil {
		return err
	}

	return nil
}

func NewPaymentRepository() PaymentRepository {
	repo := &paymentRepositoryImpl{}

	// Open the JSON files
	usersFile, err := os.Open("data/users.json")
	if err != nil {
		return nil
	}
	defer usersFile.Close()

	merchantsFile, err := os.Open("data/merchants.json")
	if err != nil {
		return nil
	}
	defer merchantsFile.Close()

	// Decode the files into the respective slices
	err = json.NewDecoder(usersFile).Decode(&repo.users)
	if err != nil {
		return nil
	}

	err = json.NewDecoder(merchantsFile).Decode(&repo.merchants)
	if err != nil {
		return nil
	}

	// Load data from payments.json (if exists)
	payments, err := loadPaymentsFromJSON("data/payments.json")
	if err != nil {
		payments = []model.PaymentModel{}
	}
	repo.payments = payments

	return repo
}

func loadPaymentsFromJSON(fileName string) ([]model.PaymentModel, error) {
	file, err := os.OpenFile(fileName, os.O_RDONLY|os.O_CREATE, 0755)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var payments []model.PaymentModel
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&payments)
	if err != nil {
		return nil, err
	}

	return payments, nil
}

func writeJSONFile(fileName string, data interface{}) error {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return err
	}

	return nil
}
