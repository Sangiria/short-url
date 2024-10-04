package jwt

import (
	"short_url/enviroment"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateJwtToken() (string, error) {
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(48 * time.Hour).Unix(),
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, err := rawToken.SignedString([]byte(enviroment.GoDotEnvVariable("Secretkey")))
	
	if err != nil {
		return "", err
	}
	return token, nil
}

