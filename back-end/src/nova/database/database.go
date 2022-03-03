package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "usdnova"
	dbname   = "postgres"
)

var DB *gorm.DB

func SetupDB() error {
	var err error

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return err
}

type Account struct {
	ID         string `json:"id"`
	Password   string `json:"password"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string `json:"email"`
	LastLogin  string `json:"last_login"`
	DateJoined string `json:"date_joined"`
	IsAdmin    bool   `json:"is_admin"`
}

type NewAccount struct {
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type Profile struct {
	AccountID string `json:"account_id"`
	Bio       string `json:"bio"`
	Photo     []byte `json:"photo"`
}

type Review struct {
	ID string `json:"id"`
	SourceReview string `json:"source_review"`
	Account string `json:"account"`
	Plugin string `json:"plugin"`
	Rating float32 `json:"rating"`
	Content string `json:"content"`
}

type NewReview struct {
	SourceReview string `json:"source_review"`
	Account string `json:"account"`
	Plugin string `json:"plugin"`
	Rating float32 `json:"rating"`
	Content string `json:"content"`
}

type Report struct {
	ID string `json:"id"`
	Account string `json:"account"`
	Plugin string `json:"plugin"`
	Content string `json:"content"`
}

type NewReport struct {
	Account string `json:"account"`
	Plugin string `json:"plugin"`
	Content string `json:"content"`
}