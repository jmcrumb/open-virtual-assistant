package apitest

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/jmcrumb/nova/database"
)

// add ability to check on the result given back by a request (for GET)
type APITest struct {
	URL    string
	Body   interface{}
	Status int
	Result interface{}

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
		ctx := context.WithValue(req.Context(), "account_id", database.GetTestAccount().ID)
		req = req.WithContext(ctx)
		args.Router.ServeHTTP(w, req)

		// check result body against expected result
		var result string
		if reflect.TypeOf(test.Result) == reflect.TypeOf("") {
			result = test.Result.(string)
			assert.MatchRegex(args.T, w.Body.String(), result)
		} else {
			marshalled, _ := json.Marshal(&test.Result)
			result = string(marshalled)
			assert.Equal(args.T, result, w.Body.String())
		}

		// check http result values
		assert.Equal(args.T, test.Status, w.Code)

		// check database
		rows := args.QueryRows()
		AssertResultsEqual(args.T, test.Rows, rows, args.Comparator)
	}
}
