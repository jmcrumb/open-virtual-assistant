package accounts

import (
	"encoding/json"
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
func compareProfiles(p1, p2 interface{}) bool {
	a, b := p1.(database.Profile), p2.(database.Profile)
	bytesEqual := func(a, b []byte) bool {
		if len(a) != len(b) {
			return false
		}
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}
	return a.AccountID == b.AccountID && a.Bio == b.Bio && bytesEqual(a.Photo, b.Photo)
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
func queryProfileRows() []interface{} {
	var profiles []database.Profile
	err := database.DB.Table("profile").Find(&profiles).Error
	if err != nil {
		log.Fatalf("%v", err)
	}

	var rows []interface{}
	for _, account := range profiles {
		rows = append(rows, account)
	}

	return rows
}

func TestMain(m *testing.M) {
	database.SetupDB()
	database.InitializeDB()

	// router = gin.New()
	// router.Use(gin.Recovery())
	router = gin.Default()
	Route(router.Group("/accounts"))
	router.SetTrustedProxies([]string{"localhost"})

	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestGetAccountByID(t *testing.T) {
	database.ClearDB()

	id, _ := database.GetTestAccount()
	var account database.Account
	database.DB.Table("account").Where("id = ?", id).First(&account)

	var accountBytes []byte
	accountBytes, _ = json.Marshal(&account)

	tests := []apitest.APITest{
		{
			URL:    id,
			Status: http.StatusOK,
			Result: string(accountBytes),
			Rows: []interface{}{
				account,
			},
		},
		{
			URL:    "badID",
			Status: http.StatusNotFound,
			Result: "no account found with id: badID",
			Rows: []interface{}{
				account,
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

func TestGetProfileByID(t *testing.T) {
	database.ClearDB()

	id, _ := database.GetTestAccount()
	database.DB.Table("profile").Create(&database.Profile{
		AccountID: id,
		Bio:       "account bio",
		Photo:     []byte{1, 2, 3, 4, 5},
	})
	var profile database.Profile
	database.DB.Table("profile").Where("account_id = ?", id).First(&profile)

	var profileBytes []byte
	profileBytes, _ = json.Marshal(&profile)

	tests := []apitest.APITest{
		{
			URL:    profile.AccountID,
			Status: http.StatusOK,
			Result: string(profileBytes),
			Rows: []interface{}{
				profile,
			},
		},
		{
			URL:    "badID",
			Status: http.StatusNotFound,
			Result: "no profile found with id: badID",
			Rows: []interface{}{
				profile,
			},
		},
	}

	apitest.TryRequests(apitest.APITestArgs{
		T:      t,
		Router: router,
		Tests:  tests,

		Method:  "GET",
		BaseURL: "/accounts/profiles/",

		QueryRows:  queryProfileRows,
		Comparator: compareProfiles,
	})
}

func TestPostAccount(t *testing.T) {
	database.ClearDB()

	tests := []apitest.APITest{
		{
			Body: database.NewAccount{
				Password:  "account1_password",
				FirstName: "John",
				LastName:  "Doe",
				Email:     "jdoe@novatest.com",
			},
			Status: http.StatusCreated,
			Result: `.+`,
			Rows: []interface{}{
				database.Account{
					Password:  "account1_password",
					FirstName: "John",
					LastName:  "Doe",
					Email:     "jdoe@novatest.com",
				},
			},
		},
		{
			Body:   `invalid can't be unmarshalled`,
			Status: http.StatusBadRequest,
			Result: "unable to unmarshall request body",
			Rows: []interface{}{
				database.Account{
					Password:  "account1_password",
					FirstName: "John",
					LastName:  "Doe",
					Email:     "jdoe@novatest.com",
				},
			},
		},
	}

	apitest.TryRequests(apitest.APITestArgs{
		T:      t,
		Router: router,
		Tests:  tests,

		Method:  "POST",
		BaseURL: "/accounts/",

		QueryRows:  queryAccountRows,
		Comparator: compareAccounts,
	})
}
