package models

import (
	"fmt"
	"os"
	"time"

	DBConfig "../Config"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
)

type Toaken struct {
	ToakenID string `json:"toakenId"`
	UserID   int    `json:"userId"`
	//limitTime time.Time `json:"limitTime"`
}

func (b *Toaken) TableName() string {
	return "toaken_id"
}

func CreateToaken(token *Toaken) (err error) {
	if err = DBConfig.DB.Create(token).Error; err != nil {
		return err
	}
	return nil
}

type JWTService interface {
	GenerateToaken(loginID string, isUser bool) string
	ValidateToken(token string) (*jwt.Token, error)
}

type authCustomClaims struct {
	Name string `json:"name"`
	User bool   `json:"user"`
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
	issure    string
}

func JWTAuthService() JWTService {
	return &jwtServices{
		secretKey: getSecretKey(),
		issure:    "Bikash",
	}
}

func getSecretKey() string {
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (service *jwtServices) GenerateToaken(loginID string, isUser bool) string {
	claims := &authCustomClaims{
		loginID,
		isUser,
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
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])

		}
		return []byte(service.secretKey), nil
	})

}
