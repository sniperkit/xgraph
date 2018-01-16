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

func GetClinic(cmd *flag.FlagSet) {
	configPath := cmd.String("config", "", "Config file (Optional)")
	clinicId := cmd.String("id", "", "Clinic's Id. (Required)")
	cmd.Parse(os.Args[2:])

	if !cmd.Parsed() {
		return
	}

	if *configPath == "" {
		*configPath = config.GetPathOfConfig()
	}

	// Required Flags
	if *clinicId == "" {
		cmd.PrintDefaults()
		os.Exit(1)
	}

	configuration := config.ReadConf(*configPath)

	adminService, err := bolt.NewAdminService(configuration.DbPath)
	if err != nil {
		log.Fatal(err)
	}

	var clinic admin.Clinic
	clinic, err = adminService.GetClinic(*clinicId)

	if err != nil {
		log.Fatal(err)
	}

	printClinic(clinic)
}

func printClinic(clinic admin.Clinic) {
	fmt.Println("\n==== clinic ====")

	fmt.Println("\tId: ", clinic.ID)

	fmt.Println("\tGroupName: ", clinic.GroupName)
	fmt.Println("\tClinicFullName: ", clinic.ClinicFullName)
	fmt.Println("\tClinicShortName: ", clinic.ClinicShortName)
	fmt.Println("\tBranchName: ", clinic.BranchName)
	fmt.Println("\tAddress1: ", clinic.Address1)
	fmt.Println("\tAddress2: ", clinic.Address2)
	fmt.Println("\tAddressCity: ", clinic.AddressCity)
	fmt.Println("\tAddressZipCode: ", clinic.AddressZipCode)
	fmt.Println("\tAddressCountry: ", clinic.AddressCountry)
	fmt.Println("\tOfficeTel: ", clinic.OfficeTel)
	fmt.Println("\tOfficeFax: ", clinic.OfficeFax)
	fmt.Println("\tWebsite: ", clinic.Website)
	fmt.Println("\tClinicEmail: ", clinic.ClinicEmail)
	fmt.Println("\tClinicAdminFullName: ", clinic.ClinicAdminFullName)
	fmt.Println("\tClinicAdminMobile: ", clinic.ClinicAdminMobile)
	fmt.Println("\tClinicAdminEmail: ", clinic.ClinicAdminEmail)
	fmt.Println("\tClinicAdminPassword: ", clinic.ClinicAdminPassword)
	fmt.Println("\tClinicAdminConfirmPassword: ", clinic.ClinicAdminConfirmPassword)
	fmt.Println("\tAutoBidToggle: ", clinic.AutoBidToggle)
	fmt.Println("\tPriority: ", clinic.Priority)
	fmt.Println("\tMaxNbOfPatientsPerHour: ", clinic.MaxNbOfPatientsPerHour)
	fmt.Println("\tNotificationsMobile: ", clinic.NotificationsMobile)
	fmt.Println("\tNotificationsEmail: ", clinic.NotificationsEmail)
	fmt.Println("\tTogglePushNotifications: ", clinic.TogglePushNotifications)
	fmt.Println("\tPocFullName: ", clinic.PocFullName)
	fmt.Println("\tPocMobile: ", clinic.PocMobile)
	fmt.Println("\tPocEmail: ", clinic.PocEmail)
	fmt.Println("\tCreated By: ", clinic.CreatedBy)

	fmt.Println("\n==== emyloyees ====")

	for _, employee := range clinic.Employees {
		fmt.Println("\tId: ", employee.ID)
		fmt.Println("\tFName: ", employee.FName)
		fmt.Println("\tLName: ", employee.LName)
		fmt.Println("\tEmail: ", employee.Email)
		fmt.Println("\tWork for: ", employee.WorkFor)
		fmt.Println("\tCreated By: ", employee.CreatedBy)
		fmt.Println("\n")

	}
}
