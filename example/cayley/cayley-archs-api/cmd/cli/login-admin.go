package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/oren/doc-api/bolt"
	"github.com/oren/doc-api/config"
)

func LoginAdmin(cmd *flag.FlagSet) {
	configPath := cmd.String("config", "", "Config file (Optional)")
	email := cmd.String("email", "", "Admin's email. (Required)")
	password := cmd.String("password", "", "Admin's password. (Required)")

	cmd.Parse(os.Args[2:])

	if !cmd.Parsed() {
		return
	}

	if *email == "" || *password == "" {
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

	var jwt string
	jwt, err = adminService.Login(*password, *email)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Admin exist. jwt:", jwt)
}
