package authentication

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func TokenGenerate(userID uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userId"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	return token.SignedString(config.SecretKey)
}

func CheckToken(r *http.Request) error {
	tokenString := getToken(r)
	token, erro := jwt.Parse(tokenString, returnVerificationKey)

	if erro != nil {
		return erro
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("invalid token")
}

func GetUserIdFromToken(r *http.Request) (uint64, error) {
	tokenString := getToken(r)
	token, erro := jwt.Parse(tokenString, returnVerificationKey)

	if erro != nil {
		return 0, erro
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId, erro := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["userId"]), 10, 64)
		if erro != nil {
			return 0, erro
		}

		return userId, nil
	}

	return 0, errors.New("invalid token")
}

func getToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func returnVerificationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected method %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}
