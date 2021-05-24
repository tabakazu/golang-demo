package rest

import (
	"encoding/json"
	"net/http"

	"github.com/tabakazu/webapi-app/app"
	"github.com/tabakazu/webapi-app/app/input"
	"github.com/tabakazu/webapi-app/app/service"
	"github.com/tabakazu/webapi-app/domain"
)

type UserAccountController interface {
	RegisterHandler(http.ResponseWriter, *http.Request)
	LoginHandler(http.ResponseWriter, *http.Request)
}

type userAccountController struct {
	registerService app.RegisterUserAccountService
	loginService    app.LoginUserAccountService
}

func NewUserAccountController(repo domain.UserAccountRepository) UserAccountController {
	return &userAccountController{
		registerService: service.NewRegisterUserAccount(repo),
		loginService:    service.NewLoginUserAccount(repo),
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
