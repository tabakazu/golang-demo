package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/tabakazu/webapi-app/app"
	"github.com/tabakazu/webapi-app/app/input"
	"github.com/tabakazu/webapi-app/app/output"
	"github.com/tabakazu/webapi-app/app/service"
	"github.com/tabakazu/webapi-app/domain"
)

type UserAccountController interface {
	RegisterHandler(http.ResponseWriter, *http.Request)
	LoginHandler(http.ResponseWriter, *http.Request)
	ShowHandler(http.ResponseWriter, *http.Request)
}

type userAccountController struct {
	registerService app.RegisterUserAccountService
	loginService    app.LoginUserAccountService
	showService     app.ShowUserAccountService
}

func NewUserAccountController(repo domain.UserAccountRepository) UserAccountController {
	return &userAccountController{
		registerService: service.NewRegisterUserAccount(repo),
		loginService:    service.NewLoginUserAccount(repo),
		showService:     service.NewShowUserAccount(repo),
	}
}

func (ctrl userAccountController) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var p input.RegisterUserAccountParam
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := ctrl.registerService.Execute(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func (ctrl userAccountController) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var p input.LoginUserAccountParam
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	data, err := ctrl.loginService.Execute(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat":     now.Unix(),
		"exp":     now.Add(time.Hour * 24).Unix(),
		"user_id": fmt.Sprintf("%d", data.ID),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	res := output.LoginUserAccountResult{
		Email: data.Email,
		Token: tokenString,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (ctrl userAccountController) ShowHandler(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, http.StatusText(401), http.StatusUnauthorized)
		return
	}

	authHeaderParts := strings.Fields(authHeader)
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		err := errors.New("Authorization header format must be Bearer {token}")
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	tokenString := authHeaderParts[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		http.Error(w, http.StatusText(401), http.StatusUnauthorized)
		return
	}

	val, ok := claims["user_id"].(string)
	if !ok {
		http.Error(w, http.StatusText(401), http.StatusUnauthorized)
		return
	}

	id, err := strconv.Atoi(val)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	data, err := ctrl.showService.Execute(input.ShowUserAccountParam{ID: uint(id)})
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)

}
