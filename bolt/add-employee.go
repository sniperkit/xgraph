package bolt

import (
	"fmt"

	"github.com/cayleygraph/cayley/quad"
	"github.com/oren/doc-api"
	uuid "github.com/satori/go.uuid"
)

func (a *AdminService) AddEmployee(employee *admin.NewEmployee, adminEmail string) error {
	if !validateEmployeeFields(employee) {
		return fmt.Errorf("Employee fields are not valid")
	}

	// get admin.ID from bolt
	// var foundAdmin Admin
	// foundAdmin, err = FindAdmin(store, claim.Email)
	// var id quad.IRI
	adminId, err := findAdminID(a.Store, adminEmail)

	if err != nil {
		return err
	}

	e := &admin.Employee{
		ID:        quad.IRI(uuid.NewV1().String()),
		FName:     employee.FName,
		LName:     employee.LName,
		Email:     employee.Email,
		WorkFor:   quad.IRI(employee.ClinicId),
		CreatedBy: adminId,
	}

	err = insert(a.Store, e)

	if err != nil {
		return err
	}

	return nil
}

// TODO: validate more fields of employee
func validateEmployeeFields(c *admin.NewEmployee) bool {
	if c.FName == "" || c.LName == "" {
		return false
	}

	return true
}
