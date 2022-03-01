package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "usdnova"
	dbname   = "postgres"
)

type Suite struct {
	DB     *gorm.DB
	Mock   sqlmock.Sqlmock
	MockDB *sql.DB
}

func (*Suite) Setup() {
	DB.Exec(`\i <project root path>database_INIT.sql`)
}
func (*Suite) Teardown() {
	DB.Exec(`DROP ALL TABLES PLEASE`)
}

var TestSuite *Suite
var DB *gorm.DB

func init() {
	var err error

	if _, exists := os.LookupEnv("TEST_DB"); exists {
		TestSuite = new(Suite)
		TestSuite.MockDB, TestSuite.Mock, err = sqlmock.New()
		if err != nil {
			log.Fatalf("failed to create database connection: %v", err)
		}

		TestSuite.DB, err = gorm.Open(postgres.New(postgres.Config{
			Conn: TestSuite.MockDB,
		}), &gorm.Config{})
		if err != nil {
			log.Fatalf("failed to open postgres database: %v", err)
		}

		TestSuite.DB.Logger.LogMode(logger.Info)
		DB = TestSuite.DB
	} else {
		dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			panic("failed to connect database")
		}
	}
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
