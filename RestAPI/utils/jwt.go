package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "secret123!"

func GenerateToken(email string, userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(tokenString string) error {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("could not parse token")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return err
	}

	if !parsedToken.Valid {
		return errors.New("invalid token")
	}

	//claims, ok := parsedToken.Claims.(jwt.MapClaims)
	//if !ok {
	//	return nil, errors.New("could not parse claims")
	//}
	//email, userId := claims["email"].(string), claims["userId"].(string)
	return nil
}
