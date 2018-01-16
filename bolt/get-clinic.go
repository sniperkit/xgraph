package bolt

import (
	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/graph/path"
	"github.com/cayleygraph/cayley/quad"
	"github.com/cayleygraph/cayley/schema"
	"github.com/oren/doc-api"
)

func (a *AdminService) GetClinic(clinicId string) (admin.Clinic, error) {
	clinic, err := getClinic(a.Store, clinicId)

	if err != nil {
		return admin.Clinic{}, err
	}

	employees, err := getEmployees(a.Store, clinicId)
	clinic.Employees = employees

	if err != nil {
		return admin.Clinic{}, err
	}

	return clinic, nil
}

// get employeees of clinic id
func getEmployees(store *cayley.Handle, clinicId string) ([]admin.Employee, error) {
	var employees []admin.Employee
	p := path.StartPath(store).Has(quad.IRI("rdf:type"), quad.IRI("Employee")).
		Has(quad.IRI("workFor"), quad.IRI(clinicId))

	err := schema.LoadPathTo(nil, store, &employees, p)

	if err != nil {
		return employees, err
	}

	return employees, nil
}

// get clinic by it's id
func getClinic(store *cayley.Handle, clinicId string) (admin.Clinic, error) {
	v := quad.IRI(clinicId)

	var c admin.Clinic
	p := path.StartPath(store, v).Has(quad.IRI("rdf:type"), quad.IRI("Clinic"))

	err := schema.LoadPathTo(nil, store, &c, p)

	if err != nil {
		return c, err
	}

	return c, nil
}
