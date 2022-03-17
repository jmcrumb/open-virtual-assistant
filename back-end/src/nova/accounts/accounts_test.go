package accounts

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

func compareAccounts(a1, a2 interface{}) bool {
	a, b := a1.(database.Account), a2.(database.Account)
	return a.Password == b.Password && a.FirstName == b.FirstName &&
		a.LastName == b.LastName && a.Email == b.Email
}
func queryAccountRows() []interface{} {
	var accounts []database.Account
	err := database.DB.Table("account").Find(&accounts).Error
	if err != nil {
		log.Fatalf("%v", err)
	}

	var rows []interface{}
	for _, account := range accounts {
		rows = append(rows, account)
	}

	return rows
}

func TestMain(m *testing.M) {
	database.SetupDB()
	database.InitializeDB()

	router = gin.Default()
	Route(router.Group("/accounts"))
	router.SetTrustedProxies([]string{"localhost"})

	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestGetAccountByID(t *testing.T) {
	database.ClearDB()

	account, info := database.GetTestAccount()
	tests := []apitest.APITest{
		{
			URL:    account,
			Status: http.StatusOK,
			Err:    "",
			Rows: []interface{}{
				database.Account{
					ID:        account,
					Password:  info.Password,
					FirstName: info.FirstName,
					LastName:  info.LastName,
					Email:     info.Email,
				},
			},
		},
	}

	apitest.TryRequests(apitest.APITestArgs{
		T:      t,
		Router: router,
		Tests:  tests,

		Method:  "GET",
		BaseURL: "/accounts/",

		QueryRows:  queryAccountRows,
		Comparator: compareAccounts,
	})
}
