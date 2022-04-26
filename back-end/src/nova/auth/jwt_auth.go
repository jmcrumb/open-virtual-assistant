package auth

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jmcrumb/nova/database"
)

//jwt service
type JWTService interface {
	GenerateToken(id string) string
	ValidateToken(token string) (*jwt.Token, error)
}
type authCustomClaims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
	issure    string
}

//auth-jwt
func JWTAuthService() JWTService {
	return &jwtServices{
		secretKey: getSecretKey(),
		issure:    "NOVA-PREALPHA",
	}
}

func getSecretKey() string {
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (service *jwtServices) GenerateToken(id string) string {
	claims := &authCustomClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    service.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("invalid token: %v", token.Header["alg"])

		}
		return []byte(service.secretKey), nil
	})
}

func EnforceMiddlewareAuthentication(c *gin.Context, id string, f func(id string)) {
	fmt.Println(c.Request.Context().Value("account_id").(string))
	if c.Request.Context().Value("account_id").(string) == id {
		f(id)
		return
	}
	c.String(http.StatusUnauthorized, "Permission Denied")
}

func EnforceMiddlewareAuthenticatedAdmin(c *gin.Context, f func()) {
	auth_id, _ := c.Request.Context().Value("account_id").(string)
	var account database.Account
	database.DB.Table("account").Where("id = ?", auth_id).First(&account)

	if account.IsAdmin {
		f()
		return
	}
	c.String(http.StatusUnauthorized, "Permission Denied")
}
