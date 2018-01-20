package main_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/oren/doc-api"
)

// POST /adminlogin
// expect 200
// POST /clinics
// expect 200
func TestAddClinic(t *testing.T) {
	email := "vera@gmail.com"
	password := "112233"

	groupName := "Group Name"
	clinicFullName := "Clinic Full Name"
	clinicShortName := "Clinic Short Name (displayed in search results)"
	branchName := "Branch Name"
	address1 := "Address 1"
	address2 := "Address 2"
	addressCity := "Address City"
	addressZipCode := "Address Zip Code"
	addressCountry := "Address Country"
	officeTel := "Office Tel"
	officeFax := "Office Fax"
	website := "Website"
	clinicEmail := "Clinic Email"
	clinicAdminFullName := "Clinic Admin Full Name"
	clinicAdminMobile := "Clinic Admin Mobile"
	clinicAdminEmail := "Clinic Admin Email"
	clinicAdminPassword := "Clinic Admin Password"
	clinicAdminConfirmPassword := "Clinic Admin Confirm Password"
	autoBidToggle := "Auto bid toggle"
	priority := "Priority"
	maxNbOfPatientsPerHour := "Max Nb of Patients per Hour"
	notificationsMobile := "Notifications Mobile"
	notificationsEmail := "Notifications Email"
	togglePushNotifications := "Toggle Push Notifications"
	pocFullName := "POC Full Name"
	pocMobile := "POC Mobile"
	pocEmail := "POC Email"

	jwt := login(t, email, password)
	addClinic(t, jwt, groupName, clinicFullName, clinicShortName, branchName, address1, address2, addressCity, addressZipCode, addressCountry, officeTel, officeFax, website, clinicEmail, clinicAdminFullName, clinicAdminMobile, clinicAdminEmail, clinicAdminPassword, clinicAdminConfirmPassword, autoBidToggle, priority, maxNbOfPatientsPerHour, notificationsMobile, notificationsEmail, togglePushNotifications, pocFullName, pocMobile, pocEmail)
}

// curl -v -H "Authorization: Bearer verylongtokenstring" -H "Accept: application/json" -H "Content-type: application/json" POST -d '{"name":"a nice place","address1":"2323 hesting road, singapore"}' "http://localhost:3000/clinics"
func addClinic(t *testing.T, jwt string, groupName string, clinicFullName string, clinicShortName string, branchName string, address1 string, address2 string, addressCity string, addressZipCode string, addressCountry string, officeTel string, officeFax string, website string, clinicEmail string, clinicAdminFullName string, clinicAdminMobile string, clinicAdminEmail string, clinicAdminPassword string, clinicAdminConfirmPassword string, autoBidToggle string, priority string, maxNbOfPatientsPerHour string, notificationsMobile string, notificationsEmail string, togglePushNotifications string, pocFullName string, pocMobile string, pocEmail string) {
	url := "http://localhost:3000/clinics"

	newClinic := &admin.NewClinic{
		GroupName:                  groupName,
		ClinicFullName:             clinicFullName,
		ClinicShortName:            clinicShortName,
		BranchName:                 branchName,
		Address1:                   address1,
		Address2:                   address2,
		AddressCity:                addressCity,
		AddressZipCode:             addressZipCode,
		AddressCountry:             addressCountry,
		OfficeTel:                  officeTel,
		OfficeFax:                  officeFax,
		Website:                    website,
		ClinicEmail:                clinicEmail,
		ClinicAdminFullName:        clinicAdminFullName,
		ClinicAdminMobile:          clinicAdminMobile,
		ClinicAdminEmail:           clinicAdminEmail,
		ClinicAdminPassword:        clinicAdminPassword,
		ClinicAdminConfirmPassword: clinicAdminConfirmPassword,
		AutoBidToggle:              autoBidToggle,
		Priority:                   priority,
		MaxNbOfPatientsPerHour:     maxNbOfPatientsPerHour,
		NotificationsMobile:        notificationsMobile,
		NotificationsEmail:         notificationsEmail,
		TogglePushNotifications:    togglePushNotifications,
		PocFullName:                pocFullName,
		PocMobile:                  pocMobile,
		PocEmail:                   pocEmail,
	}

	js, err := json.Marshal(newClinic)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(js))
	req.Header.Set("Authorization", "Bearer "+jwt)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.Status == "200 OK" {
		return
	}

	t.Fatal("POST /clinics returned", resp.Status)
}
