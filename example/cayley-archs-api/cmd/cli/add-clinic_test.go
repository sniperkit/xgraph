package main

import "testing"

// TestLoadJSON tests loading a JSON file
func TestLoadJSON(t *testing.T) {
	clinic := LoadJSON("test-data/clinic-all.json")
	if clinic.ClinicFullName != "Glenecclesia Clinic" {
		t.Errorf("expecting clinic name to be Glenecclesia Clinic, got %s", clinic.ClinicFullName)
	}
}

func TestRequiredFields(t *testing.T) {
	// load json
	// call validation function
	// make sure it's valid
	// result := Bar()
	// if result != "bar" {
	// 	t.Errorf("expecting bar, got %s", result)
	// }
}

func TestInsertToDB(t *testing.T) {
	// load json
	// call insert clinic
	// fetch new clinic from DB
	// compare clinic name
}
