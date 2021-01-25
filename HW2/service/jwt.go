package service

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTClaimsModel struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}

func GenerateJWT(userId string) (string, error) {
	mySigningKey := []byte("AllYourBase")

	// Create the Claims
	claims := JWTClaimsModel{
		userId,
		jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * 50).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)

	if err != nil {
		return "", err
	}
	return ss, nil
}
func DecodeJWT(tokenString string) {
	// at(time.Unix(0, 0), func() {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaimsModel{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})

	if claims, ok := token.Claims.(*JWTClaimsModel); ok && token.Valid {
		fmt.Printf("%v %v", claims.UserId, claims.StandardClaims.ExpiresAt)
	} else {
		fmt.Println(err)
	}
	// })
}
func at(t time.Time, f func()) {
	jwt.TimeFunc = func() time.Time {
		return t
	}
	f()
	jwt.TimeFunc = time.Now
}
