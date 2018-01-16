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

func ListEmployees(cmd *flag.FlagSet) {
	configPath := cmd.String("config", "", "Config file (Optional)")
	cmd.Parse(os.Args[2:])

	if !cmd.Parsed() {
		return
	}

	if *configPath == "" {
		*configPath = config.GetPathOfConfig()
	}

	configuration := config.ReadConf(*configPath)

	adminService, err := bolt.NewAdminService(configuration.DbPath)
	if err != nil {
		log.Fatal(err)
	}

	var results []admin.Employee
	results, err = adminService.AllEmployees()

	if err != nil {
		log.Fatal(err)
	}

	printEmployees(results)
}

func printEmployees(employees []admin.Employee) {
	fmt.Println("\n==== All employees ====")

	for _, employee := range employees {
		fmt.Println("\tId: ", employee.ID)
		fmt.Println("\tFName: ", employee.FName)
		fmt.Println("\tLName: ", employee.LName)
		fmt.Println("\tEmail: ", employee.Email)
		fmt.Println("\tWork for: ", employee.WorkFor)
		fmt.Println("\tCreated By: ", employee.CreatedBy)

	}
}
