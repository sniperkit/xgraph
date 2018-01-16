package main

import (
	"encoding/json"
	"net/http"

	"github.com/oren/doc-api"
)

func (ac AdminController) adminLogin(w http.ResponseWriter, r *http.Request) {
	a := &admin.EmailAndPassword{}

	if err := json.NewDecoder(r.Body).Decode(a); err != nil {
		ServerError(w, err)
		return
	}

	if a.Email == "" || a.Password == "" {
		http.Error(w, "Login is missing email or password", http.StatusBadRequest)
		return
	}

	jwt, err := ac.adminService.Login(a.Password, a.Email)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	user := admin.User{a.Email, jwt}
	js, err := json.Marshal(user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(js)
}
