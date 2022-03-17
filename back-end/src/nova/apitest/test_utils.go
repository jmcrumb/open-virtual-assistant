package apitest

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

// add ability to check on the result given back by a request (for GET)
type APITest struct {
	URL    string
	Body   interface{}
	Status int
	Err    string

	Rows []interface{}
}
type ComparatorFunc func(a, b interface{}) bool

type APITestArgs struct {
	T      *testing.T
	Router *gin.Engine
	Tests  []APITest

	Method  string
	BaseURL string

	QueryRows  func() []interface{}
	Comparator ComparatorFunc
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

func TryRequests(args APITestArgs) {
	for _, test := range args.Tests {
		// perform request
		w := httptest.NewRecorder()
		var body string
		if reflect.TypeOf(test.Body) == reflect.TypeOf("") {
			body = test.Body.(string)
		} else {
			marshalled, _ := json.Marshal(&test.Body)
			body = string(marshalled)
		}
		req, _ := http.NewRequest(args.Method, args.BaseURL+test.URL, strings.NewReader(body))
		args.Router.ServeHTTP(w, req)

		// check http result values
		assert.Equal(args.T, test.Status, w.Code)
		if w.Code == http.StatusBadRequest {
			assert.Equal(args.T, test.Err, w.Body.String())
		}

		// check database
		rows := args.QueryRows()
		AssertResultsEqual(args.T, test.Rows, rows, args.Comparator)
	}
}
