package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/oren/doc-api"
)

// authenticate admin
// add new employee
func (ac AdminController) addEmployee(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value(0).(*admin.MyCustomClaims)

	newEmployee := &admin.NewEmployee{}

	err := json.NewDecoder(r.Body).Decode(newEmployee)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if newEmployee.FName == "" || newEmployee.LName == "" || newEmployee.Email == "" || newEmployee.ClinicId == "" {
		http.Error(w, "Employee must have fname, lname, email, and clinic_id fields", http.StatusInternalServerError)
		return
	}

	employee := &admin.NewEmployee{
		FName:    newEmployee.FName,
		LName:    newEmployee.LName,
		Email:    newEmployee.Email,
		ClinicId: newEmployee.ClinicId,
	}

	// // should i pass the employeeID from bolt instead of email?
	err = ac.adminService.AddEmployee(employee, claims.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(claims)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("in add employee")

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
