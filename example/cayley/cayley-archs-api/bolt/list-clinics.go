package bolt

import (
	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/schema"
	"github.com/oren/doc-api"
)

func (a *AdminService) AllClinics() ([]admin.Clinic, error) {
	clinics, err := readAllClinics(a.Store)
	clinics = sort(clinics, 1,3)

	if err != nil {
		return []admin.Clinic{}, err
	}

	return clinics, nil
}

func readAllClinics(store *cayley.Handle) ([]admin.Clinic, error) {
	// get all admins
	var clinics []admin.Clinic
	err := schema.LoadTo(nil, store, &clinics)

	return clinics, err
}

func sort(clinics []admin.Clinic, x float64, y float64)[]admin.Clinic {
	return clinics
}
