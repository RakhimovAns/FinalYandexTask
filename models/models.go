package models

import (
	"errors"
	"github.com/golang-jwt/jwt"
)

type Expression struct {
	ID           int64  `json:"id"`
	Expression   string `json:"expression"`
	AddTime      int64  `json:"addTime"`
	SubTime      int64  `json:"subTime"`
	MultiplyTime int64  `json:"multiplyTime"`
	DivideTime   int64  `json:"divideTime"`
	IsCounted    bool
	Result       int64
}

type ID struct {
	ID int64 `json:"id"`
}

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

var ErrInvalidPassword = errors.New("invalid password")
var ErrUserExist = errors.New("user already exist, try other name")
var ErrUserNotExist = errors.New("no user with this name")

type TokenClaim struct {
	jwt.StandardClaims
	UserID int64 `json:"userid"`
}
