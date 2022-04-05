package reviews

import (
	"fmt"
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

	account := database.GetTestAccount().ID
	plugin := database.GetTestPlugin(account).ID
	review := database.GetTestReview(account, plugin)
	tests := []apitest.APITest{
		{
			URL:    plugin,
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

	account := database.GetTestAccount().ID
	plugin := database.GetTestPlugin(account).ID
	tests := []apitest.APITest{
		{
			Body: database.NewReview{
				Account: account,
				Plugin:  plugin,
				Rating:  4.5,
				Content: "This plugin is great except that it's a little bit slow",
			},
			Status: http.StatusCreated,
			Result: `.+`,
			Rows: []interface{}{
				database.Review{
					Account: account,
					Plugin:  plugin,
					Rating:  4.5,
					Content: "This plugin is great except that it's a little bit slow",
				},
			},
		},
		{
			Body:   `{"invalid":"test"}`,
			Status: http.StatusBadRequest,
			Result: "invalid account or plugin id provided: {account: \"\", plugin: \"\"}",
			Rows: []interface{}{
				database.Review{
					Account: account,
					Plugin:  plugin,
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

	account := database.GetTestAccount().ID
	plugin := database.GetTestPlugin(account).ID
	review := database.GetTestReview(account, plugin)

	tests := []apitest.APITest{
		{
			Body: database.Review{
				ID:      review.ID,
				Account: "new account (this won't actually be updated)",
				Rating:  2.2,
				Content: "this plugin is actually terrible",
			},
			Status: http.StatusCreated,
			Result: "",
			Rows: []interface{}{
				database.Review{
					Account: account,
					Plugin:  plugin,
					Rating:  2.2,
					Content: "this plugin is actually terrible",
				},
			},
		},
		{
			Body:   "non-unmarshallable",
			Status: http.StatusBadRequest,
			Result: "unable to unmarshall request body",
			Rows: []interface{}{
				database.Review{
					Account: account,
					Plugin:  plugin,
					Rating:  2.2,
					Content: "this plugin is actually terrible",
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

	account := database.GetTestAccount().ID
	plugin := database.GetTestPlugin(account).ID
	review := database.GetTestReview(account, plugin)

	tests := []apitest.APITest{
		{
			URL:    fmt.Sprintf("%v/%v", plugin, review.ID),
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

	account := database.GetTestAccount().ID
	plugin := database.GetTestPlugin(account).ID
	review := database.GetTestReview(account, plugin)

	tests := []apitest.APITest{
		{
			URL:    fmt.Sprintf("%v/%v", plugin, review.ID),
			Status: http.StatusOK,
			Result: review,
			Rows: []interface{}{
				review,
			},
		},
		{
			URL:    fmt.Sprintf("%v/%v", plugin, "invalid"),
			Status: http.StatusBadRequest,
			Result: "invalid review ID",
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
