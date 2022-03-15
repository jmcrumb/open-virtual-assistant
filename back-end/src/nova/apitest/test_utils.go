package apitest

import (
	"testing"

	"github.com/jmcrumb/nova/database"
)

type APITest struct {
	Body   interface{}
	Status int
	Err    string

	Rows []interface{}
}
type ComparatorFunc func(a, b interface{}) bool

func AddDBItems(table string, rows ...interface{}) {
	db := database.DB

	for _, item := range rows {
		db.Table(table).Create(&item)
	}
}
func AssertResultsEqual(t *testing.T, expected []interface{}, actual []interface{}, areEqual ComparatorFunc) {
	if len(expected) != len(actual) {
		t.Errorf("len(rows) == %d, want %d", len(actual), len(expected))
	}

	for _, got := range actual {
		found := false
		for _, want := range expected {
			if !areEqual(want, got) {
				continue
			}

			found = true
			break
		}

		if !found {
			t.Errorf("missing record from database: %v", got)
		}
	}
}
