package plugins

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jmcrumb/nova/apitest"
	"github.com/jmcrumb/nova/database"
	"github.com/stretchr/testify/assert"
)

var router *gin.Engine

func comparePlugins(p1, p2 interface{}) bool {
	a, b := p1.(database.Plugin), p2.(database.Plugin)
	return a.Publisher == b.Publisher && a.SourceLink == b.SourceLink && a.About == b.About && a.DownloadCount == b.DownloadCount
}

func TestMain(m *testing.M) {
	database.SetupDB()
	database.InitializeDB()

	router = gin.Default()
	Route(router.Group("/plugin"))
	router.SetTrustedProxies([]string{"localhost"})

	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestPostPlugin(t *testing.T) {
	database.ClearDB()

	accounts := []string{
		database.GetTestAccount(),
	}
	tests := []apitest.APITest{
		{
			Body: database.NewPlugin{
				Publisher:  accounts[0],
				SourceLink: "https://source.com/plugin/download",
				About:      "a short description about the plugin",
			},
			Status: http.StatusCreated,
			Err:    "",
			Rows: []interface{}{
				database.Plugin{
					Publisher:     accounts[0],
					SourceLink:    "https://source.com/plugin/download",
					About:         "a short description about the plugin",
					DownloadCount: 0,
				},
			},
		},
		{
			Body:   `{"invalid":"test"}`,
			Status: http.StatusBadRequest,
			Err:    "invalid publisher id provided: \"\"",
			Rows: []interface{}{
				database.Plugin{
					Publisher:     accounts[0],
					SourceLink:    "https://source.com/plugin/download",
					About:         "a short description about the plugin",
					DownloadCount: 0,
				},
			},
		},
	}

	for _, test := range tests {
		// perform request
		w := httptest.NewRecorder()
		var body string
		if reflect.TypeOf(test.Body) == reflect.TypeOf("") {
			body = test.Body.(string)
		} else {
			marshalled, _ := json.Marshal(&test.Body)
			body = string(marshalled)
		}
		req, _ := http.NewRequest("POST", "/plugin/", strings.NewReader(body))
		router.ServeHTTP(w, req)

		// check http result values
		assert.Equal(t, test.Status, w.Code)
		if w.Code != http.StatusCreated {
			assert.Equal(t, test.Err, w.Body.String())
		}

		// check database
		var plugins []database.Plugin
		err := database.DB.Table("plugin").Find(&plugins).Error
		if err != nil {
			t.Errorf("%v", err)
		}

		var rows []interface{}
		for _, plugin := range plugins {
			rows = append(rows, plugin)
		}
		apitest.AssertResultsEqual(t, test.Rows, rows, comparePlugins)
	}
}

func TestPutPlugin(t *testing.T) {

}

func TestDeletePlugin(t *testing.T) {

}
