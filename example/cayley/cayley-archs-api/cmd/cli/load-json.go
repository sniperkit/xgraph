package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/oren/doc-api"
)

// LoadJson loads json of a clinic and returns a Clinic struct
func LoadJSON(JSONFile string) *admin.Clinic {
	raw, err := ioutil.ReadFile(JSONFile)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c admin.Clinic

	err = json.Unmarshal(raw, &c)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("clinic", &c)
	return &c
}
