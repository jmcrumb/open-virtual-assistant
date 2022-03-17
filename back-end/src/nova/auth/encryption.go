package auth

import (
	"errors"
	"log"

	"github.com/jmcrumb/nova/database"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pwd string) string {
	bpwd := []byte(pwd)

	hash, err := bcrypt.GenerateFromPassword(bpwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}

func ComparePassword(account database.Account, pwd string) bool {
	bpwd := []byte(pwd)
	bhash := []byte(account.Password)

	err := bcrypt.CompareHashAndPassword(bhash, bpwd)

	return err == nil
}

func ResetPassword(account database.Account, submitted_old_pwd string, new_pwd string) error {
	verify_pwd := ComparePassword(account, submitted_old_pwd)
	if verify_pwd {
		account.Password = HashPassword(new_pwd)
		database.DB.Table("account").Where("id = ?", account.ID).Select("Password").Updates(&account)
		return nil
	}
	return errors.New("password reset failed")
}
