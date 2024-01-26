package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "dummySecretKey"
func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"userId": userId,
		"exp": time.Now().Add(time.Hour*2).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

// In anonymous function we are checking input token encryption type is it same as the one in which we encrypted.
func VerifyToken(token string) (int64,error){
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error){
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok{
			return nil, errors.New("Unexpected signing method.")
		}
		return []byte(secretKey), nil
	})

	if err !=nil{
		return 0, errors.New("Could not parse token.")
	}
	tokenIsValid := parsedToken.Valid
	if !tokenIsValid{
		return 0, errors.New("Invalid token!")
	}
	//For learning purpose
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok{
		return 0, errors.New("Invalid token claims.")
	}
	// .(string) is used for type checking and telling email is of type string. This will return 2 values first string second ok(bool)
	// email, ok :=claims["email"].(string)
	// like this we can extract the email value from the token
	// email := claims["email"].(string)
	userId := int64(claims["userId"].(float64))

	return userId, nil
}