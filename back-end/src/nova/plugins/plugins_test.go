package plugins

import (
	"database/sql"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmcrumb/nova/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Suite struct {
	db     *gorm.DB
	mock   sqlmock.Sqlmock
	mockDB *sql.DB
}

func NewSuite() *Suite {
	var (
		s   Suite
		err error
	)

	s.mockDB, s.mock, err = sqlmock.New()
	if err != nil {
		log.Fatalf("failed to create database connection: %v", err)
	}

	s.db, err = gorm.Open(postgres.New(postgres.Config{
		Conn: s.mockDB,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to open postgres database: %v", err)
	}

	s.db.Logger.LogMode(logger.Info)
	return &s
}

func TestPostPlugin(t *testing.T) {
	s := NewSuite()
	defer s.mockDB.Close()
	tests := []database.Plugin{
		{
			ID:         "12345",
			SourceLink: "https://plugin.com",
		},
		{
			ID:            "123",
			Publisher:     "maxon",
			SourceLink:    "github.com/maxon/plugin",
			About:         "a great plugin with many featurers",
			DownloadCount: 500,
			PublishedOn:   "2/22/22 22:22:22",
		},
	}

	query := regexp.QuoteMeta(
		`INSERT INTO \"plugin\" (\"id\", \"publisher\", \"source_link\", \"about\", \"download_count\", \"published_on\"
		VALUES (?,?,?,?,?,?) RETURNING \"plugin\"`)
	for _, test := range tests {
		s.mock.ExpectQuery(query).
			WithArgs(test.ID, test.Publisher, test.SourceLink, test.About, test.DownloadCount, test.PublishedOn).
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "publisher", "source_link", "about", "download_count", "published_on"}).
					AddRow(test.ID, test.Publisher, test.SourceLink, test.About, test.DownloadCount, test.PublishedOn),
			)

		s.db.Exec(query, test.ID, test.Publisher, test.SourceLink)
		err := s.db.Table("plugin").Create(&test).Error
		if err != nil {
			t.Errorf("Query failed: %v", err)
		} else if err := s.mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	}
}

func TestPutPlugin(t *testing.T) {

}

func TestDeletePlugin(t *testing.T) {

}
