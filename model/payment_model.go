package model

import "time"

type PaymentModel struct {
	Id            string    `json:"id"`
	UserId        string    `json:"user_id"`
	MerchantNoRek string    `json:"merchant_no_rek"`
	Amount        string    `json:"amount"`
	CreatedAt     time.Time `json:"created_at"`
}

