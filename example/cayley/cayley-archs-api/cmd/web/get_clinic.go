package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// authenticate admin
// return clinic
func (ac AdminController) getClinic(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	clinicId := vars["id"]

	clinic, err := ac.adminService.GetClinic(clinicId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(clinic)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
