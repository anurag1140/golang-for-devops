package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("my-super-secret-key")

func GenerateToken(username, role string) (string, error) {

	claims := Claims{
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}

func ValidateToken(tokenString string) (*Claims, error) {

	fmt.Println("Received Token:")
	fmt.Println(tokenString)

	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		},
	)

	if err != nil {
		fmt.Println("JWT ERROR:", err)
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)

	if !ok || !token.Valid {
		fmt.Println("Claims invalid")
		return nil, jwt.ErrTokenInvalidClaims
	}

	fmt.Println("JWT VALID")
	fmt.Printf("%+v\n", claims)

	return claims, nil
}
