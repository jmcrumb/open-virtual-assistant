package plugins

import (
	"net/http"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jmcrumb/nova/apitest"
	"github.com/jmcrumb/nova/database"
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

	apitest.TryRequests(apitest.APITestArgs{
		T:      t,
		Router: router,
		Tests:  tests,

		Method: "POST",
		URL:    "/plugin/",

		QueryRows: func() []interface{} {
			var plugins []database.Plugin
			err := database.DB.Table("plugin").Find(&plugins).Error
			if err != nil {
				t.Errorf("%v", err)
			}

			var rows []interface{}
			for _, plugin := range plugins {
				rows = append(rows, plugin)
			}

			return rows
		},
		Comparator: comparePlugins,
	})
}

func TestPutPlugin(t *testing.T) {

}

func TestDeletePlugin(t *testing.T) {

}
