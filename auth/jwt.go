package auth

import (
	"os"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
)

func CreateToken(email string) (string, error) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"iat":   time.Now(),
		"exp":   time.Now().In(jst).Add(time.Minute * 3).Unix(),
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
