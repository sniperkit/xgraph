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

func AddEmployee(cmd *flag.FlagSet) {
	configPath := cmd.String("config", "", "Config file (Optional)")
	employeeFName := cmd.String("fname", "", "Employee's first name. (Required)")
	employeeLName := cmd.String("lname", "", "Employee's last name. (Required)")
	employeeEmail := cmd.String("email", "", "Employee's email. (Required)")
	clinicId := cmd.String("clinic-id", "", "Clinic's Id. (Required)")
	adminJWT := cmd.String("jwt", "", "Admin's JWT. (Required)")

	cmd.Parse(os.Args[2:])

	if !cmd.Parsed() {
		return
	}

	// Required Flags
	if *employeeFName == "" || *employeeLName == "" || *employeeEmail == "" || *clinicId == "" || *adminJWT == "" {
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

	fmt.Println(claims)

	employee := &admin.NewEmployee{
		FName:    *employeeFName,
		LName:    *employeeLName,
		Email:    *employeeEmail,
		ClinicId: *clinicId,
	}

	// // should i pass the employeeID from bolt instead of email?
	err = adminService.AddEmployee(employee, claims.Email)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Employee was added")
}
