package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func init() {
}

func main() {
	addCommand := flag.NewFlagSet("add-admin", flag.ExitOnError)
	listCommand := flag.NewFlagSet("list-admins", flag.ExitOnError)
	listClinics := flag.NewFlagSet("list-clinics", flag.ExitOnError)
	listEmployees := flag.NewFlagSet("list-employees", flag.ExitOnError)
	getClinic := flag.NewFlagSet("get-clinics", flag.ExitOnError)
	listQuads := flag.NewFlagSet("list-quads", flag.ExitOnError)
	loginAdminCommand := flag.NewFlagSet("login-admin", flag.ExitOnError)
	addClinic := flag.NewFlagSet("add-clinic", flag.ExitOnError)
	addEmployee := flag.NewFlagSet("add-employee", flag.ExitOnError)
	deleteClinicCmd := flag.NewFlagSet("delete-clinic", flag.ExitOnError)

	// os.Arg[1] will be the subcommand
	if len(os.Args) < 2 {
		data, err := ioutil.ReadFile("README.md")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(data))
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add-admin":
		AddAdmin(addCommand)
	case "list-admins":
		ListAdmins(listCommand)
	case "login-admin":
		LoginAdmin(loginAdminCommand)
	case "add-clinic":
		AddClinic(addClinic)
	case "add-employee":
		AddEmployee(addEmployee)
	case "list-clinics":
		ListClinics(listClinics)
	case "list-employees":
		ListEmployees(listEmployees)
	case "get-clinic":
		GetClinic(getClinic)
	case "list-quads":
		ListQuads(listQuads)
	case "delete-clinic":
		DeleteClinic(deleteClinicCmd)
	default:
		flag.PrintDefaults()
		fmt.Println("Command not found")
		os.Exit(1)
	}
}
