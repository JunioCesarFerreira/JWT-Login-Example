package controller

import (
	"encoding/json"
	"fmt"
	"main/data"
	"main/middleware/cors"
	pkgLogin "main/middleware/login"
	"net/http"
)

type Controller struct {
	loginMethod pkgLogin.Login
	repository  data.Repository
}

func New(l pkgLogin.Login, r data.Repository) Controller {
	return Controller{
		loginMethod: l,
		repository:  r,
	}
}

// Rota para autenticação de login
func (c Controller) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login")

	cors.EnableCors(&w)

	if r.Method == "OPTIONS" {
		return
	}

	var user pkgLogin.DataLogin
	err := json.NewDecoder(r.Body).Decode(&user)
	fmt.Println("nameid:", user.UserId)
	fmt.Println("pwdHashed:", user.Pwd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	check, err := c.repository.CheckUser(user.UserId, user.Pwd)
	if !check || err != nil {
		http.Error(w, "Invalid name or password", http.StatusUnauthorized)
		return
	}

	resp, err := c.loginMethod.GetToken(user, 5)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

// Rota protegida com autenticação JWT
func (c Controller) ProtectedEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\nprotected endpoint")

	resp := map[string]string{"message": "Protected route.\nIt works!"}
	json.NewEncoder(w).Encode(resp)
}
