package utils

import (
	"github.com/fox-wei/go_project_training/account_management/myerr"
	"golang.org/x/crypto/bcrypt"
)

func CheckParam(accountName, password string) myerr.Err {
	if accountName == "" {
		return myerr.New(*myerr.ErrValidation, myerr.ErrValidation).Add("acountName is empty.")
	}

	if password == "" {
		return myerr.New(*myerr.ErrValidation, myerr.ErrValidation).Add("password is empty.")
	}

	return myerr.Err{ErrNum: *myerr.PassParamCheck, Err: nil}
}

//?加密
func Encrypt(source string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(source), bcrypt.DefaultCost)
	return string(hashedBytes), err
}

//*比较密码
func Compare(hasdhPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hasdhPassword), []byte(password))
}
