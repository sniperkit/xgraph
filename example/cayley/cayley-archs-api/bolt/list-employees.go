package bolt

import (
	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/schema"
	"github.com/oren/doc-api"
)

func (a *AdminService) AllEmployees() ([]admin.Employee, error) {
	employees, err := readAllEmployees(a.Store)

	if err != nil {
		return []admin.Employee{}, err
	}

	return employees, nil
}

func readAllEmployees(store *cayley.Handle) ([]admin.Employee, error) {
	// get all admins
	var employees []admin.Employee
	err := schema.LoadTo(nil, store, &employees)

	return employees, err
}
