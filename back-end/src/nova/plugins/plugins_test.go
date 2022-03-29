package plugins

import (
	"log"
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
func queryPluginRows() []interface{} {
	var plugins []database.Plugin
	Result := database.DB.Table("plugin").Find(&plugins).Error
	if Result != nil {
		log.Fatalf("%v", Result)
	}

	var rows []interface{}
	for _, plugin := range plugins {
		rows = append(rows, plugin)
	}

	return rows
}

func TestMain(m *testing.M) {
	database.SetupDB()
	database.InitializeDB()

	router = gin.Default()
	Route(router.Group("/plugins"))
	router.SetTrustedProxies([]string{"localhost"})

	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestPostPlugin(t *testing.T) {
	database.ClearDB()

	account, _ := database.GetTestAccount()
	tests := []apitest.APITest{
		{
			Body: database.NewPlugin{
				Publisher:  account,
				SourceLink: "https://source.com/plugin/download",
				About:      "a short description about the plugin",
			},
			Status: http.StatusCreated,
			Result: "",
			Rows: []interface{}{
				database.Plugin{
					Publisher:     account,
					SourceLink:    "https://source.com/plugin/download",
					About:         "a short description about the plugin",
					DownloadCount: 0,
				},
			},
		},
		{
			Body:   `{"invalid":"test"}`,
			Status: http.StatusBadRequest,
			Result: "invalid publisher id provided: \"\"",
			Rows: []interface{}{
				database.Plugin{
					Publisher:     account,
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

		Method:  "POST",
		BaseURL: "/plugins/",

		QueryRows:  queryPluginRows,
		Comparator: comparePlugins,
	})
}

func TestPutPlugin(t *testing.T) {
	database.ClearDB()

	account, _ := database.GetTestAccount()
	newPlugin := database.NewPlugin{
		Publisher:  account,
		SourceLink: "https://source.com/plugin/download",
		About:      "a short description about the plugin",
	}
	database.DB.Table("plugin").Create(&newPlugin)
	var plugin database.Plugin
	database.DB.Table("plugin").Where("source_link = ?", newPlugin.SourceLink).Find(&plugin)

	tests := []apitest.APITest{
		{
			Body: database.Plugin{
				ID:         plugin.ID,
				SourceLink: "https://source.com/plugin/new-download",
				About:      "the description has changed",
			},
			Status: http.StatusCreated,
			Result: "",
			Rows: []interface{}{
				database.Plugin{
					Publisher:     account,
					SourceLink:    "https://source.com/plugin/new-download",
					About:         "the description has changed",
					DownloadCount: 0,
				},
			},
		},
		{
			Body:   "non-unmarshallable",
			Status: http.StatusBadRequest,
			Result: "unable to unmarshall request body",
			Rows: []interface{}{
				database.Plugin{
					Publisher:     account,
					SourceLink:    "https://source.com/plugin/new-download",
					About:         "the description has changed",
					DownloadCount: 0,
				},
			},
		},
	}

	apitest.TryRequests(apitest.APITestArgs{
		T:      t,
		Router: router,
		Tests:  tests,

		Method:  "PUT",
		BaseURL: "/plugins/",

		QueryRows:  queryPluginRows,
		Comparator: comparePlugins,
	})
}

func TestDeletePlugin(t *testing.T) {
	database.ClearDB()

	account, _ := database.GetTestAccount()
	newPlugin := database.NewPlugin{
		Publisher:  account,
		SourceLink: "https://source.com/plugin/download",
		About:      "a short description about the plugin",
	}
	database.DB.Table("plugin").Create(&newPlugin)
	var plugin database.Plugin
	database.DB.Table("plugin").Where("source_link = ?", newPlugin.SourceLink).Find(&plugin)

	tests := []apitest.APITest{
		{
			URL:    plugin.ID,
			Body:   "",
			Status: http.StatusNoContent,
			Result: "",
			Rows:   []interface{}{},
		},
	}

	apitest.TryRequests(apitest.APITestArgs{
		T:      t,
		Router: router,
		Tests:  tests,

		Method:  "DELETE",
		BaseURL: "/plugins/",

		QueryRows:  queryPluginRows,
		Comparator: comparePlugins,
	})
}

func TestGetPlugin(t *testing.T) {
	database.ClearDB()

	account, _ := database.GetTestAccount()
	newPlugin := database.NewPlugin{
		Publisher:  account,
		SourceLink: "https://source.com/plugin/download",
		About:      "a short description about the plugin",
	}
	database.DB.Table("plugin").Create(&newPlugin)
	var plugin database.Plugin
	database.DB.Table("plugin").Where("source_link = ?", newPlugin.SourceLink).Find(&plugin)

	tests := []apitest.APITest{
		{
			URL:    plugin.ID,
			Status: http.StatusNoContent,
			Result: plugin,
			Rows: []interface{}{
				plugin,
			},
		},
		{
			URL:    "invalid",
			Status: http.StatusBadRequest,
			Result: "invalid plugin ID",
			Rows: []interface{}{
				plugin,
			},
		},
	}

	apitest.TryRequests(apitest.APITestArgs{
		T:      t,
		Router: router,
		Tests:  tests,

		Method:  "GET",
		BaseURL: "/plugins/",

		QueryRows:  queryPluginRows,
		Comparator: comparePlugins,
	})
}
