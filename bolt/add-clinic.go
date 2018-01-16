package bolt

import (
	"fmt"

	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/quad"
	"github.com/oren/doc-api"
	uuid "github.com/satori/go.uuid"
)

func (a *AdminService) AddClinic(clinic *admin.Clinic, email string) error {
	// get admin.ID from bolt
	// var foundAdmin Admin
	// foundAdmin, err = FindAdmin(store, claim.Email)
	// var id quad.IRI
	id, err := findAdminID(a.Store, email)

	if err != nil {
		return err
	}

	clinic.CreatedBy = id
	clinic.ID = quad.IRI(uuid.NewV1().String())
	err = insert(a.Store, clinic)

	if err != nil {
		fmt.Println("ooops")
		return err
	}

	return nil
}

func findAdminID(store *cayley.Handle, email string) (quad.IRI, error) {
	p := cayley.StartPath(store).Has(quad.IRI("email"), quad.String(email))
	id, err := p.Iterate(nil).FirstValue(nil)

	if err != nil {
		return "", err
	}

	return id.(quad.IRI), nil
}

// func validateClinicFields(c *admin.Clinic) bool {
// 	if c.GroupName == "" || c.ClinicFullName == "" || c.ClinicShortName == "" || c.BranchName == "" || c.Address1 == "" || c.Address2 == "" || c.AddressCity == "" || c.AddressZipCode == "" || c.AddressCountry == "" || c.OfficeTel == "" || c.OfficeFax == "" || c.Website == "" || c.ClinicEmail == "" || c.ClinicAdminFullName == "" || c.ClinicAdminMobile == "" || c.ClinicAdminEmail == "" || c.ClinicAdminPassword == "" || c.ClinicAdminConfirmPassword == "" || c.AutoBidToggle == "" || c.Priority == "" || c.MaxNbOfPatientsPerHour == "" || c.NotificationsMobile == "" || c.NotificationsEmail == "" || c.TogglePushNotifications == "" || c.PocFullName == "" || c.PocMobile == "" || c.PocEmail == "" {
// 		return false
// 	}

// 	return true
// }
