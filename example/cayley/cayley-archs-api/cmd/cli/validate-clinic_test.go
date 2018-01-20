package main

import (
	"testing"

	"github.com/oren/doc-api"
)

// TestValidateName test that clinic name is not empty
func TestValidateName(t *testing.T) {
	// create clinic with no name
	// call validate()

	clinic := admin.Clinic{ClinicFullName: ""}
	err := clinic.Validate()

	if err == nil {
		t.Errorf("expecting to get an error if clinic name is empty but got no error")
	}
}
