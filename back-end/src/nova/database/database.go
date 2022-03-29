package database

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres-test"
)

var DB *gorm.DB

func SetupDB() {
	var err error
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Fatal("failed to connect database")
	}
}
func InitializeDB() error {
	return DB.Exec(DBInit).Error
}
func ClearDB() error {
	return DB.Exec(DBClear).Error
}

func GetTestAccount() (id string, info NewAccount) {
	now := time.Now().Nanosecond()

	acc := NewAccount{
		Password:  "test123",
		FirstName: "user",
		LastName:  "test",
		Email:     fmt.Sprintf("user-%d@novatest.com", now),
	}
	var result Account

	DB.Table("account").Create(&acc)
	DB.Table("account").Where("email = ?", acc.Email).First(&result)

	return result.ID, acc
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

type Plugin struct {
	ID            string `json:"id"`
	Publisher     string `json:"publisher"`
	SourceLink    string `json:"source_link"`
	About         string `json:"about"`
	DownloadCount int    `json:"download_count"`
	PublishedOn   string `json:"published_on"`
}

type NewPlugin struct {
	Publisher  string `json:"publisher"`
	SourceLink string `json:"source_link"`
	About      string `json:"about"`
}

type UpdatePassword struct {
	AccountID   string `json:"account_id"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type Review struct {
	ID           string  `json:"id"`
	SourceReview string  `json:"source_review"`
	Account      string  `json:"account"`
	Plugin       string  `json:"plugin"`
	Rating       float32 `json:"rating"`
	Content      string  `json:"content"`
}

type NewReview struct {
	SourceReview string  `json:"source_review"`
	Account      string  `json:"account"`
	Plugin       string  `json:"plugin"`
	Rating       float32 `json:"rating"`
	Content      string  `json:"content"`
}

type Report struct {
	ID      string `json:"id"`
	Account string `json:"account"`
	Plugin  string `json:"plugin"`
	Content string `json:"content"`
}

type NewReport struct {
	Account string `json:"account"`
	Plugin  string `json:"plugin"`
	Content string `json:"content"`
}
