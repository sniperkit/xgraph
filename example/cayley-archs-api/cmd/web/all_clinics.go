package main

import (
	"encoding/json"
	"net/http"
)

// authenticate admin
// return all clinics
func (ac AdminController) allClinics(w http.ResponseWriter, r *http.Request) {
	clinics, err := ac.adminService.AllClinics()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(clinics)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
