package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/oren/doc-api"
	"github.com/oren/doc-api/bolt"
	"github.com/oren/doc-api/config"
)

// InsertClinic inserts clinic
func AddClinic(cmd *flag.FlagSet) {
	configPath := cmd.String("config", "", "Config file (Optional)")
	adminJWT := cmd.String("jwt", "", "Admin's JWT. (Required)")
	clinicPath := cmd.String("clinic-file", "", "Clinic JSON file (Required)")

	cmd.Parse(os.Args[2:])

	if !cmd.Parsed() {
		return
	}

	// Required Flags
	if *adminJWT == "" || *clinicPath == "" {
		cmd.PrintDefaults()
		os.Exit(1)
	}

	if *configPath == "" {
		*configPath = config.GetPathOfConfig()
	}

	configuration := config.ReadConf(*configPath)

	adminService, err := bolt.NewAdminService(configuration.DbPath)
	if err != nil {
		log.Fatal(err)
	}

	var claims *admin.MyCustomClaims
	claims, err = adminService.Authenticate(*adminJWT)
	if err != nil {
		log.Fatal(err)
	}

	clinic := LoadJSON(*clinicPath)
	err = clinic.Validate()

	err = adminService.AddClinic(clinic, claims.Email)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Clinic was added")
}
