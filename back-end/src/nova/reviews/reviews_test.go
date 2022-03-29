package reviews

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

func compareReviews(r1, r2 interface{}) bool {
	a, b := r1.(database.Review), r2.(database.Review)
	return a.SourceReview == b.SourceReview && a.Account == b.Account && a.Plugin == b.Plugin && a.Rating == b.Rating && a.Content == b.Content
}
func queryReviewRows() []interface{} {
	var reviews []database.Review
	Result := database.DB.Table("review").Find(&reviews).Error
	if Result != nil {
		log.Fatalf("%v", Result)
	}

	var rows []interface{}
	for _, review := range reviews {
		rows = append(rows, review)
	}

	return rows
}

func TestMain(m *testing.M) {
	database.SetupDB()
	database.InitializeDB()

	router = gin.Default()
	Route(router.Group("/reviews"))
	router.SetTrustedProxies([]string{"localhost"})

	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestGetReviews(t *testing.T) {
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
	newReview := database.NewReview{
		Account: account,
		Plugin:  plugin.ID,
		Rating:  4.5,
		Content: "This plugin is great except that it's a little bit slow",
	}
	database.DB.Table("review").Create(&newReview)
	var review database.Review
	database.DB.Table("review").Where("plugin = ?", newReview.Plugin).Find(&review)

	tests := []apitest.APITest{
		{
			URL:    plugin.ID,
			Status: http.StatusOK,
			Result: []database.Review{review},
			Rows: []interface{}{
				review,
			},
		},
		{
			URL:    "invalid",
			Status: http.StatusBadRequest,
			Result: "invalid plugin ID",
			Rows: []interface{}{
				review,
			},
		},
	}

	apitest.TryRequests(apitest.APITestArgs{
		T:      t,
		Router: router,
		Tests:  tests,

		Method:  "GET",
		BaseURL: "/reviews/",

		QueryRows:  queryReviewRows,
		Comparator: compareReviews,
	})
}

func TestPostReview(t *testing.T) {
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
			URL: plugin.ID,
			Body: database.NewReview{
				Account: account,
				Plugin:  plugin.ID,
				Rating:  4.5,
				Content: "This plugin is great except that it's a little bit slow",
			},
			Status: http.StatusCreated,
			Result: `.+`,
			Rows: []interface{}{
				database.Review{
					Account: account,
					Plugin:  plugin.ID,
					Rating:  4.5,
					Content: "This plugin is great except that it's a little bit slow",
				},
			},
		},
		{
			Body:   `{"invalid":"test"}`,
			Status: http.StatusBadRequest,
			Result: "invalid account id provided: \"\"",
			Rows: []interface{}{
				database.Review{
					Account: account,
					Plugin:  plugin.ID,
					Rating:  4.5,
					Content: "This plugin is great except that it's a little bit slow",
				},
			},
		},
	}

	apitest.TryRequests(apitest.APITestArgs{
		T:      t,
		Router: router,
		Tests:  tests,

		Method:  "POST",
		BaseURL: "/reviews/",

		QueryRows:  queryReviewRows,
		Comparator: compareReviews,
	})
}

func TestPutReview(t *testing.T) {
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
	newReview := database.NewReview{
		Account: account,
		Plugin:  plugin.ID,
		Rating:  4.5,
		Content: "This plugin is great except that it's a little bit slow",
	}
	database.DB.Table("review").Create(&newReview)
	var review database.Review
	database.DB.Table("review").Where("plugin = ?", newReview.Plugin).Find(&review)

	tests := []apitest.APITest{
		{
			Body: database.Review{
				ID:      review.ID,
				Rating:  2.2,
				Content: "this plugin is actually terrible",
			},
			Status: http.StatusCreated,
			Result: "",
			Rows: []interface{}{
				database.Review{
					Account: account,
					Plugin:  plugin.ID,
					Rating:  2.2,
					Content: "this plugin is actually terrible",
				},
			},
		},
		{
			/*
				TODO
			*/
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
		BaseURL: "/reviews/",

		QueryRows:  queryReviewRows,
		Comparator: compareReviews,
	})
}

func TestDeleteReview(t *testing.T) {
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
		BaseURL: "/reviews/",

		QueryRows:  queryReviewRows,
		Comparator: compareReviews,
	})
}

func TestGetReview(t *testing.T) {
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
		BaseURL: "/reviews/",

		QueryRows:  queryReviewRows,
		Comparator: compareReviews,
	})
}
