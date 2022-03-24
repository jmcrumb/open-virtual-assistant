package accounts

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	Route(router.Group("/account"))
	return router
}

func TestGetAccountByID(t *testing.T) {
	router := setupRouter()

	tests := []struct {
		body   string
		status int
		err    string
	}{
		{
			`{"Text":"This is a test"}`,
			http.StatusOK,
			"",
		},
		{
			`{"wrong":"This is an invalid test"}`,
			http.StatusBadRequest,
			"empty 'text' field in json request",
		},
		{
			`This is also an invalid test`,
			http.StatusBadRequest,
			"unable to unmarhsall request body",
		},
	}

	for _, test := range tests {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/nlp/text2speech", strings.NewReader(test.body))
		router.ServeHTTP(w, req)

		assert.Equal(t, test.status, w.Code)
		if w.Code == http.StatusOK {
			assert.Equal(t, test.body, w.Body.String())
		} else {
			assert.Equal(t, test.err, w.Body.String())
		}
	}
}

func TestSTT(t *testing.T) {

}
