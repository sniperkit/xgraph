package main

import (
	"encoding/json"
	"net/http"

	"github.com/oren/doc-api"
)

// authenticate admin
// add new clinic
func (ac AdminController) addClinic(w http.ResponseWriter, r *http.Request) {
	newClinic := &admin.NewClinic{}

	err := json.NewDecoder(r.Body).Decode(newClinic)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if newClinic.GroupName == "" || newClinic.ClinicFullName == "" || newClinic.ClinicShortName == "" || newClinic.BranchName == "" || newClinic.Address1 == "" || newClinic.Address2 == "" || newClinic.AddressCity == "" || newClinic.AddressZipCode == "" || newClinic.AddressCountry == "" || newClinic.OfficeTel == "" || newClinic.OfficeFax == "" || newClinic.Website == "" || newClinic.ClinicEmail == "" || newClinic.ClinicAdminFullName == "" || newClinic.ClinicAdminMobile == "" || newClinic.ClinicAdminEmail == "" || newClinic.ClinicAdminPassword == "" || newClinic.ClinicAdminConfirmPassword == "" || !newClinic.AutoBidToggle || !newClinic.Priority || newClinic.NotificationsMobile == "" || newClinic.NotificationsEmail == "" || !newClinic.TogglePushNotifications || newClinic.PocFullName == "" || newClinic.PocMobile == "" || newClinic.PocEmail == "" {
		http.Error(w, "Clinic must have Group Name, Clinic Full Name, Clinic Short Name (displayed in search results), Branch Name, Address 1, Address 2, Address City, Address Zip Code, Address Country, Office Tel, Office Fax, Email, Clinic Admin Full Name, Clinic Admin Mobile, Clinic Admin Email, Clinic Admin Password, Clinic Admin Confirm Password, Auto bid toggle, Max Nb of Patients per Hour, Notifications Mobile, Notifications Email, Toggle Push Notifications, POC Full Name, POC Mobile, POC Email, fields", http.StatusInternalServerError)
		return
	}

	c := &admin.Clinic{
		GroupName:                  newClinic.GroupName,
		ClinicFullName:             newClinic.ClinicFullName,
		ClinicShortName:            newClinic.ClinicShortName,
		BranchName:                 newClinic.BranchName,
		Address1:                   newClinic.Address1,
		Address2:                   newClinic.Address2,
		AddressCity:                newClinic.AddressCity,
		AddressZipCode:             newClinic.AddressZipCode,
		AddressCountry:             newClinic.AddressCountry,
		OfficeTel:                  newClinic.OfficeTel,
		OfficeFax:                  newClinic.OfficeFax,
		Website:                    newClinic.Website,
		ClinicEmail:                newClinic.ClinicEmail,
		ClinicAdminFullName:        newClinic.ClinicAdminFullName,
		ClinicAdminMobile:          newClinic.ClinicAdminMobile,
		ClinicAdminEmail:           newClinic.ClinicAdminEmail,
		ClinicAdminPassword:        newClinic.ClinicAdminPassword,
		ClinicAdminConfirmPassword: newClinic.ClinicAdminConfirmPassword,
		AutoBidToggle:              newClinic.AutoBidToggle,
		Priority:                   newClinic.Priority,
		MaxNbOfPatientsPerHour:     newClinic.MaxNbOfPatientsPerHour,
		NotificationsMobile:        newClinic.NotificationsMobile,
		NotificationsEmail:         newClinic.NotificationsEmail,
		TogglePushNotifications:    newClinic.TogglePushNotifications,
		PocFullName:                newClinic.PocFullName,
		PocMobile:                  newClinic.PocMobile,
		PocEmail:                   newClinic.PocEmail,
	}

	claims := r.Context().Value(0).(*admin.MyCustomClaims)

	err = ac.adminService.AddClinic(c, claims.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(claims)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
