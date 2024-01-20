package helpers

import (
	"os"
	"time"

	"github.com/automa8e_clone/config"
	"github.com/automa8e_clone/models"
	"github.com/golang-jwt/jwt"
)

func GenerateJWT(tokenType string,expireMin int, user *models.User) (string, error) {
	
	secretKey := []byte(config.AppConfig.JWT_SECRET)


	var claims jwt.MapClaims = jwt.MapClaims{}

	claims["exp"] = time.Now().Add(time.Minute * time.Duration(expireMin)).Unix()
	claims["sub"] = user.Id
	claims["email"] = user.Email
	claims["version"] = os.Getenv("JWT_TOKEN_VERSION")
	claims["ref"] = tokenType

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseJWT(tokenString string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AppConfig.JWT_SECRET),nil
	})

	if err != nil {
		return claims, err
	}

	return claims, nil
}