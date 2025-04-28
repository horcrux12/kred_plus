package helper

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"kredi-plus.com/be/config"
	"time"
)

type MyCustomClaim struct {
	jwt.RegisteredClaims
	UserID   int64  `json:"user_id"`
	UserName string `json:"user_name"`
}

var secretKey = []byte(config.Attr.JWTSecretKey)

func GenerateJWT(userId int64, username string) (string, error) {
	claims := MyCustomClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    config.Attr.App.Name,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserID:   userId,
		UserName: username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(tokenStr string) (data *MyCustomClaim, err error) {
	token, err := jwt.ParseWithClaims(tokenStr, &MyCustomClaim{}, func(token *jwt.Token) (interface{}, error) {
		// Validate Alg
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*MyCustomClaim)
	if ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}
