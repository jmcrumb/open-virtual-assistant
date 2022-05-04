package reports

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

func compareReports(r1, r2 interface{}) bool {
	a, b := r1.(database.Report), r2.(database.Report)
	return a.Account == b.Account && a.Plugin == b.Plugin && a.Content == b.Content
}
func queryReportRows() []interface{} {
	var reports []database.Report
	Result := database.DB.Table("report").Find(&reports).Error
	if Result != nil {
		log.Fatalf("%v", Result)
	}

	var rows []interface{}
	for _, report := range reports {
		rows = append(rows, report)
	}

	return rows
}

func TestMain(m *testing.M) {
	database.SetupDB()
	database.InitializeDB()

	router = gin.Default()
	Route(router.Group("/reports"))
	router.SetTrustedProxies([]string{"localhost"})

	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestGetReports(t *testing.T) {
	database.ClearDB()

	account := database.GetTestAccount().ID
	plugin := database.GetTestPlugin(account).ID
	report := database.GetTestReport(account, plugin)
	tests := []apitest.APITest{
		{
			URL:    plugin,
			Status: http.StatusOK,
			Result: []database.Report{report},
			Rows: []interface{}{
				report,
			},
			AuthorizedUser: account,
		},
		{
			URL:    "invalid",
			Status: http.StatusBadRequest,
			Result: "invalid plugin ID",
			Rows: []interface{}{
				report,
			},
			AuthorizedUser: account,
		},
	}

	apitest.TryRequests(apitest.APITestArgs{
		T:      t,
		Router: router,
		Tests:  tests,

		Method:  "GET",
		BaseURL: "/reports/",

		QueryRows:  queryReportRows,
		Comparator: compareReports,
	})
}

func TestPostReport(t *testing.T) {
	database.ClearDB()

	account := database.GetTestAccount().ID
	plugin := database.GetTestPlugin(account).ID
	tests := []apitest.APITest{
		{
			Body: database.NewReport{
				Account: account,
				Plugin:  plugin,
				Content: "This plugin broke all of my devices",
			},
			Status: http.StatusCreated,
			Result: `.+`,
			Rows: []interface{}{
				database.Report{
					Account: account,
					Plugin:  plugin,
					Content: "This plugin broke all of my devices",
				},
			},
			AuthorizedUser: account,
		},
		{
			Body:   `{"invalid":"test"}`,
			Status: http.StatusBadRequest,
			Result: "invalid account or plugin id provided: {account: \"\", plugin: \"\"}",
			Rows: []interface{}{
				database.Report{
					Account: account,
					Plugin:  plugin,
					Content: "This plugin broke all of my devices",
				},
			},
			AuthorizedUser: account,
		},
	}

	apitest.TryRequests(apitest.APITestArgs{
		T:      t,
		Router: router,
		Tests:  tests,

		Method:  "POST",
		BaseURL: "/reports/",

		QueryRows:  queryReportRows,
		Comparator: compareReports,
	})
}

func TestPutReport(t *testing.T) {
	database.ClearDB()

	account := database.GetTestAccount().ID
	plugin := database.GetTestPlugin(account).ID
	report := database.GetTestReport(account, plugin)

	tests := []apitest.APITest{
		{
			Body: database.Report{
				ID:      report.ID,
				Account: "new account (this won't actually be updated)",
				Content: "this plugin made my phone blow up",
			},
			Status: http.StatusCreated,
			Result: "",
			Rows: []interface{}{
				database.Report{
					Account: account,
					Plugin:  plugin,
					Content: "this plugin made my phone blow up",
				},
			},
			AuthorizedUser: account,
		},
		{
			Body:   "non-unmarshallable",
			Status: http.StatusBadRequest,
			Result: "unable to unmarshall request body",
			Rows: []interface{}{
				database.Report{
					Account: account,
					Plugin:  plugin,
					Content: "this plugin made my phone blow up",
				},
			},
			AuthorizedUser: account,
		},
	}

	apitest.TryRequests(apitest.APITestArgs{
		T:      t,
		Router: router,
		Tests:  tests,

		Method:  "PUT",
		BaseURL: "/reports/",

		QueryRows:  queryReportRows,
		Comparator: compareReports,
	})
}

func TestDeleteReport(t *testing.T) {
	database.ClearDB()

	account := database.GetTestAccount().ID
	plugin := database.GetTestPlugin(account).ID
	report := database.GetTestReport(account, plugin)

	tests := []apitest.APITest{
		{
			URL:            fmt.Sprintf("%v/%v", plugin, report.ID),
			Body:           "",
			Status:         http.StatusNoContent,
			Result:         "",
			Rows:           []interface{}{},
			AuthorizedUser: account,
		},
	}

	apitest.TryRequests(apitest.APITestArgs{
		T:      t,
		Router: router,
		Tests:  tests,

		Method:  "DELETE",
		BaseURL: "/reports/",

		QueryRows:  queryReportRows,
		Comparator: compareReports,
	})
}

func TestGetReport(t *testing.T) {
	database.ClearDB()

	account := database.GetTestAccount().ID
	plugin := database.GetTestPlugin(account).ID
	report := database.GetTestReport(account, plugin)

	tests := []apitest.APITest{
		{
			URL:    fmt.Sprintf("%v/%v", plugin, report.ID),
			Status: http.StatusOK,
			Result: report,
			Rows: []interface{}{
				report,
			},
			AuthorizedUser: account,
		},
		{
			URL:    fmt.Sprintf("%v/%v", plugin, "invalid"),
			Status: http.StatusBadRequest,
			Result: "invalid report ID",
			Rows: []interface{}{
				report,
			},
			AuthorizedUser: account,
		},
	}

	apitest.TryRequests(apitest.APITestArgs{
		T:      t,
		Router: router,
		Tests:  tests,

		Method:  "GET",
		BaseURL: "/reports/",

		QueryRows:  queryReportRows,
		Comparator: compareReports,
	})
}
