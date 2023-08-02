package model

type UserModel struct {
	Id       string `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
}