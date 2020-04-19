package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	key []byte = []byte(os.Getenv("SIGNINGKEY"))
)

type User struct {
	Name string `json:"name"`
}

type Error struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/auth", GetTokenHandler)
	http.HandleFunc("/user", JwtMiddleware(GetUserInfoHandler))

	fmt.Println("Staring web server...")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func GetTokenHandler(w http.ResponseWriter, r *http.Request) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":  "Tabakazu. or https://tabakazu.com/",                  // "iss" (Issuer)
		"iat":  time.Now().Unix(),                                     // "iat" (Issued At)
		"nbf":  time.Date(2020, 01, 01, 12, 0, 0, 0, time.UTC).Unix(), // "nbf" (Not Before)
		"exp":  time.Now().Add(time.Hour * 24).Unix(),                 // "exp" (Expiration Time)
		"name": "Taro Yamada",
	})

	tokenString, _ := token.SignedString(key)
	w.Write([]byte(tokenString))
}

func GetUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	v := r.Context().Value("User")
	token, _ := v.(*jwt.Token)

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if name, ok := GetUserName(claims); ok {
			res := &User{name}
			json.NewEncoder(w).Encode(res)
			return
		}
	}

	res := &Error{"API Error"}
	json.NewEncoder(w).Encode(res)
}

func GetUserName(m jwt.MapClaims) (string, bool) {
	val, ok := m["name"].(string)
	return val, ok
}

func JwtMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := CheckJwt(w, r); err == nil && next != nil {
			next.ServeHTTP(w, r)
		}
	}
}

func CheckJwt(w http.ResponseWriter, r *http.Request) error {
	tokenString, err := FromAuthHeader(r)
	if err != nil {
		return err
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return key, nil
	})

	newRequest := r.WithContext(context.WithValue(r.Context(), "User", token))
	*r = *newRequest

	return nil
}

func FromAuthHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", nil
	}

	authHeaderParts := strings.Fields(authHeader)
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return "", errors.New("Authorization header format must be Bearer {token}")
	}

	return authHeaderParts[1], nil
}
