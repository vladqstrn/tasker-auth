package utils

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

type SignedClaims struct {
	Username string
	jwt.StandardClaims
}

var SECRET_KEY string = viper.GetString("JWT.SECRET_KEY")

func GenarateTokens(username string) (token string, refreshToken string) {
	claims := SignedClaims{

		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Minute * time.Duration(1)).Unix(),
		},
	}

	refreshClaims := SignedClaims{

		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))

	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Fatal(err)
	}
	return accessToken, refreshToken
}

func ValidateToken(cookieToken string) (string, error) {
	token, err := jwt.ParseWithClaims(
		cookieToken,
		&SignedClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	var returnClaims *SignedClaims
	if claims, ok := token.Claims.(*SignedClaims); ok && token.Valid {
		returnClaims = claims
		return returnClaims.Username, err
	} else {
		return "", err
	}

}

func ParseJWTWithClaims(tokenString string) (*SignedClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &SignedClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		if validationError, ok := err.(*jwt.ValidationError); ok {
			if validationError.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, fmt.Errorf("token is expired")
			}
		}

		return nil, fmt.Errorf("couldn't parse token: %v", err)
	}

	claims, ok := token.Claims.(*SignedClaims)
	if !ok {
		return nil, fmt.Errorf("invalid claims format")
	}

	return claims, nil
}
