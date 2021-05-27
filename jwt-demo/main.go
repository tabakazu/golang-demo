package main

import (
	"crypto/md5"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	// My secret key
	key := []byte(os.Getenv("SIGNINGKEY"))

	//
	// building and signing a token
	//
	keyAndiat := fmt.Sprintf("%s:%d", string(key), time.Now().Unix())
	jwtID := fmt.Sprintf("%x", md5.Sum([]byte(keyAndiat)))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		// Claims https://tools.ietf.org/html/rfc7519#section-4
		"iss":     "My Awesome Company Inc. or https://my.awesome.website/", // "iss" (Issuer)
		"aud":     []string{"Young", "Old"},                                 // "aud" (Audience)
		"sub":     "Subject",                                                // "sub" (Subject)
		"iat":     time.Now().Unix(),                                        // "iat" (Issued At)
		"nbf":     time.Date(2020, 01, 01, 12, 0, 0, 0, time.UTC).Unix(),    // "nbf" (Not Before)
		"exp":     time.Now().Add(time.Hour * 24).Unix(),                    // "exp" (Expiration Time)
		"jti":     jwtID,                                                    // "jti" (JWT ID)
		"name":    "Taro Yamada",
		"admin":   false,
		"user_id": 1,
	})

	tokenString, err := token.SignedString(key)
	fmt.Println(tokenString, err)

	//
	// parsing and validating a token
	//
	token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return key, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["name"])
		fmt.Println(claims["user_id"])
		fmt.Println(claims["exp"])
		fmt.Println(claims["jti"])

		val, ok := claims["user_id"].(float64)
		if ok {
			fmt.Println(uint(val))
		}

	} else {
		fmt.Println(err)
	}
}
