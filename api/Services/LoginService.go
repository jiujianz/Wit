package service

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	models "github.com/jiujianz/Wit/Models"

	"github.com/gin-gonic/gin"
)

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

func getSecretKey() string {
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func GenerateToaken(loginID string, isUser bool) string {
	claims := &authCustomClaims{
		loginID,
		isUser,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    "Wit",
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(getSecretKey()))
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

func GetUserByLoginID(c *gin.Context) (models.User, error) {
	var user models.User
	fmt.Println(user)
	c.BindJSON(&user)
	err := models.GetUserByLoginID(&user)
	fmt.Println(err)
	if err != nil {
		//fmt.Println(c.AbortWithStatus(http.StatusNotFound))
		// TODO error response
		fmt.Println("error contents : ", err)
		return user, err
	}
	fmt.Println(user)
	return user, nil
}
