package main

type UserInfoType string

const (
	UserInfoKey UserInfoType = "UserInfoKey"
)

type UserInfo struct {
	UserID  string
	Email   string
	IsAdmin bool
}

type UserRegister struct {
	Email   string `json:"email"`
	IsAdmin bool   `json:"isAdmin"`
}
