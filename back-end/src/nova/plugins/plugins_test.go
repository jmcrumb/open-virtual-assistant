package plugins

import (
	"log"
	"os"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmcrumb/nova/database"
)

func TestMain(m *testing.M) {
	os.Setenv("TEST_DB", "true")

	log.Printf("Beginning tests")
	exitVal := m.Run()
	database.TestSuite.MockDB.Close()

	os.Exit(exitVal)
}

func TestPostPlugin(t *testing.T) {
	s := database.TestSuite
	s.Setup()
	defer s.Teardown()
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
		s.Mock.ExpectQuery(query).
			WithArgs(test.ID, test.Publisher, test.SourceLink, test.About, test.DownloadCount, test.PublishedOn).
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "publisher", "source_link", "about", "download_count", "published_on"}).
					AddRow(test.ID, test.Publisher, test.SourceLink, test.About, test.DownloadCount, test.PublishedOn),
			)

		s.DB.Exec(query, test.ID, test.Publisher, test.SourceLink)
		err := s.DB.Table("plugin").Create(&test).Error
		if err != nil {
			t.Errorf("Query failed: %v", err)
		} else if err := s.Mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %v", err)
		}
	}
}

func TestPutPlugin(t *testing.T) {

}

func TestDeletePlugin(t *testing.T) {

}
