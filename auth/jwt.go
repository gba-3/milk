package auth

import (
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
	"os"
	"time"
)

func CreateToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"iat":   time.Now(),
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})
	// 電子署名
	tokenString, err := token.SignedString([]byte(os.Getenv("SIGNINKEY")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SIGNINKEY")), nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})
